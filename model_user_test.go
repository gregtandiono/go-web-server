package main

import (
	"encoding/json"
	"testing"

	"github.com/boltdb/bolt"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	nu := NewUser(
		uuid.FromStringOrNil("39232b56-b094-4ca0-ac52-e5ad9cdb7f8d"),
		"Minnie Mouse",
		"minnie@disney.com",
	)
	assert.Equal(t, nu.Name, "Minnie Mouse", "OK")
	assert.Equal(t, nu.Email, "minnie@disney.com", "OK")
}

func TestUserSave(t *testing.T) {

	u := NewUser(
		uuid.FromStringOrNil("c52639a8-d1ba-4886-8fe8-49818a84d314"),
		"Mickey Mouse",
		"mickey@disney.com",
	)
	err := u.Save()
	if err != nil {
		t.Error("TestUserSave failed \n", err)
	}

	// Now let's verify if the record is in the DB
	var s Storage
	db := s.Init()
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		var user User
		bkt := tx.Bucket([]byte("USERS"))
		v := bkt.Get([]byte("c52639a8-d1ba-4886-8fe8-49818a84d314"))
		err := json.Unmarshal(v, &user)
		if err != nil {
			return err
		}
		assert.Equal(t, user.Name, "Mickey Mouse", "OK")
		assert.Equal(t, user.Email, "mickey@disney.com", "OK")
		return nil
	})

}

func TestUserFetch(t *testing.T) {
	// we're going to fetch the record that was created in the previous test
	var user User
	user.ID = uuid.FromStringOrNil("c52639a8-d1ba-4886-8fe8-49818a84d314")
	u := user.Fetch()

	assert.Equal(t, u.Name, "Mickey Mouse", "OK")
	assert.Equal(t, u.Email, "mickey@disney.com", "OK")
}

func TestUserDelete(t *testing.T) {
	var user User
	user.ID = uuid.FromStringOrNil("c52639a8-d1ba-4886-8fe8-49818a84d314")
	err := user.Delete()
	if err != nil {
		t.Error("failed to delete user", err)
	}

	// let's verify if the record is still there

	var s Storage
	db := s.Init()
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("USERS"))
		v := bkt.Get([]byte("c52639a8-d1ba-4886-8fe8-49818a84d314"))
		assert.Equal(t, []byte(nil), v, "OK. User does not exist anymore")
		return nil
	})

}

func TestUserFetchAll(t *testing.T) {
	var user User
	people := user.FetchAll()
	assert.Equal(t, len(people), 20, "OK")
}
