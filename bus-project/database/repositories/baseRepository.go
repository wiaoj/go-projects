package repositories

import (
	"golang_projects/database"
	"golang_projects/models"
)

var DB = database.DB

type modelsType interface {
	*models.User //| *models.TestModel
}

type IBaseRepository[Type modelsType] interface {
	Add(*Type) (*Type, error)
	// Update(*Type) (*Type, error)
	// Delete(*Type) (*Type, error)
	// DeleteById(*Type) (*Type, error)
	// Get(*Type) (*Type, error)
	// GetById(*Type) (*Type, error)
	// GetAll(*Type) (*Type, error)
}

type BaseRepository[Type modelsType] struct{}

// type Repository[Type modelsType] interface {
// 	Create(item Type) error
// 	Read(id int) (Type, error)
// 	Update(item Type) error
// 	Delete(id int) error
// }

func (*BaseRepository[Type]) Add(entity *Type) (*Type, error) {
	if err := DB.Create(entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (*BaseRepository[Type]) DeleteById(entity *Type, entityId uint) (*Type, error) {
	if err := DB.Unscoped().Delete(entity, entityId).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (*BaseRepository[Type]) GetAll(entity *Type) *[]Type {
	var entities *[]Type
	DB.Model(entity).Find(&entities)
	return entities
}
