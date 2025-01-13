package routes

import (
	"github.com/anujmritunjay/uber-backend/sockets"
	"github.com/gin-gonic/gin"
)

// SocketRoutes defines the WebSocket routes
func SocketRoutes(router *gin.Engine) {
	router.GET("/ws", sockets.HandleWebSocket)
}
