package handler

import (
	"net/http"
	"technical-test/priverion/models"
	"technical-test/priverion/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// manage the validation of auth user and books

func SignUpValidation(ctx *gin.Context, user models.User, err error) error {
	// get errors
	validationErrors := err.(validator.ValidationErrors)

	// create a list to save all errorsFields
	var errorFields []utils.FieldsMessage
	for _, fieldError := range validationErrors {
		errorFields = append(errorFields, utils.FieldsMessage{
			Field:   fieldError.Field(),
			Message: fieldError.Tag(),
		})
	}

	// create a error response
	errorMsg := "Datos de usuario no válidos"
	errorResponse := utils.ErrorResponse{
		Message: errorMsg,
		Fields:  errorFields,
	}

	ctx.IndentedJSON(http.StatusBadRequest, errorResponse)
	return err
}
