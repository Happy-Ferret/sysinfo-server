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

// GetAll comment
func GetAll(s *server) (data []string, err error) {
	var output []string
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("bucket"))
		if err := b.ForEach(func(key []byte, value []byte) error {
			output = append(output, string(value))
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
	return output, nil
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
		}

		log.Printf("%+v\n", si.Node.MachineID)

		data, err := Get(s, "bucket", si.Node.MachineID)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		log.Println(string(data))
	} else {
		fmt.Println(string(msg), "wasn't write")
	}
}
