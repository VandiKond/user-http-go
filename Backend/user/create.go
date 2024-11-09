package user

import (
	"errors"

	hash_password "github.com/VandiKond/user-http-go/Backend/hash"
	"github.com/VandiKond/user-http-go/Backend/operations"
)

// Creates a new user
//
// login -- users login
// password -- users password
// allUsers -- a slice of users
//
// Returns:
// val 1 : error
func NewUser(login string, password string, allUsers *[]User) (*User, error) {
	// Checks that user exists
	ok := CheckExistence(login, *allUsers)
	if ok != nil {
		return nil, errors.New(UAE)
	}

	// Creates a password
	hash, err := hash_password.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Logs the operation
	operations.NewOperation(CREATE_USER, "{}", login, 0)

	// Ads the user
	user := User{Login: login, passwordHash: hash}
	*allUsers = append(*allUsers, user)

	return &user, nil

}
