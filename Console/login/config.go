package login

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/VandiKond/user-http-go/Console/config"
)

func GetResponse(login string, password string, url string, method string, body io.Reader, errorLoginMessage string) (*http.Response, error) {
	req, err := http.NewRequest(method, config.PATH+url, body)
	if err != nil {
		config.EXIT_WITH_ERROR()
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(login, password)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		config.EXIT_WITH_ERROR()
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Println(errorLoginMessage)
		return resp, errors.New(ELI)
	}
	if resp.StatusCode != http.StatusOK {
		config.EXIT_WITH_ERROR()
		return resp, errors.New("unknown error")
	}

	return resp, nil
}
