package main

import (
	"golang_projects/config"
	"golang_projects/controllers"
	"golang_projects/database"
	"golang_projects/database/repositories"
	"golang_projects/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

// func init() {
// 	database.ConnectToDatabase()
// }

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	dbClient := database.ConnectToDatabase()
	database.MigrateDatabase()

	userRepository := repositories.NewUserRepository(dbClient)

	authorizationController := controllers.AuthorizationController{
		UserService: services.NewUserService(userRepository),
	}
	api := app.Group("/api")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	auth := api.Group("/auth")
	{
		auth.Post("/register", authorizationController.Register)
		auth.Post("/login", authorizationController.Login)
	}

	addBrandRoutes(dbClient, api)
	addModelRoutes(dbClient, api)
	addTypeRoutes(dbClient, api)
	addPropertyRoutes(dbClient, api)
	addBusRoutes(dbClient, api)
	addLocationsRoutes(dbClient, api)
	addBusTravelRoutes(dbClient, api)
	app.Listen(config.GetPort())
}

func addBrandRoutes(dbClient *gorm.DB, api fiber.Router) {
	brandRepository := repositories.NewBrandRepository(dbClient)
	brandsController := controllers.BrandsController{
		BrandService: services.NewBrandService(brandRepository),
	}
	brands := api.Group("/brands")
	{
		brands.Get("/", brandsController.GetAllBrands)

		brands.Post("/", brandsController.CreateBrand)
		brands.Delete("/:id", brandsController.DeleteBrand)
	}
}

func addModelRoutes(dbClient *gorm.DB, api fiber.Router) {
	modelRepository := repositories.NewModelRepository(dbClient)
	modelsController := controllers.ModelsController{
		ModelService: services.NewModelService(modelRepository),
	}
	models := api.Group("/models")
	{
		models.Get("/", modelsController.GetModelByBrandId)

		models.Post("/", modelsController.CreateModel)
		models.Delete("/:id", modelsController.DeleteModel)
	}
}

func addTypeRoutes(dbClient *gorm.DB, api fiber.Router) {
	typeRepository := repositories.NewTypeRepository(dbClient)
	typesController := controllers.TypesController{
		TypeService: services.NewTypeService(typeRepository),
	}
	types := api.Group("/types")
	{
		types.Get("/", typesController.GetAllTypes)

		types.Post("/", typesController.CreateType)
		types.Delete("/:id", typesController.DeleteType)
	}
}

func addPropertyRoutes(dbClient *gorm.DB, api fiber.Router) {
	propertyRepository := repositories.NewPropertyRepository(dbClient)
	propertiesController := controllers.PropertiesController{
		PropertyService: services.NewPropertyService(propertyRepository),
	}
	properties := api.Group("/properties")
	{
		properties.Get("/", propertiesController.GetAllProperties)

		properties.Post("/", propertiesController.CreateProperty)
		properties.Delete("/:id", propertiesController.DeleteProperty)
	}
}

func addBusRoutes(dbClient *gorm.DB, api fiber.Router) {
	busRepository := repositories.NewBusRepository(dbClient)
	propertyRepository := repositories.NewPropertyRepository(dbClient)
	bussesController := controllers.BussesController{
		BussesService:     services.NewBusService(busRepository),
		PropertiesService: services.NewPropertyService(propertyRepository),
	}
	busses := api.Group("/busses")
	{
		busses.Get("/", bussesController.GetAllBusses)
		busses.Get("/bus-definition", bussesController.GetBusDefinition)
		busses.Get("/:id", bussesController.GetById)
		busses.Post("/", bussesController.CreateBus)
		busses.Put("/", bussesController.UpdateBus)
		busses.Delete("/:id", bussesController.DeleteBus)
	}
}

func addLocationsRoutes(dbClient *gorm.DB, api fiber.Router) {
	locationRepository := repositories.NewLocationRepository(dbClient)
	locationController := controllers.LocationsController{
		LocationService: services.NewLocationService(locationRepository),
	}

	busses := api.Group("/locations")
	{
		busses.Get("/", locationController.GetAll)
	}
}

func addBusTravelRoutes(dbClient *gorm.DB, api fiber.Router) {
	travelRepository := repositories.NewTravelRepository(dbClient)
	travelsController := controllers.TravelsController{
		TravelService: services.NewTravelService(travelRepository),
		BusService:    services.NewBusService(repositories.NewBusRepository(dbClient)),
		SeatService:   services.NewSeatService(repositories.NewSeatRepository(dbClient)),
	}
	busses := api.Group("/voyage")
	{
		busses.Get("/:from-:to/day::day-time::at", travelsController.GetTravelQuery)
		busses.Get("/", travelsController.GetAllTravel)

		busses.Post("/", travelsController.CreateTravel)
		busses.Post("/buy-ticket", travelsController.BuyTicket)
	}
}
