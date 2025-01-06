package http

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

// comparing password with hash
func CheckPasswordH(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// changes the password into string of letters or numbers using encryption algorithm
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// checking if the given input is email or not.
func IsEmail(em string) bool {
	_, err := mail.ParseAddress(em)
	return err == nil
}
