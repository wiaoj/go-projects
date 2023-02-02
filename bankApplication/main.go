package main

import (
	"bank-application/controllers"
	"bank-application/initializers"
	"bank-application/middlewares"
	"bank-application/migration"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migration.Migration()
}

func main() {
	router := initRouter()
	router.Run()
}

func initRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", middlewares.AdminAuthorization(), controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		bank := api.Group("/banks")
		{
			bank.POST("/", middlewares.AdminAuthorization(), controllers.CreateBank)
			bank.DELETE("/:id", middlewares.AdminAuthorization(), controllers.DeleteBank)

			bank.GET("/", controllers.GetAllBanks)
			bank.GET("/:id", controllers.GetByIdBank)
		}

		interest := api.Group("/interests")
		{
			interest.POST("/", middlewares.AdminAuthorization(), controllers.CreateInterest)
			interest.DELETE("/:id", middlewares.AdminAuthorization(), controllers.DeleteInterest)

			interest.GET("/", controllers.GetAllInterest)
			interest.GET("/q", controllers.GetInterestsQuery)
		}

		admin := api.Group("/admin")
		{
			admin.POST("/migrate", middlewares.AdminAuthorization(), controllers.CreateMigration)
			admin.POST("/deletedatabase", middlewares.AdminAuthorization(), controllers.DeleteDatabase)
		}
	}
	return router
}
