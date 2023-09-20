package handler

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsSetup(router *gin.Engine) {
	// allow swagger UI requests
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))
}
