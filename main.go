package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
)

const (
	maxDatagramSize = 4096
)

var (
	host               = flag.String("host", "0.0.0.0", "The sysinfo-server address.")
	uport              = flag.String("uport", "9000", "The sysinfo-server udp port.")
	tport              = flag.String("tport", "9000", "The sysinfo-server http managment port.")
	dbFile             = "bolt.db"
	databaseConnection server
)

// Env exported
type Env struct {
	db *bolt.DB
}

func handleUDPConnection(conn *net.UDPConn, s *server) {
	buffer := make([]byte, maxDatagramSize)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	go writeToDb(buffer[:n], s)
}

func (env *Env) handler(w http.ResponseWriter, r *http.Request) {
	var output []string

	fmt.Fprintf(w, "%v", output)
}

// func httpServer() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(*host+":"+*tport, nil)
// }

func init() {
	flag.Parse()
	db, err := bolt.Open(dbFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Printf("Error: %s", err)
	}
}

func httpServer() {
	env := &Env{db: db}
	http.HandleFunc("/", env.handler)
	http.ListenAndServe(*host+":"+*tport, nil)
}
func main() {
	service := *host + ":" + *uport

	// bolt
	server, err := newServer(dbFile)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	databaseConnection = *server
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

	log.Println("Server up over proto udp and listening on port", *uport)
	log.Println("HTTP managment Server up and listening on port", *tport)

	for {
		// wait for UDP client to connect
		handleUDPConnection(ln, server)
	}
}
