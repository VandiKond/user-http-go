package http_service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/VandiKond/user-http-go/Backend/file_worker"
	"github.com/VandiKond/user-http-go/Backend/user"
)

type RegistrationData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Creates a new user
func registerHandler(w http.ResponseWriter, r *http.Request) {
	// If the wrong method returning status 405
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Getting body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Getting registration data
	var data RegistrationData
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// getting all users
	users, err := file_worker.GetUsersArray()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Creates a new user
	newUser, err := user.NewUser(data.Username, data.Password, &users)
	if err != nil {
		// If user exists returning status 409
		if err.Error() == user.UAE {
			http.Error(w, "User already exists", http.StatusConflict)
			return
			// if input data is not valid returning status 406
		} else if err.Error() == user.NVD {
			http.Error(w, "Not valid input data", http.StatusNotAcceptable)
			return
			// if input data is to short returning status 406
		} else if err.Error() == user.DTS {
			http.Error(w, "Input data to short", http.StatusNotAcceptable)
			return
		}
		// retuning status 500
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Saves users
	err = file_worker.SaveUsers(users)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Encodes and sends the user
	json.NewEncoder(w).Encode(newUser)
}
