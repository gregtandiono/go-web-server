package main

import (
	"encoding/json"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/boltdb/bolt"
	uuid "github.com/satori/go.uuid"
)

// User struct
type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

// NewUser func returns a new instance of User
func NewUser(id uuid.UUID, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

// Save adds a new User record
func (u *User) Save() error {
	var s Storage
	db := s.Init()
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("USERS"))
		mu, _ := json.Marshal(u)
		err := bkt.Put([]byte(u.ID.String()), []byte(string(mu)))
		return err
	})
}

// Fetch returns one single user based on user id
func (u *User) Fetch() User {
	var s Storage
	db := s.Init()
	defer db.Close()

	var user User

	db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("USERS"))
		v := bkt.Get([]byte(u.ID.String()))
		err := json.Unmarshal(v, &user)
		if err != nil {
			return err
		}
		return nil
	})

	return user
}

// FetchAll returns an array of users
func (u *User) FetchAll() []User {
	users := []User{}
	for i := 0; i < 20; i++ {
		var user User
		user.ID = uuid.NewV4()
		user.Name = randomdata.FullName(randomdata.RandomGender)
		user.Email = randomdata.Email()
		users = append(users, user)
	}
	return users
}
