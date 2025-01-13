package routes

import (
	"github.com/anujmritunjay/uber-backend/config"
	"github.com/anujmritunjay/uber-backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, appCtx *config.AppContext) {
	userGroups := router.Group("/auth")

	{
		userGroups.POST("/register", func(c *gin.Context) {
			controllers.SignUp(c, appCtx)
		})
	}
}
