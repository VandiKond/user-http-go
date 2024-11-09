package file_worker

import (
	"github.com/VandiKond/user-http-go/Backend/user"
)

// Gets users and backup users from files and returning them as an array
//
// Returns
// Users -- an array of users
// BackupUsers -- an array of backup users
// Error -- the error
func GetUsersArray() (Users []user.User, Error error) {
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
