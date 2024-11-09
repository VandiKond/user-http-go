package file_worker

import "github.com/VandiKond/user-http-go/Backend/user"

// Gets users and backup users from files and returning them as an array
//
// Returns
// Users -- an array of users
// BackupUsers -- an array of backup users
// Error -- the error
func GetUsersArray() (Users []user.User, BackupUsers []user.User, Error error) {
	mainUserText, BackupUserText, err := SortAll()
	if err != nil {
		return nil, nil, err
	}

	users, err := GetUserFromJSON(mainUserText)
	if err != nil {
		return nil, nil, err
	}

	backupUsers, err := GetUserFromJSON(BackupUserText)
	if err != nil {
		return nil, nil, err
	}

	return users, backupUsers, nil
}
