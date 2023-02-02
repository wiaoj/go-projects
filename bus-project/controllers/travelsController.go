package controllers

import (
	"golang_projects/constants/messages"
	contracts "golang_projects/contracts"
	jsonResponse "golang_projects/contracts/response"
	response "golang_projects/contracts/response"
	"golang_projects/models"
	"golang_projects/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TravelsController struct {
	TravelService services.ITravelService
	BusService    services.IBusService
	SeatService   services.ISeatService
}

func (controllers TravelsController) CreateTravel(context *fiber.Ctx) error {
	request := new(contracts.CreateTravelRequest)

	if err := context.BodyParser(&request); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	var busses []models.Bus
	bus := new(models.Bus)
	bus.ID = request.BusId
	busses = append(busses, *bus)

	if err := controllers.TravelService.AddTravel(&models.Travel{
		Fee:          request.Fee,
		FromLocation: request.From,
		ToLocation:   request.To,
		Day:          request.Date,
		Buses:        busses,
	}); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	return context.JSON(response.SingleResponse{
		Message: messages.CreateVoyageSuccess,
		Status:  response.SuccessStatus,
	})
}

func (controllers TravelsController) GetAllTravel(context *fiber.Ctx) error {
	travels := controllers.TravelService.GetAllTravels()

	response := []contracts.GetTravelResponse{}
	for i := 0; i < len(travels); i++ {
		travel := travels[i]

		busses := []contracts.GetTravelBusDTO{}

		for i := 0; i < len(travel.Buses); i++ {

			bus := travel.Buses[i]

			properties := []contracts.GetTravelBusPropertyDTO{}
			for i := 0; i < len(bus.Properties); i++ {
				property := bus.Properties[i]
				properties = append(properties, contracts.GetTravelBusPropertyDTO{
					Id:    property.ID,
					Value: property.Value,
				})
			}

			seats := []contracts.GetTravelSeatDTO{}
			for i := 0; i < len(bus.Seats); i++ {
				seat := bus.Seats[i]

				seatProperties := []contracts.GetTravelSeatPropertyDTO{}
				for i := 0; i < len(seat.Properties); i++ {
					seatProperty := seat.Properties[i]
					seatProperties = append(seatProperties, contracts.GetTravelSeatPropertyDTO{
						Id:     seatProperty.ID,
						No:     seatProperty.No,
						Gender: seatProperty.Gender,
					})
				}

				seats = append(seats, contracts.GetTravelSeatDTO{
					Id:         seat.ID,
					Count:      seat.Count,
					Properties: seatProperties,
				})
			}

			busses = append(busses, contracts.GetTravelBusDTO{
				Id:          bus.ID,
				PlateNumber: bus.PlateNumber,
				BusBrand:    bus.BusModel.Brand.Name,
				BusModel:    bus.BusModel.Value,
				Type:        bus.Type.Value,
				Properties:  properties,
				Seats:       seats,
			})
		}

		response = append(response, contracts.GetTravelResponse{
			Id:    travel.ID,
			Fee:   travel.Fee,
			From:  travel.FromLocation,
			To:    travel.ToLocation,
			Day:   travel.Day,
			Time:  travel.Time,
			Buses: busses,
		})
	}
	return context.Status(fiber.StatusOK).JSON(jsonResponse.ListDataResponse{
		Status: jsonResponse.SuccessStatus,
		Datas:  response,
	})
}

