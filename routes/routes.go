package routes

import (
	"technical-test/priverion/controllers"
	handler "technical-test/priverion/handlers"
	"technical-test/priverion/middleware"
	"technical-test/priverion/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *utils.Database) *gin.Engine {
	router := gin.Default()

	// handler of cors to the api
	handler.CorsSetup(router)
	// middleware of database to use it global for the controllers
	router.Use(middleware.DatabaseMiddleware(db))

	api := router.Group("/api")

	userRoutes := api.Group("/users")
	{
		userRoutes.POST("/signUp", controllers.SignUp)
		userRoutes.POST("/logIn", controllers.LogIn)
		userRoutes.PUT("/updateRole/:id", middleware.AuthMiddleware("ROLE_ADMIN"), controllers.UpdateRoleUser)
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

	return router
}
