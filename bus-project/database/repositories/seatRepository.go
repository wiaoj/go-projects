package repositories

import (
	"golang_projects/models"

	"gorm.io/gorm"
)

type SeatRepository struct {
	Seats *gorm.DB
}

type ISeatRepository interface {
	Add(seat *models.Seat) error
	AddProperty(seatProperty *models.SeatProperty) error
	Delete(id uint64) error
	GetAll() []models.Seat
}

func NewSeatRepository(client *gorm.DB) SeatRepository {
	return SeatRepository{Seats: client}
}

func (repository SeatRepository) Add(seat *models.Seat) error {
	if err := repository.Seats.Create(seat).Error; err != nil {
		return err
	}
	return nil
}
func (repository SeatRepository) AddProperty(seatProperty *models.SeatProperty) error {
	if err := repository.Seats.Create(seatProperty).Error; err != nil {
		return err
	}
	return nil
}

func (repository SeatRepository) Delete(id uint64) error {
	var Seat models.Seat
	if err := repository.Seats.Unscoped().Delete(&Seat, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository SeatRepository) GetAll() []models.Seat {
	var seats []models.Seat
	repository.Seats.Find(&seats)
	return seats
}
