package middlewares

import (
	"context"
	"strings"
	"time"

	"github.com/anujmritunjay/uber-backend/models"
	"github.com/anujmritunjay/uber-backend/services"
	"github.com/anujmritunjay/uber-backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthUser(c *gin.Context, db *mongo.Client) {
	token := c.GetHeader("Authorization")
	if token == "" {
		panic(utils.NewError("Unauthorized.", 401))
	}

	token = strings.Split(token, " ")[1]
	claim := services.DecodeJWT(token)

	id, ok := claim["_id"].(string)
	if !ok {
		panic(utils.NewError("Unauthorized.", 401))
	}

	if id == "" {
		panic(utils.NewError("Unauthorized.", 401))
	}

	objectID, _ := primitive.ObjectIDFromHex(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Database(utils.Database).Collection(utils.UserCollection)
	var user models.User
	err := collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		panic(utils.NewError(err.Error()))
	}
	c.Set("user", user)
	c.Next()

}
