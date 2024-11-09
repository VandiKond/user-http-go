package console_user

// The basic user
type User struct {
	// Users login
	Login string `json:"login"`
	// Hash of users password
	PasswordHash string `json:"password_hash"`
	// The users admin rank
	Rank int `json:"rank"`
	// The user balance
	Balance float64 `json:"balance"`
}
