package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

// Combine password and salt then hash them using the SHA-512
func HashFunc(password string, salt []byte) string {
	// Convert password string to byte slice
	var pwdByte = []byte(password)

	// Create sha-512 hasher
	var sha512 = sha512.New()

	pwdByte = append(pwdByte, salt...)

	sha512.Write(pwdByte)

	// Get the SHA-512 hashed password
	var hashedPassword = sha512.Sum(nil)

	// Convert the hashed to hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPassword)
	return hashedPasswordHex
}

// Check if two passwords match
func PasswordsMatch(hashedPassword, currPassword string, salt []byte) bool {
	var currPasswordHash = HashFunc(currPassword, salt)

	return hashedPassword == currPasswordHash
}
