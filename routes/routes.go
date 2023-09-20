package routes

import (
	"technical-test/priverion/controllers"
	"technical-test/priverion/middleware"
	"technical-test/priverion/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *utils.Database) *gin.Engine {
	router := gin.Default()

	// middleware of database to use it global for the controllers
	router.Use(middleware.DatabaseMiddleware(db))

	api := router.Group("/api")

	userRoutes := api.Group("/users")
	{
		userRoutes.POST("/signUp", controllers.SignUp)
		userRoutes.POST("/logIn", controllers.LogIn)
	}

	// TODO: implement cors
	return router
}
