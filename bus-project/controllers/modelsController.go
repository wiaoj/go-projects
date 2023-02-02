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

type ModelsController struct {
	ModelService services.IModelService
}

func (controller ModelsController) CreateModel(context *fiber.Ctx) error {

	request := new(contracts.CreateModelRequest)

	if err := context.BodyParser(&request); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	model := models.Model{
		BrandId: request.BrandId,
		Value:   request.Value,
	}

	if err := controller.ModelService.AddModel(&model); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}
	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.CreatedModel,
		Status:  response.SuccessStatus,
		Data: contracts.CreatedTypeResponse{
			Id:    model.ID,
			Value: model.Value,
		},
	})

}

func (controller ModelsController) DeleteModel(context *fiber.Ctx) error {
	paramsId := context.Params("id")

	id, err := strconv.ParseUint(paramsId, 10, 32)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	if err := controller.ModelService.DeleteModel(id); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.DeletedModel,
		Status:  response.SuccessStatus,
		Data: contracts.DeletedTypeResponse{
			Id: uint(id),
		},
	})
}

func (controller ModelsController) GetModelByBrandId(context *fiber.Ctx) error {
	paramsId := context.Query("brandId")

	id, err := strconv.ParseUint(paramsId, 10, 32)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}
	models := controller.ModelService.GetModelsByBrandId(id)

	if models == nil {

	}

	var respons []contracts.ModelResponse
	for i := 0; i < len(models); i++ {
		respons = append(respons, contracts.ModelResponse{
			Id:      models[i].ID,
			BrandId: models[i].BrandId,
			Value:   models[i].Value,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.ListDataResponse{
		Status: response.SuccessStatus,
		Datas:  respons,
	})
}
