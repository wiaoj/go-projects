package repositories

import (
	"golang_projects/models"

	"gorm.io/gorm"
)

type ModelRepository struct {
	Models *gorm.DB
}

type IModelRepository interface {
	Add(model *models.Model) error
	Delete(id uint64) error
	GetWhere(query interface{}, args ...interface{}) *models.Model
	GetModelsByBrandId(id uint64) *[]models.Model
}

func NewModelRepository(client *gorm.DB) ModelRepository {
	return ModelRepository{Models: client}
}

func (repository ModelRepository) Add(model *models.Model) error {
	if err := repository.Models.Create(model).Error; err != nil {
		return err
	}
	return nil
}
func (repository ModelRepository) Delete(id uint64) error {
	var model models.Model
	if err := repository.Models.Unscoped().Delete(&model, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository ModelRepository) GetWhere(query interface{}, args ...interface{}) *models.Model {
	var model models.Model
	repository.Models.Where(query, args).First(&model)
	return &model
}

func (repository ModelRepository) GetModelsByBrandId(id uint64) *[]models.Model {
	var models []models.Model
	repository.Models.Find(&models, "brand_id = ?", id)
	return &models
}
