package services

import (
	"golang_projects/database/repositories"
	"golang_projects/models"
)

type SeatService struct {
	Repository repositories.ISeatRepository
}

type ISeatService interface {
	AddSeat(aeat *models.Seat) error
	AddSeatProperty(seatProperty *models.SeatProperty) error
	DeleteSeat(id uint64) error
	GetAllSeats() []models.Seat
}

func NewSeatService(repository repositories.ISeatRepository) SeatService {
	return SeatService{Repository: repository}
}

func (service SeatService) AddSeat(seat *models.Seat) error {
	if err := service.Repository.Add(seat); err != nil {
		return err
	}
	return nil
}
func (service SeatService) AddSeatProperty(seatProperty *models.SeatProperty) error {
	if err := service.Repository.AddProperty(seatProperty); err != nil {
		return err
	}
	return nil
}
func (service SeatService) DeleteSeat(id uint64) error {
	if err := service.Repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (service SeatService) GetAllSeats() []models.Seat {
	return service.Repository.GetAll()
}
