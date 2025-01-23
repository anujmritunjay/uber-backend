package services

import (
	"github.com/anujmritunjay/uber-backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(utils.NewError("Failed to hash password."))

	}
	return string(hashPass)
}

func MatchPassword(userPassword, databasePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(userPassword))
	if err != nil {
		panic(utils.NewError("Invalid Credentials.", 401))
	}
	return true
}

func GenerateToken(id string) string {
	claims := jwt.MapClaims{
		"_id": id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(utils.JWT_SECRET))
	if err != nil {
		panic(utils.NewError(err.Error()))
	}
	return tokenString
}
