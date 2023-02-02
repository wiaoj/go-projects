package models

import "gorm.io/gorm"

type BusTravel struct {
	gorm.Model
	BusId    uint
	Bus      Bus
	TravelId uint
	Travel   Travel
}
