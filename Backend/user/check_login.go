package user

// Checks if the user exists
//
// login -- the users login
// users -- a slice of users
//
// Returns:
// val 1: pointer to a user or nil
func CheckExistence(login string, users []User) *User {
	for _, user := range users {
		if user.Login == login {
			return &user
		}
	}
	return nil
}
