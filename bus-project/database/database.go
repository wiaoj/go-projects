package database

import (
	"fmt"
	"golang_projects/config"
	"golang_projects/constants/messages"
	"golang_projects/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB = ConnectToDatabase()

func ConnectToDatabase() *gorm.DB {
	//var err error
	dsn := config.GetConnectionString()

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(messages.FailedConnectingDatabase)
		return nil
	}

	fmt.Println(messages.SuccessConnectingDatabase)
	return client
}

func MigrateDatabase() string {
	migrate()
	if tables, _ := DB.Migrator().GetTables(); len(tables) == 0 {
		//x := initializers.DB.Migrator().HasTable("users")
		return "veritabanı oluşturuldu"
	}
	return "veritabanı mevcut"
}

const upSchema = "Create SCHEMA public"
const dropSchema = "DROP SCHEMA IF EXISTS public CASCADE;"

func migrate() {
	if err := DB.Exec(dropSchema).Error; err != nil {
		panic(err)
	}

	if err := DB.Exec(upSchema).Error; err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Brand{})
	DB.AutoMigrate(&models.Model{})
	DB.AutoMigrate(&models.Type{})
	DB.AutoMigrate(&models.Property{})

	DB.AutoMigrate(&models.Location{})

	DB.AutoMigrate(&models.SeatProperty{})
	DB.AutoMigrate(&models.Seat{})

	DB.AutoMigrate(&models.Bus{})

	DB.AutoMigrate(&models.Travel{})

	//DB.AutoMigrate(&models.BusTravel{})
	// database.AutoMigrate(&models.Claim{})
	// database.AutoMigrate(&models.UsersClaims{})

	seedBrands()
	seedModels()
	seedProperties()
	seedTypes()

	seedLocations()

	seedBusses()

	seedTravels()
	seedSeats()
}

func seedBrands() {
	var brands []models.Brand

	brands = append(brands,
		models.Brand{
			Name: "Mercedes",
		},
		models.Brand{
			Name: "Volvo",
		})

	DB.Create(&brands)
}

func seedModels() {
	var _models []models.Model

	_models = append(_models,
		models.Model{
			BrandId: 1,
			Value:   "Serie A",
		},
		models.Model{
			BrandId: 2,
			Value:   "Serie B",
		})

	DB.Create(&_models)
}

func seedTypes() {
	var types []models.Type

	types = append(types,
		models.Type{
			Value: "2+1",
		},
		models.Type{
			Value: "2+2",
		})

	DB.Create(&types)
}

var properties []models.Property

func seedProperties() {

	properties = append(properties,
		models.Property{
			Value: "tv",
		},
		models.Property{
			Value: "ikram",
		})

	DB.Create(&properties)
}

var busses1 []models.Bus
var busses2 []models.Bus

func seedBusses() {

	busses1 = append(busses1, models.Bus{
		PlateNumber: "FX-56871",
		SeatsCount:  50,
		BusModelId:  1,
		TypeId:      1,
		Properties:  properties,
	},
		models.Bus{
			PlateNumber: "FX-56811",
			SeatsCount:  50,
			BusModelId:  1,
			TypeId:      2,
			Properties:  properties,
		})

	busses2 = append(busses2, models.Bus{
		PlateNumber: "FX-56671",
		SeatsCount:  10,
		BusModelId:  2,
		TypeId:      1,
		Properties:  properties,
	},
		models.Bus{
			PlateNumber: "FX-66811",
			SeatsCount:  20,
			BusModelId:  2,
			TypeId:      2,
			Properties:  properties,
		})

	DB.Create(&busses1)
	DB.Create(&busses2)
}

var locations []models.Location

func seedLocations() {
	locations = append(locations,
		models.Location{
			Name: "istanbul",
		},
		models.Location{
			Name: "ankara",
		},
		models.Location{
			Name: "izmir",
		},
		models.Location{
			Name: "denizli",
		})

	DB.Create(&locations)
}

func seedTravels() {
	var travels []models.Travel
	time1 := time.Now().UTC()
	travels = append(travels, models.Travel{
		Fee:          500,
		FromLocation: locations[0].Name,
		ToLocation:   locations[1].Name,
		Day:          time1.String(),
		Time:         time1.String(),
		Buses:        busses1,
	},
		models.Travel{
			Fee:          400,
			FromLocation: locations[2].Name,
			ToLocation:   locations[3].Name,
			Day:          time1.String(),
			Time:         time1.String(),
			Buses:        busses2,
		},
	)

	DB.Create(&travels)
}

func seedSeats() {
	var seats1 []models.Seat
	var seats2 []models.Seat
	var seatProperties []models.SeatProperty
	seatProperties = append(seatProperties,
		models.SeatProperty{
			No:     1,
			Gender: true,
		},
		models.SeatProperty{
			No:     2,
			Gender: false,
		})

	seats1 = append(seats1,
		models.Seat{
			Count:      50,
			BusId:      1,
			TravelId:   1,
			Properties: seatProperties,
		},
		models.Seat{
			Count:      20,
			BusId:      2,
			TravelId:   1,
			Properties: seatProperties,
		})

	seats2 = append(seats2,
		models.Seat{
			Count:      20,
			BusId:      4,
			TravelId:   2,
			Properties: seatProperties,
		},
		models.Seat{
			Count:      20,
			BusId:      2,
			TravelId:   2,
			Properties: seatProperties,
		})

	DB.Create(&seats1)
	DB.Create(&seats2)
}
