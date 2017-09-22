package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	maxDatagramSize = 4096
)

func startNetServer() {
	service := "0.0.0.0:" + *netPort

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

		fmt.Println("Server up over proto", *proto, "and listening on port", *netPort)

		for {
			// wait for TCP client to connect
			conn, err := ln.Accept()
			if err != nil {
				log.Fatalf("Error: %s", err)
			}
			handleTCPConnection(conn)
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

		log.Println("Server up over proto", *proto, "and listening on port", *netPort)

		for {
			// wait for UDP client to connect
			handleUDPConnection(ln)
		}
	}

}

func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, maxDatagramSize)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	err = writeDataToDB(a.DB, buffer[:n])
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func handleTCPConnection(conn net.Conn) {
	buf := &bytes.Buffer{}
	for {
		data := make([]byte, maxDatagramSize)
		n, err := conn.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("Fatal: %s", err)
			} else {
				break
			}
		}
		buf.Write(data[:n])

	}

	defer conn.Close()

	err := writeDataToDB(a.DB, buf.Bytes())
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
