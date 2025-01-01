package validations

import (
	"fmt"

	"github.com/anujmritunjay/uber-backend/utils"
	"github.com/go-playground/validator/v10"
)

func NewValidationError(err error) {
	errorMessage := ""
	for _, e := range err.(validator.ValidationErrors) {
		if e.Tag() == "required" {
			errorMessage = fmt.Sprintf("%v is Required.", e.Field())
		}

		if e.Tag() == "email" {
			errorMessage = "Please provide a valid email."
		}

		if e.Tag() == "min" {
			errorMessage = fmt.Sprintf("%v is should be min %v character.", e.Field(), e.Param())
		}
		if e.Tag() == "max" {
			errorMessage = fmt.Sprintf("%v is should be max %v character.", e.Field(), e.Param())
		}

		panic(utils.NewError(403, errorMessage))
	}
}
