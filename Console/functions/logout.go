package functions

import (
	"fmt"

	"github.com/VandiKond/user-http-go/Console/console"
	"github.com/VandiKond/user-http-go/Console/user"
)

func LogOut(hiMessage string, user *user.User) (bool, error) {
	if len(hiMessage) <= 0 {
		hiMessage = "Do you want to log out?"
	}
	fmt.Printf("%s y/n base:n\n", hiMessage)
	var DoContinue string
	console.Get(user, &DoContinue)
	if len(DoContinue) <= 0 {
		DoContinue = "n"
	}
	if DoContinue != "n" && DoContinue != "y" {
		fmt.Println("Wrong input")
		return LogOut(hiMessage, user)
	} else if DoContinue == "n" {
		fmt.Println("Exiting.")
		return false, nil
	}
	return true, Start()
}
