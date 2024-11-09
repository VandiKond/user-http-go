package file_worker

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/VandiKond/user-http-go/Backend/user"
)

func SaveUsers(users []user.User) error {
	return saveAny(MAIN_FILE, users)
}

func SaveBackupUsers(backupUsers []user.User) error {
	return saveAny(BACKUP_FILE, backupUsers)
}

func saveAny(fileName string, saveData any) error {
	data, err := json.Marshal(saveData)
	if err != nil {
		return errors.New(EMU)
	}

	return os.WriteFile(fileName, []byte(data), 0666)
}
