package operations

// It represents a operation
type Operation struct {
	// Operation type
	Type string `json:"type"`
	// The data (it could be almost anything, but better in json format)
	Data string `json:"data"`
	// The user that made the operation (users login)
	User string `json:"user"`
	// The operation rank
	Rank int `json:"rank"`
}
