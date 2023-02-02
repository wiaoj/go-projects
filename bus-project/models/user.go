package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string `gorm:"email;uniqueIndex;not null"`
	PasswordHash []byte
}
