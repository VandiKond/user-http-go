package functions

import (
	"fmt"

	"github.com/VandiKond/user-http-go/Console/user"
)

func Help(hiMessage string, user *user.User) (bool, error) {
	if len(hiMessage) <= 0 {
		hiMessage = "All commands:"
	}
	fmt.Println(hiMessage)
	for i, d := range GetFunctions() {
		fmt.Printf("%s: %s\n", i, d.Description)
	}
	return false, nil
}
