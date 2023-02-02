package services

import (
	"golang_projects/database/repositories"
	"golang_projects/models"
)

type ModelService struct {
	Repository repositories.IModelRepository
}

type IModelService interface {
	AddModel(Model *models.Model) error
	DeleteModel(id uint64) error
	GetModelsByBrandId(id uint64) []models.Model
}

func NewModelService(repository repositories.IModelRepository) ModelService {
	return ModelService{Repository: repository}
}

func (service ModelService) AddModel(model *models.Model) error {
	if err := service.Repository.Add(model); err != nil {
		return err
	}
	return nil
}

func (service ModelService) DeleteModel(id uint64) error {
	if err := service.Repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (service ModelService) GetModelsByBrandId(id uint64) []models.Model {
	return *service.Repository.GetModelsByBrandId(id)
}
