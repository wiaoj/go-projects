package controllers

import (
	"golang_projects/constants/messages"
	"golang_projects/contracts"
	response "golang_projects/contracts/response"
	"golang_projects/models"
	"golang_projects/services"
	encryption "golang_projects/utils/encryption"
	utils "golang_projects/utils/encryption"
	jwt "golang_projects/utils/jwt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type AuthorizationController struct {
	UserService services.IUserService
}

var validate = validator.New()

func (controller AuthorizationController) Login(context *fiber.Ctx) error {
	request := new(contracts.UserLoginRequest)

	if err := context.BodyParser(&request); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	user, err := controller.UserService.GetUserByEmail(request.Email)

	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	if err := utils.CheckPassword(user.PasswordHash, request.Password); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: messages.WrongPassword,
			Status:  response.ErrorStatus,
		})
	}

	token, err := jwt.GenerateJwtToken(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.LoginSuccessfully,
		Status:  response.SuccessStatus,
		Data: contracts.UserLoginResponse{
			Token: token,
		},
	})
}

func (controller AuthorizationController) Register(context *fiber.Ctx) error {
	request := new(contracts.UserRegisterRequest)

	if err := context.BodyParser(&request); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	if err := validate.Struct(request); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	user := models.User{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Email:        request.Email,
		PasswordHash: encryption.HashPassword(request.Password), //TODO: password nil gelebilir
	}

	//var repo repositories.UserRepository

	if err := controller.UserService.AddUser(&user); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	token, err := jwt.GenerateJwtToken(&user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.RegisterSuccessfully,
		Status:  response.SuccessStatus,
		Data: contracts.UserLoginResponse{
			Token: token,
		},
	})
}
