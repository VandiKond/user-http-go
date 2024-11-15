package file_worker

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/VandiKond/user-http-go/Backend/user"
)

// Saves users
//
// users -- a slice of users
//
// Returns
// val 1: error
func SaveUsers(users []user.User) error {
	return saveAny(MAIN_FILE, users)
}

// Saves the backup users
//
// backupUsers -- a slice of users
//
// Returns
// val 1: error
func SaveBackupUsers(backupUsers []user.User) error {
	return saveAny(BACKUP_FILE, backupUsers)
}

// Saves any file
//
// fileName -- the name of the file
// saveData (any) -- the data to save
//
// Returns:
// val 1: an error
func saveAny(fileName string, saveData any) error {
	data, err := json.Marshal(saveData)
	if err != nil {
		return errors.New(EMU)
	}

	return os.WriteFile(fileName, []byte(data), 0666)
}
