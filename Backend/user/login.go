package user

import (
	"errors"

	hash_password "github.com/VandiKond/user-http-go/Backend/hash"
	"github.com/VandiKond/user-http-go/Backend/operations"
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
	user := CheckExistence(login, *users)
	if user == nil {
		return nil, false, errors.New(UNF)
	}

	// Compares passwords
	ok, err := hash_password.CompareHash(password, user.passwordHash)
	if !ok {
		return nil, false, err
	}

	// Logs the login
	operations.NewOperation(LOGIN, "{}", login, 0)

	return user, true, nil
}
