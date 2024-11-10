package file_worker

import (
	"errors"
	"os"
)

// Gets or creates a file
//
// fileName -- the name of the file
//
// Returns :
// val 1: file
// val 2: error
func GetFile(fileName string) (*os.File, error) {
	// Trying to create a file
	file, err := os.Create(fileName)
	if err == nil {
		if fileName == MAIN_FILE {
			os.WriteFile(fileName, []byte("[]"), 0666)
		}
		// if the file was created returning it
		return file, nil
	}

	// Opening the file
	file, err = os.Open(fileName)
	if err != nil {
		// Returning the error
		return nil, errors.New(EOF)
	}

	// Returning the file
	return file, nil
}
