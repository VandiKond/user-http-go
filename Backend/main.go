package main

import (
	start "github.com/VandiKond/user-http-go/Backend/main"
)

func main() {
	err := start.StartAll()
	if err != nil {
		panic(err)
	}
}
