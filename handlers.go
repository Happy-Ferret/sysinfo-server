package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/boltdb/bolt"
)

func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, maxDatagramSize)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	go writeToDb(buffer[:n])
}

func handlerHTTP(w http.ResponseWriter, r *http.Request) {
	var output []string
	dbc.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("bucket"))
		if err := b.ForEach(func(key []byte, value []byte) error {
			output = append(output, string(value))
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", output)
}

func httpServer() {
	http.HandleFunc("/", handlerHTTP)
	http.ListenAndServe(*host+":"+*tport, nil)
}
