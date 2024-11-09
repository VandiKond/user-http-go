package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RegistrationData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// Create registration data
	registrationData := RegistrationData{
		Username: "Vandi",
		Password: "10iui78!90kj",
	}

	// Encode data to JSON
	jsonData, err := json.Marshal(registrationData)
	if err != nil {
		panic(err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "http://localhost:8080/register", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	// Set content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Print the response status code and body
	fmt.Printf("Status code: %d\n", resp.StatusCode)
	fmt.Println(string(body))
}
