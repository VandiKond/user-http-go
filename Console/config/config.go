package config

import "fmt"

const (
	// The path to the host
	PATH = "http://localhost:8080"
)

// Sends an error message
func EXIT_WITH_ERROR() {
	fmt.Println("Unknown error")
}
