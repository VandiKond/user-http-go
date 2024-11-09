package user

import (
	"errors"

	hash_password "github.com/VandiKond/user-http-go/Backend/hash"
	"github.com/VandiKond/user-http-go/Backend/operations"
)

func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && r != '!' && r != '*' && r != '_' && r != '-' && r != '&' && r != '#' {
			return false
		}
	}
	return true
}

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

	// Cheeking that login and password are valid
	if !IsLetter(login) || !IsLetter(password) {
		return nil, errors.New(NVD)
	}

	if len(login) < 4 || len(password) < 10 {
		return nil, errors.New(DTS)
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
