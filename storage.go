package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
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
	var si SysInfo
	if err := json.Unmarshal(msg, &si); err != nil {
		log.Println(err)
	}
	if si.Node.MachineID != "" {
		err := Put(s, "bucket", si.Node.MachineID, msg)
		if err != nil {
			log.Printf("Error: %s", err)
			// log.Fatalf("Error: %s", err)
		}

		fmt.Printf("%+v\n", si.Node.MachineID)

		data, err := Get(s, "bucket", si.Node.MachineID)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		fmt.Println(string(data))
	}
	fmt.Println(string(msg))
}
