package repositories

import (
	"golang_projects/models"

	"gorm.io/gorm"
)

// func AddUser(user *models.UserModel) (*models.UserModel, error) {

//		if err := database.DB.Create(user).Error; err != nil {
//			return nil, err
//		}
//		asdv, ok := asd.(Adder)
//		asdv.Add()
//		return user, nil
//	}

type UsersRepository struct {
	Users *gorm.DB
}

type IUsersRepository interface {
	Add(user *models.User) error
	GetWhere(query interface{}, args ...interface{}) *models.User
}

func NewUserRepository(client *gorm.DB) UsersRepository {
	return UsersRepository{Users: client}
}

func (repository UsersRepository) Add(user *models.User) error {
	if err := repository.Users.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (repository UsersRepository) GetWhere(query interface{}, args ...interface{}) *models.User {
	var user models.User
	repository.Users.Where(query, args).First(&user)
	return &user
}

// func (u *UserRepository) Add(user *models.User) error {
// 	var us models.User
// 	DB.Where("email = ?", user.Email).First(&us)
// 	if us.ID != 0 {
// 		return errors.New("Böyle bir email zaten kayıtlı")
// 	}

// 	if _, err := u.Base.Add(&user); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u UserRepository) Update(user *models.UserModel) (*models.UserModel, error) {

// 	return user, nil
// }
// func (u UserRepository) Delete(user *models.UserModel) (*models.UserModel, error) {

// 	return user, nil
// }
// func (u UserRepository) Get(user *models.UserModel) (*models.UserModel, error) {

// 	return user, nil
// }
// func (u UserRepository) GetAll(user *models.UserModel) (*models.UserModel, error) {

// 	return user, nil
// }
