package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

const (
	maxDatagramSize = 4096
)

var (
	host   = flag.String("host", "0.0.0.0", "The sysinfo-server address.")
	port   = flag.String("port", "9000", "The sysinfo-server port.")
	proto  = flag.String("proto", "udp", "UDP or TCP.")
	dbFile = "bolt.db"
)

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
	defer server.db.Close()

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
