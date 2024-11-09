package http_service

import "net/http"

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		basicHandler(w, r)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		registerHandler(w, r)
	})

	http.ListenAndServe(":8080", nil)

}
