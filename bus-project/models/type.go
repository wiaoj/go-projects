package models

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Value string `gorm:"uniqueIndex;not null"`
}
