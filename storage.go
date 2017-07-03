package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// Put comment
func Put(bucket, key string, val []byte) error {

	return dbc.Update(func(tx *bolt.Tx) error {
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
func Get(bucket, key string) (data []byte, err error) {
	dbc.View(func(tx *bolt.Tx) error {
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

func writeToDb(msg []byte) {
	var si SysInfo
	if err := json.Unmarshal(msg, &si); err != nil {
		log.Println(err)
	}
	if si.Node.MachineID != "" {
		err := Put("bucket", si.Node.MachineID, msg)
		if err != nil {
			log.Printf("Error: %s", err)
		}

		log.Printf("%+v\n", si.Node.MachineID)

		data, err := Get("bucket", si.Node.MachineID)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		log.Println(string(data))
	} else {
		fmt.Println(string(msg), "wasn't write")
	}
}
