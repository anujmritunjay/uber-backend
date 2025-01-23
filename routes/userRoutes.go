package routes

import (
	"github.com/anujmritunjay/uber-backend/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(router *gin.Engine, db *mongo.Client) {
	userGroups := router.Group("/auth")

	{
		userGroups.POST("/register", func(c *gin.Context) {
			controllers.SignUp(c, db)
		})

		userGroups.POST("/sign-in", func(ctx *gin.Context) {
			controllers.LogIn(ctx, db)
		})
	}
}
