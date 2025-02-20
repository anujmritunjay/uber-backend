package routes

import (
	"github.com/anujmritunjay/uber-backend/controllers"
	"github.com/anujmritunjay/uber-backend/middlewares"
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

		userGroups.POST("/log-out", func(ctx *gin.Context) {
			controllers.LogOut(ctx, db)
		})

	}

	protectedRoutes := router.Group("/auth", func(ctx *gin.Context) {
		middlewares.AuthUser(ctx, db)
	})

	protectedRoutes.GET("/me", func(c *gin.Context) {
		controllers.Me(c, db)
	})
}
