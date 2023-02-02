package services

import (
	"golang_projects/database/repositories"
	"golang_projects/models"
)

type PropertyService struct {
	Repository repositories.IPropertyRepository
}

type IPropertyService interface {
	AddProperty(property *models.Property) error
	GetPropertiesByIds(ids []uint) *[]models.Property
	DeleteProperty(id uint64) error
	GetAllProperties() []models.Property
}

func NewPropertyService(repository repositories.IPropertyRepository) PropertyService {
	return PropertyService{Repository: repository}
}

func (service PropertyService) AddProperty(property *models.Property) error {
	// if checkedProperty := service.Repository.GetWhere("value = ?", property.Value); checkedProperty.ID != 0 {
	// 	return errors.New(messages.PropertyAlreadyExists)
	// }

	if err := service.Repository.Add(property); err != nil {
		return err
	}
	return nil
}

func (service PropertyService) GetPropertiesByIds(ids []uint) *[]models.Property {
	return service.Repository.GetPropertiesByIds(ids)
}

func (service PropertyService) DeleteProperty(id uint64) error {
	if err := service.Repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (service PropertyService) GetAllProperties() []models.Property {
	return *service.Repository.GetAll()
}
