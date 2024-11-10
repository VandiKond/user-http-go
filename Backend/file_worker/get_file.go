package file_worker

import (
	"errors"
	"io"
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
	// Trying to open the file
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		// Returning the error
		return nil, errors.New(EOF)
	}

	// If the file was created and is MAIN_FILE, writing initial content
	if fileName == MAIN_FILE {
		// Check if the file is empty
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			return nil, err
		}
		if _, err := file.Read([]byte{}); err != io.EOF {
			// File is not empty, do nothing
		} else {
			// File is empty, write initial content
			_, err = file.Write([]byte("[]"))
			if err != nil {
				return nil, err
			}
		}
	}

	// Returning the file
	return file, nil
}
