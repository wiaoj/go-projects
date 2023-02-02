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

type PropertiesController struct {
	PropertyService services.IPropertyService
}

func (controller PropertiesController) CreateProperty(context *fiber.Ctx) error {

	request := new(contracts.CreateTypeRequest)

	if err := context.BodyParser(&request); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	property := models.Property{
		Value: request.Value,
	}

	if err := controller.PropertyService.AddProperty(&property); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}
	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.CreatedProperty,
		Status:  response.SuccessStatus,
		Data: contracts.CreatedPropertyResponse{
			Id:    property.ID,
			Value: property.Value,
		},
	})

}

func (controller PropertiesController) DeleteProperty(context *fiber.Ctx) error {
	paramsId := context.Params("id")

	id, err := strconv.ParseUint(paramsId, 10, 32)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	if err := controller.PropertyService.DeleteProperty(id); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.DeletedProperty,
		Status:  response.SuccessStatus,
		Data: contracts.DeletedPropertyResponse{
			Id: uint(id),
		},
	})
}

func (controller PropertiesController) GetAllProperties(context *fiber.Ctx) error {
	properties := controller.PropertyService.GetAllProperties()

	if properties == nil {

	}

	var respons []contracts.PropertyResponse
	for i := 0; i < len(properties); i++ {
		respons = append(respons, contracts.PropertyResponse{
			Id:    properties[i].ID,
			Value: properties[i].Value,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.ListDataResponse{
		Status: response.SuccessStatus,
		Datas:  respons,
	})
}
