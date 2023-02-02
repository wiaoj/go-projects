package repositories

import (
	"golang_projects/models"

	"gorm.io/gorm"
)

type TypeRepository struct {
	Types *gorm.DB
}

type ITypeRepository interface {
	Add(Type *models.Type) error
	Delete(id uint64) error
	GetWhere(query interface{}, args ...interface{}) *models.Type
	GetAll() *[]models.Type
}

func NewTypeRepository(client *gorm.DB) TypeRepository {
	return TypeRepository{Types: client}
}

func (repository TypeRepository) Add(Type *models.Type) error {
	if err := repository.Types.Create(Type).Error; err != nil {
		return err
	}
	return nil
}
func (repository TypeRepository) Delete(id uint64) error {
	var Type models.Type
	if err := repository.Types.Unscoped().Delete(&Type, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository TypeRepository) GetWhere(query interface{}, args ...interface{}) *models.Type {
	var Type models.Type
	repository.Types.Where(query, args).First(&Type)
	return &Type
}
func (repository TypeRepository) GetAll() *[]models.Type {
	var Types []models.Type
	repository.Types.First(&Types)
	return &Types
}
