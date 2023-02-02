package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Name   string  `gorm:"uniqueIndex;not null"`
	Models []Model // `gorm:"foreignKey:BrandId"`
}
