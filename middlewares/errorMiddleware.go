package middlewares

import (
	"github.com/anujmritunjay/uber-backend/utils"
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Content-Type", "application/json")
		defer func() {
			if r := recover(); r != nil {
				err := r.(error)
				utils.HandleError(ctx, err)
			}
		}()
		ctx.Next()
	}
}
