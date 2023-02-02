package services

import (
	"errors"
	"golang_projects/constants/messages"
	"golang_projects/database/repositories"
	"golang_projects/models"
)

type UserService struct {
	Repository repositories.IUsersRepository
}

type IUserService interface {
	AddUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

func NewUserService(repository repositories.IUsersRepository) UserService {
	return UserService{Repository: repository}
}

func (service UserService) AddUser(user *models.User) error {
	// if checkedUser := service.Repository.GetWhere("email = ?", user.Email); checkedUser.ID != 0 {
	// 	return errors.New(messages.UserAlreadyExists)
	// }

	if err := service.Repository.Add(user); err != nil {
		return err
	}
	return nil
}

func (service UserService) GetUserByEmail(email string) (*models.User, error) {
	user := service.Repository.GetWhere("email = ?", email)

	if user.ID == 0 {
		return nil, errors.New(messages.UserNotFound)
	}

	return user, nil
}
