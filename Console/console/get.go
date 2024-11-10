package console

import (
	"fmt"

	console_user "github.com/VandiKond/user-http-go/Console/user"
)

// Gets data with user
//
// user -- a pointer to a user (if nil sets name to guest)
// a -- any data for fmt.Scanln
func Get(user *console_user.User, a ...any) {
	// Sets the default name
	var name string = "guest"
	if user != nil {
		// Ads the users login
		name = user.Login
	}

	// Ads the user name on the string
	fmt.Printf("$:%s> ", name)
	// Scans the line
	fmt.Scanln(a...)
}
