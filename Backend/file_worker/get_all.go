package file_worker

import (
	"errors"
	"io"
	"os"

	"github.com/VandiKond/user-http-go/Backend/user"
)

// Gets all users
//
// Returns:
// val 1: a string with data
// val 2: error
func GetAll() (string, error) {
	// Getting all text
	mainText, err := os.ReadFile(MAIN_FILE)
	if err == io.EOF {
		saveErr := SaveUsers([]user.User{})
		if saveErr != nil {
			return "", saveErr
		}

	} else if err != nil {
		return "", errors.New(ERF)
	}

	// Returning all files
	return string(mainText), nil
}
