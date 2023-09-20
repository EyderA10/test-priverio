package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	handler "technical-test/priverion/handlers"
	"technical-test/priverion/models"
	"technical-test/priverion/services"
	"technical-test/priverion/utils"

	"github.com/gin-gonic/gin"
)

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
		"auth":    token,
	})
}

func LogIn(ctx *gin.Context) {
	// get the database
	db := ctx.MustGet("db").(*utils.Database)
	// import UserService to instance and then use the LogIn method
	userService := services.NewUserService(db, db.GetName(), "users")
	token, errJWT := userService.LogIn(ctx)
	if errJWT != nil {
		handler.HandleJWTError(ctx, errJWT)
		return
	}
	// User registration successful
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"auth": token,
	})
}

func UpdateRoleUser(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*utils.Database)
	userService := services.NewUserService(db, os.Getenv("DATABASE_NAME"), "users")
	var updatedUser models.User

	// bind json to the updatedBook
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		log.Print(fmt.Errorf("could not bind JSON: %w", err))
		return
	}
	modifiedCount, err := userService.UpdateRoleUser(id, updatedUser)
	if err != nil {
		log.Printf("Error updating user role: %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to update user role"})
		return
	}

	if modifiedCount == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found or not updated"})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Role Updated Succesfully!",
	})
}
