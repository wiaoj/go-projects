package repositories

import (
	"golang_projects/models"

	"gorm.io/gorm"
)

type BusRepository struct {
	Busses *gorm.DB
}

type IBusRepository interface {
	BusDefintion() *models.BusDefintion
	Update(bus *models.Bus) error
	Add(bus *models.Bus) error
	GetAll() []*models.Bus
	Delete(id uint64) error
	GetWhere(query interface{}, args ...interface{}) *models.Bus
}

func NewBusRepository(client *gorm.DB) BusRepository {
	return BusRepository{Busses: client}
}

func (repository BusRepository) Add(bus *models.Bus) error {
	if err := repository.Busses.Create(bus).Error; err != nil {
		return err
	}
	return nil
}

func (repository BusRepository) GetAll() []*models.Bus {
	var busses []*models.Bus
	repository.Busses.
		Preload("BusModel.Brand").
		Preload("BusModel").
		Preload("Type").
		Preload("Properties").
		Find(&busses)
	return busses
}

func (repository BusRepository) Update(bus *models.Bus) error {
	if err := repository.Busses.Model(&models.Bus{}).Where("plate_number = ?", bus.PlateNumber).Updates(&bus).Error; err != nil {
		return err
	}
	return nil
}

func (repository BusRepository) Delete(id uint64) error {
	var bus models.Bus
	if err := repository.Busses.Unscoped().Delete(&bus, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository BusRepository) GetWhere(query interface{}, args ...interface{}) *models.Bus {
	var bus models.Bus
	repository.Busses.Where(query, args).
		Preload("BusModel.Brand").
		Preload("BusModel").
		Preload("Type").
		Preload("Properties").
		First(&bus)
	return &bus
}

func (repository BusRepository) BusDefintion() *models.BusDefintion {
	var brands []models.Brand
	var types []models.Type
	var properties []models.Property

	repository.Busses.Find(&brands)
	repository.Busses.Find(&types)
	repository.Busses.Find(&properties)

	var resp = models.BusDefintion{Brand: brands, Type: types, Property: properties}
	return &resp
}
