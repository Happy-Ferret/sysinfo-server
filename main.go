package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/boltdb/bolt"
)

const (
	maxDatagramSize = 4096
)

var (
	host   = flag.String("host", "0.0.0.0", "The sysinfo-server address.")
	uport  = flag.String("uport", "9000", "The sysinfo-server udp port.")
	tport  = flag.String("tport", "9000", "The sysinfo-server http managment port.")
	dbFile = "bolt.db"
	dbc    *bolt.DB
)

func init() {
	flag.Parse()
	var err error
	dbc, err = bolt.Open(dbFile, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(dbc, "connected")
}

func main() {
	service := *host + ":" + *uport
	defer dbc.Close()
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
	log.Println("Server up over proto udp and listening on port", *uport)
	defer ln.Close()

	go httpServer()
	log.Println("HTTP managment Server up and listening on port", *tport)

	for {
		// wait for UDP client to connect
		handleUDPConnection(ln)
	}
}
