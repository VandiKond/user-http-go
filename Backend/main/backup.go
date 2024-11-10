package start

import (
	"log"

	"github.com/VandiKond/user-http-go/Backend/file_worker"
	"github.com/robfig/cron/v3"
)

// Starting the job
func StartJob() error {
	// Creating a new cron data
	c := cron.New()
	// Setting a function on 9 AM
	_, err := c.AddFunc("* 9 * * *", Job)
	if err != nil {
		return err
	}
	// Starting all jobs
	c.Start()
	return nil
}

// The backup job
func Job() {
	// Gets all data
	mainFile, err := file_worker.GetUsersArray()
	if err != nil {
		log.Println("[ERROR]: ", err)
	}
	// Saves the backup users
	file_worker.SaveBackupUsers(mainFile)
}
