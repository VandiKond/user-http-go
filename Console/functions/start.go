package functions

import (
	"fmt"

	"github.com/VandiKond/user-http-go/Console/config"
	"github.com/VandiKond/user-http-go/Console/console"
	"github.com/VandiKond/user-http-go/Console/login"
	"github.com/VandiKond/user-http-go/Console/register"
	console_user "github.com/VandiKond/user-http-go/Console/user"
)

func Start() error {
	fmt.Println("Do you want to sign up or sing in? write 'up' to create a new account, write 'in' to log in")
	var answer string
	console.Get(nil, &answer)
	var user *console_user.User
	var err error
	switch answer {
	case "up":
		user, err = register.Register("You've chosen creating a new account. Do you want to continue?")
		if err != nil || user == nil {
			return Start()
		}
	case "in":
		user, err = login.LogIn("You've chosen logging in. Do you want to continue?")
		if err != nil || user == nil {
			return Start()
		}
	default:
		fmt.Println("Invalid input")
		return Start()
	}
	for {
		console.Get(user, &answer)
		commands := GetFunctions()
		if command, ok := commands[answer]; ok {
			doReturn, err := command.Usage(command.HiMessage, user)
			if err != nil {
				config.EXIT_WITH_ERROR()
			}
			if doReturn {
				if err != nil {
					return err
				}
				return Start()
			}
		} else {
			fmt.Println("Invalid command.")
		}
	}
}
