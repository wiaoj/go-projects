package repositories

import (
	"golang_projects/models"
	"strings"

	"gorm.io/gorm"
)

type LocationRepository struct {
	Locations *gorm.DB
}

type ILocationRepository interface {
	Add(location *models.Location) error
	Delete(id uint64) error
	GetAll() []models.Location
}

func NewLocationRepository(client *gorm.DB) LocationRepository {
	return LocationRepository{Locations: client}
}

func (repository LocationRepository) Add(location *models.Location) error {
	strings.ToLower(location.Name)
	if err := repository.Locations.Create(location).Error; err != nil {
		return err
	}
	return nil
}
func (repository LocationRepository) Delete(id uint64) error {
	var location models.Location
	if err := repository.Locations.Unscoped().Delete(&location, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository LocationRepository) GetAll() []models.Location {
	var locations []models.Location
	repository.Locations.Find(&locations)
	return locations
}
