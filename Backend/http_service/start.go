package http_service

import (
	"encoding/json"
	"net/http"
)

// Starts the http server
func Start() {
	// Handler for login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Uses the basic handler
		data, err := basicHandler(w, r)
		if err != nil {
			return
		}

		// Sends the user data
		json.NewEncoder(w).Encode(data)

	})

	// Handler for register
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		// using the register handler
		registerHandler(w, r)
	})

	// Starting the server
	http.ListenAndServe(":8080", nil)

}
