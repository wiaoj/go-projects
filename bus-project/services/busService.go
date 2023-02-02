package services

import (
	"golang_projects/database/repositories"
	"golang_projects/models"
)

type BusService struct {
	Repository repositories.IBusRepository
}

type IBusService interface {
	GetBusDefinition() *models.BusDefintion
	AddBus(Bus *models.Bus) error
	UpdateBus(Bus *models.Bus) error
	GetAllBusses() []*models.Bus
	DeleteBus(id uint64) error
	GetById(id uint64) *models.Bus
}

func NewBusService(repository repositories.IBusRepository) BusService {
	return BusService{Repository: repository}
}

func (service BusService) AddBus(Bus *models.Bus) error {
	if err := service.Repository.Add(Bus); err != nil {
		return err
	}
	return nil
}

func (service BusService) GetAllBusses() []*models.Bus {
	return service.Repository.GetAll()
}

func (service BusService) GetById(id uint64) *models.Bus {
	return service.Repository.GetWhere("id = ?", id)
}

func (service BusService) DeleteBus(id uint64) error {
	if err := service.Repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (service BusService) UpdateBus(Bus *models.Bus) error {
	if err := service.Repository.Update(Bus); err != nil {
		return err
	}
	return nil
}

func (service BusService) GetBusDefinition() *models.BusDefintion {
	return service.Repository.BusDefintion()
}
