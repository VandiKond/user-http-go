package file_worker

import (
	"fmt"
	"os"
)

// Gets all files that are in config.go
//
// Returns:
// MainFile -- the main data file
// BackupFile -- the backup data file
// LogFile -- the file with logs
// Error -- the error if has one
func GetAll() (MainFile *os.File, BackupFile *os.File, LogFile *os.File, Error error) {
	// getting the main file
	mainFile, err := GetFile(MAIN_FILE)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("%s %s: %w", EGF, MAIN_FILE, err)
	}

	// getting the backup file
	backupFile, err := GetFile(BACKUP_FILE)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("%s %s: %w", EGF, BACKUP_FILE, err)
	}

	// getting the log file
	logFile, err := GetFile(LOG_FILE)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("%s %s: %w", EGF, LOG_FILE, err)
	}

	// Returning all files
	return mainFile, backupFile, logFile, nil
}
