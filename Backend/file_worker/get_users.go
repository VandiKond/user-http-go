package file_worker

import (
	"github.com/VandiKond/user-http-go/Backend/user"
)

// Gets users and backup users from files and returning them as an array
//
// Returns:
// val 1: an array of users
// val 2: an error
func GetUsersArray() ([]user.User, error) {
	mainUserText, err := GetAll()
	if err != nil {
		return nil, err
	}
	users, err := GetUserFromJSON(mainUserText)
	if err != nil {
		return nil, err
	}

	return users, nil
}
