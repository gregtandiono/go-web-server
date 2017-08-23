package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Pallinder/go-randomdata"

	uuid "github.com/satori/go.uuid"

	"github.com/boltdb/bolt"
)

func TestMain(m *testing.M) {
	seedAndTeardownDB()
}

func seedAndTeardownDB() {
	var s Storage
	db := s.Init()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		fmt.Println("######## Tearing Down Bucket(s) ########")
		err := tx.DeleteBucket([]byte("USERS"))
		if err != nil {
			return err
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		fmt.Println("####### Create and Seed Bucket(s) #######")
		bkt, err := tx.CreateBucketIfNotExists([]byte("USERS"))
		if err != nil {
			return err
		}
		for i := 0; i < 20; i++ {
			u := User{uuid.NewV4(), randomdata.FullName(randomdata.RandomGender), randomdata.Email()}
			mu, _ := json.Marshal(u)
			bkt.Put([]byte(u.ID.String()), []byte(string(mu)))
		}
		return nil
	})
}
