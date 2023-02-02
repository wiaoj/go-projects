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

type BussesController struct {
	BussesService     services.IBusService
	PropertiesService services.IPropertyService
}

func (controller BussesController) GetById(context *fiber.Ctx) error {
	paramsId := context.Params("id")

	id, err := strconv.ParseUint(paramsId, 10, 32)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	bus := controller.BussesService.GetById(id)

	if bus.ID == 0 {
		return context.Status(fiber.StatusNotFound).JSON(response.SingleResponse{
			Message: messages.BusNotFound,
			Status:  response.ErrorStatus,
		})
	}

	var properties []contracts.BusPropertyResponse

	for i := 0; i < len(bus.Properties); i++ {
		properties = append(properties, contracts.BusPropertyResponse{
			Id:    bus.Properties[i].ID,
			Value: bus.Properties[i].Value,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.GetBus,
		Status:  response.SuccessStatus,
		Data: contracts.GetByIdBusResponse{
			Id:          bus.ID,
			PlateNumber: bus.PlateNumber,
			Seats:       bus.SeatsCount,
			BusBrand:    bus.BusModel.Brand.Name,
			BusModel:    bus.BusModel.Value,
			Type:        bus.Type.Value,
			Properties:  properties,
		},
	})
}

func (controller BussesController) CreateBus(context *fiber.Ctx) error {

	request := new(contracts.CreateBusRequest)

	if err := context.BodyParser(&request); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	var propertyIds []uint

	for i := 0; i < len(request.Properties); i++ {
		propertyIds = append(propertyIds, request.Properties[i].Id)
	}

	properties := controller.PropertiesService.GetPropertiesByIds(propertyIds)

	bus := models.Bus{

		PlateNumber: request.PlateNumber,
		SeatsCount:  request.SeatsCount,
		BusModelId:  request.BusModelId,
		TypeId:      request.TypeId,
		Properties:  *properties,
	}

	if err := controller.BussesService.AddBus(&bus); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}
	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.CreatedBus,
		Status:  response.SuccessStatus,
		Data: contracts.CreatedBusResponse{
			Id: bus.ID,
			// PlateNumber: bus.PlateNumber,
			// Seats:       bus.Seats,
			// BusBrand:    bus.BusModel.Brand.Name,
			// BusModel:    bus.BusModel.Value,
			// Type:        bus.Type.Value,
			// Properties:  bus.Properties,
		},
	})

}

func (controller BussesController) UpdateBus(context *fiber.Ctx) error {

	request := new(contracts.CreateBusRequest)

	if err := context.BodyParser(&request); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	var propertyIds []uint

	for i := 0; i < len(request.Properties); i++ {
		propertyIds = append(propertyIds, request.Properties[i].Id)
	}

	properties := controller.PropertiesService.GetPropertiesByIds(propertyIds)

	bus := models.Bus{

		PlateNumber: request.PlateNumber,
		SeatsCount:  request.SeatsCount,
		BusModelId:  request.BusModelId,
		TypeId:      request.TypeId,
		Properties:  *properties,
	}

	if err := controller.BussesService.UpdateBus(&bus); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}
	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.CreatedBus,
		Status:  response.SuccessStatus,
	})

}

func (controller BussesController) DeleteBus(context *fiber.Ctx) error {
	paramsId := context.Params("id")

	id, err := strconv.ParseUint(paramsId, 10, 32)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	if err := controller.BussesService.DeleteBus(id); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.DeletedBus,
		Status:  response.SuccessStatus,
		Data: contracts.DeletedBrandResponse{
			Id: uint(id),
		},
	})
}

func (controller BussesController) GetBusDefinition(context *fiber.Ctx) error {
	resp := controller.BussesService.GetBusDefinition()
	return context.Status(fiber.StatusOK).JSON(response.SingleDataResponse{
		Message: messages.GetBus,
		Status:  response.SuccessStatus,
		Data:    resp,
	})
}

func (controller BussesController) GetAllBusses(context *fiber.Ctx) error {
	busses := controller.BussesService.GetAllBusses()

	if busses == nil {

	}

	var respons []contracts.ListBusResponse
	for i := 0; i < len(busses); i++ {
		respons = append(respons, contracts.ListBusResponse{
			Id:          busses[i].ID,
			PlateNumber: busses[i].PlateNumber,
			Seats:       busses[i].SeatsCount,
			BusBrand:    busses[i].BusModel.Brand.Name,
			BusModel:    busses[i].BusModel.Value,
			Type:        busses[i].Type.Value,
			Properties:  busses[i].Properties,
		})
	}

	return context.Status(fiber.StatusOK).JSON(response.ListDataResponse{
		Status: response.SuccessStatus,
		Datas:  respons,
	})
}
