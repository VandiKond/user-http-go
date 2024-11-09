package file_worker

import (
	"encoding/json"
	"fmt"

	"github.com/VandiKond/user-http-go/Backend/user"
)

// Unmarshal json data to user.User
//
// Returns:
// Users -- an array of users
// Error -- an error
func GetUserFromJSON(jsonData string) (Users []user.User, Error error) {
	users := []user.User{}
	err := json.Unmarshal([]byte(jsonData), &users)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", EUJ, err)
	}
	return users, nil
}
