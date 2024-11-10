package functions

import (
	"fmt"

	console_user "github.com/VandiKond/user-http-go/Console/user"
)

func Help(hiMessage string, user *console_user.User) (bool, error) {
	if len(hiMessage) <= 0 {
		hiMessage = "All commands:"
	}
	fmt.Println(hiMessage)
	for i, d := range GetFunctions() {
		fmt.Printf("%s: %s\n", i, d.Description)
	}
	return false, nil
}
