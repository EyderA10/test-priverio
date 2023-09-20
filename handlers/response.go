package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRegistrationError(ctx *gin.Context, err error) {
	ctx.IndentedJSON(http.StatusBadRequest, gin.H{
		"error": "User registration failed",
	})
	log.Print(fmt.Errorf("could not save a new user: %s", err.Error()))
}

func HandleJWTError(ctx *gin.Context, err error) {
	ctx.IndentedJSON(http.StatusBadRequest, gin.H{
		"error": "token generation failed, send email or password correct",
	})
	log.Print(fmt.Errorf("could not generate JWT token: %s", err.Error()))
}
