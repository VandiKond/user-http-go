package register

import (
	"fmt"

	"github.com/VandiKond/user-http-go/Console/console"
)

func GetData() (string, string) {
	fmt.Println("Please enter login")
	var login string
	console.Get(nil, &login)
	fmt.Println("Please enter password")
	var password string
	console.GetPassword(nil, &password)
	return login, password
}
