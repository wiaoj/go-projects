package repositories

import (
	"golang_projects/models"

	"gorm.io/gorm"
)

type BrandRepository struct {
	Brands *gorm.DB
}

type IBrandRepository interface {
	Add(brand *models.Brand) error
	Delete(id uint64) error
	GetWhere(query interface{}, args ...interface{}) *models.Brand
	GetAll() *[]models.Brand
}

func NewBrandRepository(client *gorm.DB) BrandRepository {
	return BrandRepository{Brands: client}
}

func (repository BrandRepository) Add(brand *models.Brand) error {
	if err := repository.Brands.Create(brand).Error; err != nil {
		return err
	}
	return nil
}
func (repository BrandRepository) Delete(id uint64) error {
	var brand models.Brand
	if err := repository.Brands.Unscoped().Delete(&brand, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository BrandRepository) GetWhere(query interface{}, args ...interface{}) *models.Brand {
	var brand models.Brand
	repository.Brands.Where(query, args).First(&brand)
	return &brand
}

func (repository BrandRepository) GetAll() *[]models.Brand {
	var brands []models.Brand
	repository.Brands.Find(&brands)
	return &brands
}
