package models

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	Count      int
	TravelId   uint
	Travel     Travel
	BusId      uint
	Bus        Bus
	Properties []SeatProperty
}

type SeatProperty struct {
	gorm.Model
	SeatId uint
	Seat   Seat
	No     int
	Gender bool
}
