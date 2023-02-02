package repositories

import (
	"golang_projects/models"

	"gorm.io/gorm"
)

type PropertyRepository struct {
	Properties *gorm.DB
}

type IPropertyRepository interface {
	Add(property *models.Property) error
	Delete(id uint64) error
	GetWhere(query interface{}, args ...interface{}) *models.Property
	GetPropertiesByIds(ids []uint) *[]models.Property
	GetAll() *[]models.Property
}

func NewPropertyRepository(client *gorm.DB) PropertyRepository {
	return PropertyRepository{Properties: client}
}

func (repository PropertyRepository) Add(property *models.Property) error {
	if err := repository.Properties.Create(property).Error; err != nil {
		return err
	}
	return nil
}
func (repository PropertyRepository) Delete(id uint64) error {
	var property models.Property
	if err := repository.Properties.Unscoped().Delete(&property, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository PropertyRepository) GetWhere(query interface{}, args ...interface{}) *models.Property {
	var property models.Property
	repository.Properties.Where(query, args).First(&property)
	return &property
}

func (repository PropertyRepository) GetPropertiesByIds(ids []uint) *[]models.Property {
	var properties []models.Property
	repository.Properties.Find(&properties, ids)
	return &properties
}

func (repository PropertyRepository) GetAll() *[]models.Property {
	var properties []models.Property
	repository.Properties.Find(&properties)
	return &properties
}
