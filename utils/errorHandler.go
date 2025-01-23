package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomError struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewError(message string, codes ...int) *CustomError {
	code := 500
	if len(codes) > 0 {
		code = codes[0]
	}

	return &CustomError{
		Code:    code,
		Success: false,
		Message: message,
	}
}

func ErrorFormatter(err error) {
	errorMessage := ""
	for _, e := range err.(validator.ValidationErrors) {
		fmt.Printf("Field: %s, Error: %s\n", e.Field(), e.Tag())
		if e.Tag() == "required" {
			errorMessage = fmt.Sprintf("%v is Required.", e.Field())
		}

		panic(NewError(errorMessage, 403))
	}
}

func HandleError(c *gin.Context, err error) {
	c.Header("Content-Type", "application/json")
	if customErr, ok := err.(*CustomError); ok {
		c.JSON(customErr.Code, gin.H{
			"success": customErr.Success,
			"message": customErr.Message,
		})
	} else {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Internal server error",
		})
	}
}
