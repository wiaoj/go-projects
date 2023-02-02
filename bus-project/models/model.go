package models

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Value   string `gorm:"not null"`
	BrandId uint
	Brand   Brand
	Buses   []Bus `gorm:"foreignKey:BusModelId"`
}
