package user

import (
	"errors"
	"log"
)

// Deletes a user
//
// login -- users login
// allUsers -- a slice of users
//
// Returns:
// val 1 : User that is deleted
// val 2 : error
func Delete(login string, allUsers []User) (*User, error) {
	// Gets the user
	_, id := CheckExistence(login, allUsers)
	if id < 0 {
		// if user not exists
		return nil, errors.New(UNF)
	}
	// Deletes the user
	allUsers[id].Deleted = true
	log.Println(DELETE, login)
	// Returning the deleted user
	return &allUsers[id], nil
}
