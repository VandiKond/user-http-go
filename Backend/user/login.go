package user

import (
	"errors"
	"log"

	hash_password "github.com/VandiKond/user-http-go/Backend/hash"
)

// Logs in a user
//
// login -- users login
// password -- users password
//
// Returns:
// val 1 : the user
// val 2 : does the user log in
// val 3 : error
func Login(login string, password string, users *[]User) (*User, bool, error) {
	// Checks that the user exists
	user, _ := CheckExistence(login, *users)
	if user == nil {
		return nil, false, errors.New(UNF)
	}

	// Compares passwords
	ok, err := hash_password.CompareHash(password, user.PasswordHash)
	if !ok {
		return nil, false, err
	}

	log.Println(LOGIN, login)

	return user, true, nil
}
