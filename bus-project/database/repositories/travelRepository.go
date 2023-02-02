package repositories

import (
	"golang_projects/models"
	"time"

	"gorm.io/gorm"
)

type TravelRepository struct {
	Travels *gorm.DB
}

type ITravelRepository interface {
	Add(travel *models.Travel) error
	Delete(id uint64) error
	GetAll() []models.Travel
	GetTravel(from string, to string, day time.Time, time time.Time) []models.Travel
	GetTravelById(id uint64) models.Travel
}

func NewTravelRepository(client *gorm.DB) TravelRepository {
	return TravelRepository{Travels: client}
}

func (repository TravelRepository) Add(travel *models.Travel) error {
	if err := repository.Travels.Create(travel).Error; err != nil {
		return err
	}
	return nil
}
func (repository TravelRepository) Delete(id uint64) error {
	var travel models.Travel
	if err := repository.Travels.Unscoped().Delete(&travel, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository TravelRepository) GetAll() []models.Travel {
	var travels []models.Travel
	repository.Travels.
		Preload("Buses").
		Preload("Buses.BusModel").
		Preload("Buses.BusModel.Brand").
		Preload("Buses.Type").
		Preload("Buses.Properties").
		Preload("Buses.Seats").
		Preload("Buses.Seats.Properties").
		Find(&travels)
	return travels
}

func (repository TravelRepository) GetTravel(from string, to string, day time.Time, time time.Time) []models.Travel {
	var travels *[]models.Travel
	//repository.Travels.Preload("Bus.BusModel.Brand").Preload("Bus.BusModel").Preload("Bus").Find(&travels)
	repository.Travels.
		Preload("Buses").
		Preload("Buses.BusModel").
		Preload("Buses.BusModel.Brand").
		Preload("Buses.Type").
		Preload("Buses.Properties").
		Preload("Buses.Seats").
		Preload("Buses.Seats.Properties").
		Find(&travels, `from_location = ? AND to_location = ?`, from, to)

	return *travels
}

func (repository TravelRepository) GetTravelById(id uint64) models.Travel {
	var travel *models.Travel

	repository.Travels.
		Preload("Buses").
		Preload("Buses.Seats").
		Preload("Buses.Seats.Properties").
		Find(&travel, id)
	return *travel
}
