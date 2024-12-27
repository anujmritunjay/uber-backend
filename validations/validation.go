package validations

import (
	"fmt"

	"github.com/anujmritunjay/uber-backend/utils"
	"github.com/go-playground/validator/v10"
)

func NewValidationError(err error) {
	errorMessage := ""
	for _, e := range err.(validator.ValidationErrors) {
		fmt.Printf("Field: %s, Error: %s\n", e.Field(), e.Tag())
		if e.Tag() == "required" {
			errorMessage = fmt.Sprintf("%v is Required.", e.Field())
		}

		if e.Tag() == "email" {
			errorMessage = "Please provide a valid email."
		}

		panic(utils.NewError(403, errorMessage))
	}
}