func (controllers TravelsController) GetTravelQuery(context *fiber.Ctx) error {
	from := context.Params("from")
	to := context.Params("to")
	day := context.Params("day")
	at := context.Params("at")

	day1, _ := time.Parse("21-11-2000", day)
	at1, _ := time.Parse("00:00:00", at)
	travels := controllers.TravelService.GetTravel(from, to, day1, at1)

	if len(travels) == 0 {
		return context.Status(fiber.StatusNotFound).JSON(jsonResponse.SingleResponse{
			Message: messages.TravelNotFound,
			Status:  jsonResponse.ErrorStatus,
		})
	}

	response := []contracts.GetTravelResponse{}
	for i := 0; i < len(travels); i++ {
		travel := travels[i]

		busses := []contracts.GetTravelBusDTO{}

		for i := 0; i < len(travel.Buses); i++ {

			bus := travel.Buses[i]

			properties := []contracts.GetTravelBusPropertyDTO{}
			for i := 0; i < len(bus.Properties); i++ {
				property := bus.Properties[i]
				properties = append(properties, contracts.GetTravelBusPropertyDTO{
					Id:    property.ID,
					Value: property.Value,
				})
			}

			seats := []contracts.GetTravelSeatDTO{}
			for i := 0; i < len(bus.Seats); i++ {
				seat := bus.Seats[i]

				seatProperties := []contracts.GetTravelSeatPropertyDTO{}
				for i := 0; i < len(seat.Properties); i++ {
					seatProperty := seat.Properties[i]
					seatProperties = append(seatProperties, contracts.GetTravelSeatPropertyDTO{
						Id:     seatProperty.ID,
						No:     seatProperty.No,
						Gender: seatProperty.Gender,
					})
				}

				seats = append(seats, contracts.GetTravelSeatDTO{
					Id:         seat.ID,
					Count:      seat.Count,
					Properties: seatProperties,
				})
			}

			busses = append(busses, contracts.GetTravelBusDTO{
				Id:          bus.ID,
				PlateNumber: bus.PlateNumber,
				BusBrand:    bus.BusModel.Brand.Name,
				BusModel:    bus.BusModel.Value,
				Type:        bus.Type.Value,
				Properties:  properties,
				Seats:       seats,
			})
		}

		response = append(response, contracts.GetTravelResponse{
			Id:    travel.ID,
			Fee:   travel.Fee,
			From:  travel.FromLocation,
			To:    travel.ToLocation,
			Day:   travel.Day,
			Time:  travel.Time,
			Buses: busses,
		})
	}
	return context.Status(fiber.StatusOK).JSON(jsonResponse.ListDataResponse{
		Status: jsonResponse.SuccessStatus,
		Datas:  response,
	})
}

func (controllers TravelsController) BuyTicket(context *fiber.Ctx) error {
	request := new(contracts.BuyTicketRequest)

	if err := context.BodyParser(&request); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
			Message: err.Error(),
			Status:  response.ErrorStatus,
		})
	}

	travel := controllers.TravelService.GetTravelById(uint64(request.TravelId))

	if travel.ID == 0 {
		return context.Status(fiber.StatusNotFound).JSON(response.SingleResponse{
			Message: messages.TravelNotFound,
			Status:  response.ErrorStatus,
		})
	}

	if bus := controllers.BusService.GetById(uint64(request.BusId)); bus.ID == 0 {
		return context.Status(fiber.StatusNotFound).JSON(response.SingleResponse{
			Message: messages.BusNotFound,
			Status:  response.ErrorStatus,
		})
	}

	for i := 0; i < len(travel.Buses); i++ {
		bus := travel.Buses[i]

		if bus.SeatsCount < request.No {
			return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
				Message: messages.InvalidSeatPropertyNo,
				Status:  response.ErrorStatus,
			})
		}

		if bus.ID == request.BusId {
			seats := bus.Seats

			for i := 0; i < len(seats); i++ {
				seatProperties := seats[i].Properties

				for i := 0; i < len(seatProperties); i++ {
					property := seatProperties[i]
					if property.No == request.No {
						return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
							Message: messages.SeatPropertyAlredyBuyed,
							Status:  response.ErrorStatus,
						})
					}
				}
				var seatProperty = models.SeatProperty{
					SeatId: seats[i].ID,
					No:     request.No,
					Gender: request.Gender,
				}
				if err := controllers.SeatService.AddSeatProperty(&seatProperty); err != nil {
					return context.Status(fiber.StatusInternalServerError).JSON(response.SingleResponse{
						Message: err.Error(),
						Status:  response.ErrorStatus,
					})

				}
			}
		}
	}

	// if travel.Buses == nil {
	// 	return context.Status(fiber.StatusBadRequest).JSON(response.SingleResponse{
	// 		Message: "otobüs ve sefer bulunamadı",
	// 		Status:  response.ErrorStatus,
	// 	})
	// }

	return context.Status(fiber.StatusOK).JSON(response.SingleResponse{
		Message: messages.BuyTicketSuccess,
		Status:  response.SuccessStatus,
	})
}

func (controllers TravelsController) GetData(context *fiber.Ctx) error {

	return context.JSON("")
}
