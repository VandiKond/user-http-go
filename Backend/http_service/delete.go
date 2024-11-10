package http_service

import (
	"encoding/json"
	"net/http"

	"github.com/VandiKond/user-http-go/Backend/file_worker"
	"github.com/VandiKond/user-http-go/Backend/user"
)

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// If the wrong method returning status 405
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// getting all users
	users, err := file_worker.GetUsersArray()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Deletes the user
	username, _, _ := r.BasicAuth()
	delUser, err := user.Delete(username, users)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Saves users
	err = file_worker.SaveUsers(users)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Encodes and sends the user
	json.NewEncoder(w).Encode(delUser)

}
