package functions

import console_user "github.com/VandiKond/user-http-go/Console/user"

type Function struct {
	Description string
	Usage       func(string, *console_user.User) (bool, error)
	HiMessage   string
}

func GetFunctions() map[string]Function {
	return map[string]Function{
		"out": {
			Description: "Log out of your account",
			Usage:       LogOut,
		},
		"help": {
			Description: "Shows all commands",
			Usage:       Help,
		},
	}
}
