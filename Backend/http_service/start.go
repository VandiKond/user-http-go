package http_service

import (
	"encoding/json"
	"log"
	"net/http"
)

// Starts the http server
func Start() {
	// Handler for login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Uses the login handler
		data, err := loginHandler(w, r)
		if err != nil {
			return
		}

		// Sends the user data
		json.NewEncoder(w).Encode(data)

	})

	// Handler for deleting user
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		// Uses the login handler
		_, err := loginHandler(w, r)
		if err != nil {
			return
		}

		// deletes the user
		deleteHandler(w, r)
	})

	// Handler for register
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		// using the register handler
		registerHandler(w, r)
	})

	// Starting the server
	http.ListenAndServe(":8080", nil)
	log.Panicln("Started server on port 8080")

}
