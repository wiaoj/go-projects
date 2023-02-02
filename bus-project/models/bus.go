package models

import "gorm.io/gorm"

type Bus struct {
	gorm.Model
	PlateNumber string `gorm:"uniqueIndex;not null"`
	SeatsCount  int
	// SeatId     uint
	Seats      []Seat
	BusModelId uint
	BusModel   Model
	TypeId     uint
	Type       Type
	Properties []Property `gorm:"many2many:busses_properties;OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Travel     []Travel   `gorm:"many2many:busses_travels;OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
