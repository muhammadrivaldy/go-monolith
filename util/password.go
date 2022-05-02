package util

import "golang.org/x/crypto/bcrypt"

// GeneratePassword is a function for encrypted password
func GeneratePassword(s string) (string, error) {
	data, err := bcrypt.GenerateFromPassword([]byte(s), 0)
	return string(data), err
}

// ComparePassword is a function for compare hash & password
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
