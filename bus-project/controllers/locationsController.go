package controllers

import (
	"golang_projects/contracts"
	response "golang_projects/contracts/response"
	"golang_projects/services"

	"github.com/gofiber/fiber/v2"
)

type LocationsController struct {
	LocationService services.ILocationService
}

func (controller LocationsController) GetAll(context *fiber.Ctx) error {

	var responseLocations []contracts.GetAllLocationsResponse

	var locations = controller.LocationService.GetAllLocations()

	for i := 0; i < len(locations); i++ {
		responseLocations = append(responseLocations, contracts.GetAllLocationsResponse{
			Id:   locations[i].ID,
			Name: locations[i].Name,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.ListDataResponse{
		//Message: messages.GetLocations,
		Status: response.SuccessStatus,
		Datas:  responseLocations,
	})
}
