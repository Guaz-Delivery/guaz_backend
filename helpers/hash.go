package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// Helper function to hash password
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}
