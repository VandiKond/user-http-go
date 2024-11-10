package start

import (
	"log"

	"github.com/VandiKond/user-http-go/Backend/file_worker"
	"github.com/VandiKond/user-http-go/Backend/http_service"
)

// Starts all
func StartAll() error {
	// creates main file if not exists
	file_worker.GetFile(file_worker.MAIN_FILE)

	// set ups the logging
	logFile, err := file_worker.GetFile(file_worker.LOG_FILE)
	if err != nil {
		return err
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Starting the job
	file_worker.GetFile(file_worker.BACKUP_FILE)
	StartJob()

	// Starting the http server
	http_service.Start()

	return nil
}
