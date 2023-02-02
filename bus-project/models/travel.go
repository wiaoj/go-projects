package models

import (
	"gorm.io/gorm"
)

type Travel struct {
	gorm.Model
	Fee          float32
	FromLocation string
	ToLocation   string
	Day          string
	Time         string
	Buses        []Bus `gorm:"many2many:busses_travels;OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Seats        []Seat
}
