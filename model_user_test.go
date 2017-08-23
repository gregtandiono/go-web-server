package main

import (
	"testing"

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

}

func TestUserFetch(t *testing.T) {

}

func TestUserFetchAll(t *testing.T) {

}
