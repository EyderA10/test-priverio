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

	bookRoutes := api.Group("/books")
	{
		bookRoutes.Use(middleware.MiddlewareJWT())

		bookRoutes.GET("/", controllers.GetBooks)
		bookRoutes.GET("/:id", controllers.GetBookByID)
		bookRoutes.POST("/createBook", controllers.CreateBook)
		bookRoutes.PUT("/updateBook/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/deleteBook/:id", controllers.DeleteBook)
	}

	// TODO: implement cors
	return router
}
