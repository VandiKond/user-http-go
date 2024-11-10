package user

// Checks if the user exists
//
// login -- the users login
// users -- a slice of users
//
// Returns:
// val 1: pointer to a user or nil
func CheckExistence(login string, users []User) (*User, int) {
	for i, user := range users {
		if user.Login == login && !user.Deleted {
			return &user, i
		}
	}
	return nil, -1
}
