package services

import (
	"golang_projects/database/repositories"
	"golang_projects/models"
)

type BrandService struct {
	Repository repositories.IBrandRepository
}

type IBrandService interface {
	AddBrand(brand *models.Brand) error
	DeleteBrand(id uint64) error
	GetAllBrands() []models.Brand
}

func NewBrandService(repository repositories.IBrandRepository) BrandService {
	return BrandService{Repository: repository}
}

func (service BrandService) AddBrand(brand *models.Brand) error {
	// if checkedBrand := service.Repository.GetWhere("name = ?", brand.Name); checkedBrand.ID != 0 {
	// 	return errors.New(messages.BrandAlreadyExists)
	// }

	if err := service.Repository.Add(brand); err != nil {
		return err
	}
	return nil
}

func (service BrandService) DeleteBrand(id uint64) error {
	if err := service.Repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (service BrandService) GetAllBrands() []models.Brand {
	return *service.Repository.GetAll()
}
