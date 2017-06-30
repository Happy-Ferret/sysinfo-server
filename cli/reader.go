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
