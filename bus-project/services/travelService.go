package services

import (
	"golang_projects/database/repositories"
	"golang_projects/models"
	"strings"
	"time"
)

type TravelService struct {
	Repository repositories.ITravelRepository
}

type ITravelService interface {
	AddTravel(travel *models.Travel) error
	DeleteTravel(id uint64) error
	GetAllTravels() []models.Travel
	GetTravel(from string, to string, day time.Time, time time.Time) []models.Travel
	GetTravelById(id uint64) models.Travel
}

func NewTravelService(repository repositories.ITravelRepository) TravelService {
	return TravelService{Repository: repository}
}

func (service TravelService) AddTravel(travel *models.Travel) error {
	strings.ToLower(travel.FromLocation)
	strings.ToLower(travel.ToLocation)
	if err := service.Repository.Add(travel); err != nil {
		return err
	}
	return nil
}

func (service TravelService) DeleteTravel(id uint64) error {
	if err := service.Repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (service TravelService) GetAllTravels() []models.Travel {
	return service.Repository.GetAll()
}

func (service TravelService) GetTravel(from string, to string, day time.Time, time time.Time) []models.Travel {
	strings.ToLower(from)
	strings.ToLower(to)
	return service.Repository.GetTravel(from, to, day, time)
}

func (service TravelService) GetTravelById(id uint64) models.Travel {
	return service.Repository.GetTravelById(id)
}
