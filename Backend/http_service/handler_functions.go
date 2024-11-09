package http_service

import (
	"encoding/json"
	"net/http"

	"github.com/VandiKond/user-http-go/Backend/file_worker"
	"github.com/VandiKond/user-http-go/Backend/user"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	users, _, err := file_worker.GetUsersArray()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	User, ok, err := user.Login(username, password, &users)
	if err != nil {
		if err.Error() == user.UNF {
			http.Error(w, "User not found", http.StatusInternalServerError)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(User)

}
