package models

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Value  string `gorm:"uniqueIndex;not null"`
	Busses []Bus  `gorm:"many2many:busses_properties;OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
