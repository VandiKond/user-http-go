package file_worker

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/VandiKond/user-http-go/Backend/user"
)

// Unmarshal json data to user.User
//
// Returns:
// val 1: an array of users
// val 2: an error
func GetUserFromJSON(jsonData string) ([]user.User, error) {
	users := []user.User{}
	err := json.Unmarshal([]byte(jsonData), &users)
	if err != nil {
		fmt.Println(err, jsonData)
		return nil, errors.New(EUJ)
	}
	return users, nil
}
