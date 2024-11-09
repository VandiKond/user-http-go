package console

import (
	"fmt"

	console_user "github.com/VandiKond/user-http-go/Console/user"
)

func Get(user *console_user.User, a ...any) {
	var name string = "guest"
	if user != nil {
		name = user.Login
	}
	fmt.Printf("$:%s> ", name)
	fmt.Scanln(a...)
}
