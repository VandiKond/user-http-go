package http_service

import (
	"errors"
	"net/http"

	"github.com/VandiKond/user-http-go/Backend/file_worker"
	"github.com/VandiKond/user-http-go/Backend/user"
)

// A handler to authorize
func loginHandler(w http.ResponseWriter, r *http.Request) (*user.User, error) {
	// Getting the password and username
	username, password, ok := r.BasicAuth()

	// If authorization is not set returning status 401
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}

	// Getting all users
	users, err := file_worker.GetUsersArray()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return nil, errors.New("unknown error")
	}

	// Logging in
	User, ok, err := user.Login(username, password, &users)
	if err != nil {
		// If user not found returning status 401
		if err.Error() == user.UNF {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return nil, errors.New("user not found")
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return nil, errors.New("unknown error")
	}
	// If password or login wrong returning status 401
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return nil, errors.New("unauthorized")

	}

	return User, nil

}
