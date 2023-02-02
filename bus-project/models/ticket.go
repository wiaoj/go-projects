package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	SeatId uint
	Seat   Seat
	No     uint
	Gender bool
}
