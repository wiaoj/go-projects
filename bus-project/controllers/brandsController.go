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

type BrandsController struct {
	BrandService services.IBrandService
}

func (controller BrandsController) CreateBrand(context *fiber.Ctx) error {

	request := new(contracts.CreateBrandRequest)

	if err := context.BodyParser(&request); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	brand := models.Brand{
		Name: request.Name,
	}

	if err := controller.BrandService.AddBrand(&brand); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}
	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.CreatedBrand,
		Status:  response.SuccessStatus,
		Data: contracts.CreatedBrandResponse{
			Id:   brand.ID,
			Name: brand.Name,
		},
	})

}

func (controller BrandsController) DeleteBrand(context *fiber.Ctx) error {
	paramsId := context.Params("id")

	id, err := strconv.ParseUint(paramsId, 10, 32)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	if err := controller.BrandService.DeleteBrand(id); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.DeletedBrand,
		Status:  response.SuccessStatus,
		Data: contracts.DeletedBrandResponse{
			Id: uint(id),
		},
	})
}

func (controller BrandsController) GetAllBrands(context *fiber.Ctx) error {
	brands := controller.BrandService.GetAllBrands()

	if brands == nil {

	}

	var respons []contracts.BrandResponse
	for i := 0; i < len(brands); i++ {
		respons = append(respons, contracts.BrandResponse{
			Id:   brands[i].ID,
			Name: brands[i].Name,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.ListDataResponse{
		Status: response.SuccessStatus,
		Datas:  respons,
	})
}
