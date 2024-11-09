package file_worker

import (
	"fmt"
	"io"
)

// Gets all files and returns string values of main data and backup data files
//
// Returns:
// main_text -- the text of the main file
// backup_text -- the text of the backup file
// any_error -- the error
func SortAll() (main_text string, backup_text string, any_error error) {
	// Getting all files
	mainFile, backupFile, _, err := GetAll()
	if err != nil {
		return "", "", err
	}

	// Reading the main file
	mainText, err := io.ReadAll(mainFile)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", ERF, err)
	}

	// Reading the backup file
	backupText, err := io.ReadAll(backupFile)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", ERF, err)
	}

	// Returning the data, it has got
	return string(mainText), string(backupText), nil
}
