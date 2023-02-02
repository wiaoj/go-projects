package services

import (
	"golang_projects/database/repositories"
	"golang_projects/models"
)

type TypeService struct {
	Repository repositories.ITypeRepository
}

type ITypeService interface {
	AddType(Type *models.Type) error
	DeleteType(id uint64) error
	GetAllTypes() []models.Type
}

func NewTypeService(repository repositories.ITypeRepository) TypeService {
	return TypeService{Repository: repository}
}

func (service TypeService) AddType(Type *models.Type) error {
	// if checkedType := service.Repository.GetWhere("value = ?", Type.Value); checkedType.ID != 0 {
	// 	return errors.New(messages.TypeAlreadyExists)
	// }

	if err := service.Repository.Add(Type); err != nil {
		return err
	}
	return nil
}

func (service TypeService) DeleteType(id uint64) error {
	if err := service.Repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (service TypeService) GetAllTypes() []models.Type {
	return *service.Repository.GetAll()
}
