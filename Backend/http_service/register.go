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

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var data RegistrationData
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	users, _, err := file_worker.GetUsersArray()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	newUser, err := user.NewUser(data.Username, data.Password, &users)
	if err != nil {
		if err.Error() == user.UAE {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		} else if err.Error() == user.NVD {
			http.Error(w, "Not valid input data", http.StatusNotAcceptable)
			return
		} else if err.Error() == user.DTS {
			http.Error(w, "Input data to short", http.StatusNotAcceptable)
		}
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
	}
	err = file_worker.SaveUsers(users)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newUser)
}
