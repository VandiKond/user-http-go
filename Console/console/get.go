package console

import (
	"bufio"
	"fmt"
	"os"

	"github.com/VandiKond/user-http-go/Console/user"
	"golang.org/x/crypto/ssh/terminal"
)

// Gets data with user
//
// user -- a pointer to a user (if nil sets name to guest)
// val -- string pointer to a value
func Get(user *user.User, val *string) {
	// Sets the default name
	var name string = "guest"
	if user != nil {
		// Ads the users login
		name = user.Login
	}

	// Ads the user name on the string
	fmt.Printf("$:%s> ", name)
	//Scans the line
	reader := bufio.NewReader(os.Stdin)

	// Read the entire line
	input, _ := reader.ReadString('\n')

	// Remove the newline character
	*val = input[:len(input)-2]

}

// Gets data with user
//
// user -- a pointer to a user (if nil sets name to guest)
// val -- string pointer to a value
func GetPassword(user *user.User, val *string) {
	// Sets the default name
	var name string = "guest"
	if user != nil {
		// Ads the users login
		name = user.Login
	}

	// Ads the user name on the string
	fmt.Printf("$:%s> ", name)

	// Scans the password
	password, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	*val = string(password)
}
