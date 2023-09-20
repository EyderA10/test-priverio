package controllers

import (
	"net/http"
	handler "technical-test/priverion/handlers"
	"technical-test/priverion/services"
	"technical-test/priverion/utils"

	"github.com/gin-gonic/gin"
)

// TODO: al enviarlo no valida correctamente lo que llega por requestBody
func SignUp(ctx *gin.Context) {
	// get the database
	db := ctx.MustGet("db").(*utils.Database)
	// import UserService to instance and then use the SignUp method
	userService := services.NewUserService(db, db.GetName(), "users")
	user, err := userService.SignUp(ctx)
	// Check for errors
	if err != nil {
		// Handle the error and send an appropriate response
		handler.HandleRegistrationError(ctx, err)
		return
	}

	// Generate a JWT token for the registered user
	token, errJWT := userService.GenerateJWT(user)
	// Check for JWT token generation errors
	if errJWT != nil {
		handler.HandleJWTError(ctx, errJWT)
		return
	}
	// User registration successful
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"data":    user,
		"token":   token,
	})
}

func LogIn(ctx *gin.Context) {
	// implement Login from UserService
	// userService := services.NewUserService()
}
