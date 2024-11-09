package main

import (
	"errors"
	"fmt"

	"github.com/VandiKond/user-http-go/Console/console"
	"github.com/VandiKond/user-http-go/Console/login"
	"github.com/VandiKond/user-http-go/Console/register"
	console_user "github.com/VandiKond/user-http-go/Console/user"
)

func main() {
	fmt.Println("Hi! let's start")
	panic(start())
}

func start() error {
	fmt.Println("do you want to sign up or sing in? write 'up' to create a new account, write 'in' to log in")
	var answer string
	console.Get(nil, &answer)
	var user *console_user.User
	var err error
	switch answer {
	case "up":
		user, err = register.Register("You've chosen creating a new account. Do you want to continue?")
		if err != nil || user == nil {
			return start()
		}
	case "in":
		user, err = login.LogIn("You've chosen logging in. Do you want to continue?")
		if err != nil || user == nil {
			return start()
		}
	default:
		fmt.Println("Invalid input")
	}
	console.Get(user, &answer)
	return errors.New(answer)
}
