package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"

	uuid "github.com/satori/go.uuid"

	"github.com/boltdb/bolt"
)

func TestMain(m *testing.M) {
	seedAndTeardownDB()
	os.Exit(m.Run())
}

func TestFetchAllUsersHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/users", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	var people []User
	err := json.Unmarshal(response.Body.Bytes(), &people)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, len(people), 20, "OK response is expected")
}

func TestNewUserHandler(t *testing.T) {
	mockData := []byte(`{
		"name": "Gregory Tandiono",
		"email": "gregtandiono@gmail.com",
		"id": "f4bc5612-cb97-4a63-9381-a558ebc0bae5"
	}`)
	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(mockData))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "response is positive!")
}

func TestFetchOneUserHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/users/f4bc5612-cb97-4a63-9381-a558ebc0bae5", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	var user User
	err := json.Unmarshal(response.Body.Bytes(), &user)
	if err != nil {
		t.Error("Failed to fetch user", err)
	}
	assert.Equal(t, user.Name, "Gregory Tandiono", "OK response is expected")
	assert.Equal(t, user.Email, "gregtandiono@gmail.com", "OK response is expected")
	assert.Equal(t, user.ID.String(), "f4bc5612-cb97-4a63-9381-a558ebc0bae5", "OK response is expected")
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
