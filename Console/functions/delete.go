package functions

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/VandiKond/user-http-go/Console/console"
	"github.com/VandiKond/user-http-go/Console/login"
	"github.com/VandiKond/user-http-go/Console/user"
)

func Delete(hiMessage string, user *user.User) (bool, error) {
	if len(hiMessage) <= 0 {
		hiMessage = "Do you want to delete your account?"
	}
	fmt.Printf("%s y/n base:n\n", hiMessage)
	var DoContinue string
	console.Get(user, &DoContinue)
	if len(DoContinue) <= 0 {
		DoContinue = "n"
	}
	if DoContinue != "n" && DoContinue != "y" {
		fmt.Println("Wrong input")
		return Delete(hiMessage, user)
	} else if DoContinue == "n" {
		fmt.Println("Exiting.")
		return false, nil
	}

	fmt.Println("Please enter password")
	var password string
	console.Get(nil, &password)

	fmt.Printf("Are you sure? y/n base:n\n")
	console.Get(user, &DoContinue)
	if len(DoContinue) <= 0 {
		DoContinue = "n"
	}
	if DoContinue != "n" && DoContinue != "y" {
		fmt.Println("Wrong input")
		return Delete("Lets try again. Do you want to continue?", user)
	} else if DoContinue == "n" {
		fmt.Println("Exiting.")
		return false, nil
	}

	_, err := login.GetResponse(user.Login, password, "/delete", http.MethodDelete, bytes.NewBuffer([]byte("")), "Wrong password")
	if err != nil {
		if err.Error() == login.ELI {
			return Delete(hiMessage, user)
		}
		return false, err
	}

	fmt.Printf("User %s deleted. Returning to the registration page\n", user.Login)
	return true, nil
}
