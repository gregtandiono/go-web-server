package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

// Storage struct
type Storage struct{}

// Init starts the db connection pool and returns a bolt.DB
func (s *Storage) Init() *bolt.DB {
	db, err := bolt.Open("store.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// BucketInit creates buckets
func (s *Storage) BucketInit() {
	db := s.Init()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		fmt.Println("################# INITIALIZING BUCKETS #################")
		_, err := tx.CreateBucketIfNotExists([]byte("USERS"))
		if err != nil {
			return err
		}
		return nil
	})
}
