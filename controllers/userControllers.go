package controllers

import (
	"fmt"

	"github.com/anujmritunjay/uber-backend/config"
	"github.com/anujmritunjay/uber-backend/models"
	"github.com/anujmritunjay/uber-backend/services"
	"github.com/anujmritunjay/uber-backend/utils"
	"github.com/anujmritunjay/uber-backend/validations"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(c *gin.Context, appCtx *config.AppContext) {
	var user models.SignUp
	if err := c.BindJSON(&user); err != nil {
		panic(utils.NewError(500, "Invalid JSON"))
	}
	validations.SignUpValidation(user)

	hashPassword := services.HashPassword(user.Password)
	user.Password = hashPassword

	collection := appCtx.DB.Database("uber-backend").Collection("users")

	var isExists bson.M

	collection.FindOne(appCtx.Ctx, bson.M{"email": user.Email}).Decode(&isExists)
	fmt.Println(isExists)

	userData, err := collection.InsertOne(appCtx.Ctx, user)
	fmt.Println("Priting created user", userData.InsertedID)

	if err != nil {
		panic(utils.NewError(500, err.Error()))
	}
	if oid, ok := userData.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	} else {
		panic(utils.NewError(500, "Failed to convert InsertedID to ObjectID"))
	}

	c.JSON(200, gin.H{
		"success":  true,
		"password": user,
	})

}
