package user

import (
	"errors"
	"log"

	hash_password "github.com/VandiKond/user-http-go/Backend/hash"
)

func isValidText(text string) bool {
	for _, char := range text {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') ||
			char == '!' || char == '#' || char == '&' || char == '*' || char == '-' ||
			char == '_' || char == '?' || char == ',' || char == '.' ||
			(char >= '0' && char <= '9')) {
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
// val 1 : user that is created
// val 2 : error
func NewUser(login string, password string, allUsers *[]User) (*User, error) {
	// Checks that user exists
	ok, _ := CheckExistence(login, *allUsers)
	if ok != nil {
		return nil, errors.New(UAE)
	}
	// Cheeking that login and password are valid
	if !isValidText(login) || !isValidText(password) {
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

	// Ads the user
	user := User{Login: login, PasswordHash: hash}
	*allUsers = append(*allUsers, user)

	// Logs the creation
	log.Println(CREATE_USER, login)

	return &user, nil

}
