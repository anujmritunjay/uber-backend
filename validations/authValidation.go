package validations

import (
	"github.com/anujmritunjay/uber-backend/models"
	"github.com/go-playground/validator/v10"
)

func SignUpValidation(user models.SignUp) {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		// fmt.Println(err)
		NewValidationError(err)

	}
}
