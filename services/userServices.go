package services

import (
	"github.com/anujmritunjay/uber-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(utils.NewError(500, "Failed to hash password."))
	}
	return string(hashPass)
}
