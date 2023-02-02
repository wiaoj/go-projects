package services

import (
	"golang_projects/database/repositories"
	"golang_projects/models"
)

type LocationService struct {
	Repository repositories.ILocationRepository
}

type ILocationService interface {
	AddLocation(location *models.Location) error
	DeleteLocation(id uint64) error
	GetAllLocations() []models.Location
}

func NewLocationService(repository repositories.ILocationRepository) LocationService {
	return LocationService{Repository: repository}
}

func (service LocationService) AddLocation(location *models.Location) error {
	if err := service.Repository.Add(location); err != nil {
		return err
	}
	return nil
}

func (service LocationService) DeleteLocation(id uint64) error {
	if err := service.Repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (service LocationService) GetAllLocations() []models.Location {
	return service.Repository.GetAll()
}
