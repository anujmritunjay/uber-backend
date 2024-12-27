package controllers

import (
	"github.com/anujmritunjay/uber-backend/models"
	"github.com/anujmritunjay/uber-backend/utils"
	"github.com/anujmritunjay/uber-backend/validations"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.SignUp
	if err := c.BindJSON(&user); err != nil {
		panic(utils.NewError(500, "Invalid JSON"))
	}

	validations.SignUpValidation(user)

}
