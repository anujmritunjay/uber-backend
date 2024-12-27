package routes

import (
	"github.com/anujmritunjay/uber-backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroups := router.Group("/auth")

	{
		userGroups.POST("/sign-up", controllers.SignUp)
	}
}
