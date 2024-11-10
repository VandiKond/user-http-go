package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/VandiKond/user-http-go/Console/config"
	"github.com/VandiKond/user-http-go/Console/console"
	"github.com/VandiKond/user-http-go/Console/register"
	"github.com/VandiKond/user-http-go/Console/user"
)

func LogIn(hiMessage string) (*user.User, error) {
	fmt.Printf("%s y/n default: y\n", hiMessage)
	var DoContinue string
	console.Get(nil, &DoContinue)
	if len(DoContinue) <= 0 {
		DoContinue = "y"
	} else if DoContinue != "n" && DoContinue != "y" {
		fmt.Println("Wrong input")
		return LogIn(hiMessage)
	} else if DoContinue == "n" {
		fmt.Println("Exiting.")
		return nil, nil
	}

	login, password := register.GetData()

	resp, err := GetResponse(login, password, "/login", http.MethodGet, bytes.NewBuffer([]byte("{}")), "Please enter valid username and password")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil && err.Error() == ELI {
		return LogIn("Do you want to continue?")
	} else if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		config.EXIT_WITH_ERROR()
		return nil, err
	}

	var user user.User

	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		config.EXIT_WITH_ERROR()
		return nil, err
	}

	fmt.Println("Log in successful")
	return &user, nil
}
