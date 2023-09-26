package handler

import (
	"fmt"
	"net/http"
	"technical-test/priverion/models"
	"technical-test/priverion/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// manage the validation of auth user and books
func SignUpValidation(ctx *gin.Context, user *models.User) error {
	// get errors
	fmt.Println(*user)
	// validation of fields
	validate := utils.ValidatorFields()
	err := validate.Struct(*user)
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
	errorMsg := "User data invalid"
	errorResponse := utils.ErrorResponse{
		Message: errorMsg,
		Fields:  errorFields,
	}

	ctx.IndentedJSON(http.StatusBadRequest, errorResponse)
	return err
}
