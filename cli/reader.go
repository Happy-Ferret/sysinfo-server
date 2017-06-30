package main

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

type server struct {
	db *bolt.DB
}

var (
	dbFile = "../bolt.db"
)

func newServer(filename string) (s *server, err error) {
	s = &server{}
	s.db, err = bolt.Open(filename, 0600, &bolt.Options{Timeout: 1 * time.Second})
	return
}

func main() {
	// bolt
	server, err := newServer(dbFile)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	server.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("bucket"))
		if err := b.ForEach(func(key []byte, value []byte) error {
			log.Println(string(key), string(value))
			return nil
		}); err != nil {
			return err
		}
		return nil
	})

}
