package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/anujmritunjay/uber-backend/models"
	"github.com/anujmritunjay/uber-backend/services"
	"github.com/anujmritunjay/uber-backend/utils"
	"github.com/anujmritunjay/uber-backend/validations"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var database string = utils.Database
var userCollection = utils.UserCollection

func SignUp(c *gin.Context, db *mongo.Client) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		panic(utils.NewError("Invalid JSON"))
	}
	validations.SignUpValidation(user)

	hashPassword := services.HashPassword(user.Password)
	user.Password = hashPassword
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Database(database).Collection(userCollection)

	var isExists models.User

	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&isExists)
	if err == nil {
		panic(utils.NewError("User already exists.", 403))
	}

	userData, err := collection.InsertOne(ctx, user)

	if err != nil {
		panic(utils.NewError(err.Error()))
	}
	if oid, ok := userData.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	} else {
		panic(utils.NewError("Failed to convert InsertedID to ObjectID"))
	}

	c.JSON(200, gin.H{
		"success":  true,
		"password": user,
	})

}

func LogIn(c *gin.Context, db *mongo.Client) {
	var payload models.SignIn
	if err := c.BindJSON(&payload); err != nil {
		panic(utils.NewError("Invalid JSON."))
	}
	validations.LoginValidation(payload)

	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Database("uber-backend").Collection("users")

	err := collection.FindOne(ctx, bson.M{"email": payload.Email}).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		panic(utils.NewError("Invalid Credentials.", 404))
	}

	services.MatchPassword(payload.Password, user.Password)
	fmt.Println(user.ID)

	token := services.GenerateToken(user.ID.Hex())

	c.SetCookie("token", token, 3600*24, "/", "", true, true)

	c.JSON(200, gin.H{
		"success": true,
		"data":    user,
		"token":   token,
	})

}

func LogOut(c *gin.Context, db *mongo.Client) {
	c.SetCookie("token", "", -1, "/", "", true, true)
	c.JSON(200, gin.H{
		"success": true,
		"message": "Logged out successfully.",
	})
}

func Me(c *gin.Context, db *mongo.Client) {
	user, isExists := c.Get("user")

	if !isExists {
		panic(utils.NewError("Unauthorized.", 401))
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    user,
	})

}
