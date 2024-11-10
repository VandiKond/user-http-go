package functions

import (
	"fmt"

	"github.com/VandiKond/user-http-go/Console/user"
)

func GetMe(hiMessage string, user *user.User) (bool, error) {
	if len(hiMessage) <= 0 {
		hiMessage = "Your data:"
	}
	fmt.Printf("%s\nlogin: %s\nbalance: %.2f\n", hiMessage, user.Login, user.Balance)
	if user.Rank > 0 {
		fmt.Printf("admin rank: %d\n", user.Rank)
	}
	return false, nil
}
