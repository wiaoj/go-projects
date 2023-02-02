package controllers

import (
	"golang_projects/constants/messages"
	"golang_projects/contracts"
	response "golang_projects/contracts/response"
	"golang_projects/models"
	"golang_projects/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TypesController struct {
	TypeService services.ITypeService
}

func (controller TypesController) CreateType(context *fiber.Ctx) error {

	request := new(contracts.CreateTypeRequest)

	if err := context.BodyParser(&request); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	Type := models.Type{
		Value: request.Value,
	}

	if err := controller.TypeService.AddType(&Type); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}
	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.CreatedType,
		Status:  response.SuccessStatus,
		Data: contracts.CreatedTypeResponse{
			Id:    Type.ID,
			Value: Type.Value,
		},
	})

}

func (controller TypesController) DeleteType(context *fiber.Ctx) error {
	paramsId := context.Params("id")

	id, err := strconv.ParseUint(paramsId, 10, 32)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	if err := controller.TypeService.DeleteType(id); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.DeletedType,
		Status:  response.SuccessStatus,
		Data: contracts.DeletedTypeResponse{
			Id: uint(id),
		},
	})
}

func (controller TypesController) GetAllTypes(context *fiber.Ctx) error {
	types := controller.TypeService.GetAllTypes()

	if types == nil {

	}

	var respons []contracts.TypeResponse
	for i := 0; i < len(types); i++ {
		respons = append(respons, contracts.TypeResponse{
			Id:    types[i].ID,
			Value: types[i].Value,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.ListDataResponse{
		Status: response.SuccessStatus,
		Datas:  respons,
	})
}
