package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) []byte {
	passwordSalt := []byte(password)
	passwordHash, hashError := bcrypt.GenerateFromPassword(passwordSalt, bcrypt.DefaultCost)

	if hashError != nil {
		return nil
	}

	return passwordHash
}

func CheckPassword(passwordHash []byte, password string) error {
	passwordSalt := []byte(password)
	compareError := bcrypt.CompareHashAndPassword(passwordHash, passwordSalt)
	if compareError != nil {
		return compareError
	}
	return nil
}
