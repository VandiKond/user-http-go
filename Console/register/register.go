package register

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/VandiKond/user-http-go/Console/config"
	"github.com/VandiKond/user-http-go/Console/console"
	console_user "github.com/VandiKond/user-http-go/Console/user"
)

func Register(hiMessage string) (*console_user.User, error) {
	fmt.Printf("%s y/n base:y\n", hiMessage)
	var DoContinue string
	console.Get(nil, &DoContinue)
	if len(DoContinue) <= 0 {
		DoContinue = "y"
	} else if DoContinue != "n" && DoContinue != "y" {
		fmt.Println("Wrong input")
		return Register(hiMessage)
	} else if DoContinue == "n" {
		fmt.Println("Exiting.")
		return nil, nil
	}
	login, password := GetData()

	registerData := RegistrationData{Username: login, Password: password}

	jsonData, err := json.Marshal(registerData)
	if err != nil {
		config.EXIT_WITH_ERROR()
		return nil, err
	}

	req, err := http.NewRequest("POST", config.PATH+"/register", bytes.NewBuffer(jsonData))
	if err != nil {
		config.EXIT_WITH_ERROR()
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		config.EXIT_WITH_ERROR()
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		if resp.StatusCode == http.StatusConflict {
			fmt.Println("User with this login exists. please register again")

		} else if resp.StatusCode == http.StatusNotAcceptable {
			fmt.Println("Input data is not valid. please register again")
		} else {
			config.EXIT_WITH_ERROR()
			return nil, errors.New(resp.Status)
		}
		return Register("Do you want to continue?")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		config.EXIT_WITH_ERROR()
		return nil, err
	}

	var user console_user.User

	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		config.EXIT_WITH_ERROR()
		return nil, err
	}

	fmt.Printf("Registration successful. Logged in user %s\n", user.Login)
	return &user, nil

}
