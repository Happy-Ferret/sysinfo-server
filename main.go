package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/boltdb/bolt"
	"github.com/mvaleev/sysinfo"
)

const (
	maxDatagramSize = 4096
)

var (
	si     sysinfo.SysInfo
	host   = flag.String("host", "0.0.0.0", "The sysinfo-server address.")
	port   = flag.String("port", "9999", "The sysinfo-server port.")
	proto  = flag.String("proto", "udp", "UDP or TCP.")
	dbFile = "bolt.db"
)

type server struct {
	db *bolt.DB
}

func newServer(filename string) (s *server, err error) {
	s = &server{}
	s.db, err = bolt.Open(filename, 0600, &bolt.Options{Timeout: 1 * time.Second})
	return
}

// Put comment
func Put(s *server, bucket, key string, val []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		if err = b.Put([]byte(key), val); err != nil {
			return err
		}
		return err
	})
}

// Get comment
func Get(s *server, bucket, key string) (data []byte, err error) {
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		r := b.Get([]byte(key))
		if r != nil {
			data = make([]byte, len(r))
			copy(data, r)
		}
		return nil
	})
	return
}

func writeToDb(msg []byte, s *server) {
	if err := json.Unmarshal(msg, &si); err != nil {
		log.Println(err)
	}

	err := Put(s, "bucket", si.Node.MachineID, msg)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("%+v\n", si.Node.MachineID)

	data, _ := Get(s, "bucket", si.Node.MachineID)
	fmt.Println(string(data))
}

func handleUDPConnection(conn *net.UDPConn, s *server) {
	buffer := make([]byte, maxDatagramSize)

	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	go writeToDb(buffer[:n], s)
}

func handleTCPConnection(conn net.Conn, s *server) {
	buffer := make([]byte, maxDatagramSize)

	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer conn.Close()

	go writeToDb(buffer[:n], s)
}

func main() {
	flag.Parse()
	service := *host + ":" + *port

	// bolt
	server, err := newServer(dbFile)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	switch p := *proto; p {
	case "tcp":
		// tcp
		udpAddr, err := net.ResolveTCPAddr("tcp", service)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		// setup listener for incoming TCP connection
		ln, err := net.ListenTCP("tcp", udpAddr)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		defer ln.Close()

		fmt.Println("Server up over proto", *proto, "and listening on port", *port)

		for {
			// wait for TCP client to connect
			conn, err := ln.Accept()
			if err != nil {
				log.Fatalf("Error: %s", err)
			}
			handleTCPConnection(conn, server)
		}
	default:
		// udp
		udpAddr, err := net.ResolveUDPAddr("udp4", service)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		// setup listener for incoming UDP connection
		ln, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		ln.SetReadBuffer(maxDatagramSize)
		defer ln.Close()

		fmt.Println("Server up over proto", *proto, "and listening on port", *port)

		for {
			// wait for UDP client to connect
			handleUDPConnection(ln, server)
		}
	}
}
