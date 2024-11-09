package hash_password

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

const (
	EGH = "error getting hash"
)

// A function that's getting the password and returning it's hash (sha-3-256)
//
// password -- the user password
//
// Returns:
// val 1 : the hashed password
// val 2 : the error
func HashPassword(password string) (string, error) {
	// Creating a 256 byte hash as a slice of bytes
	hash := sha3.New256()
	_, err := hash.Write([]byte(password))

	// In case of error returning the error
	if err != nil {
		return "", fmt.Errorf("%s: %w", EGH, err)
	}

	// Writing the slice in sha3
	sha3 := hash.Sum(nil)

	// Returning the slice as a string
	return fmt.Sprintf("%x", sha3), nil
}

// A function to compare hash of a password and the password
//
// password -- the users password to check
// hash -- an hashed data
//
// Returns:
// val 1 : do the passwords match
// val 2 : error
func CompareHash(password string, hash string) (bool, error) {
	// Creating the hash of the password string
	hashedPassword, err := HashPassword(password)

	// In case of error returning the error
	if err != nil {
		return false, err
	}

	// Returning true if the hash are the same. If not false
	return hashedPassword == hash, nil
}
