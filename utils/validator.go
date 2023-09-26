package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func ValidatorFields() *validator.Validate {
	validate = validator.New(validator.WithRequiredStructEnabled())
	return validate
}
