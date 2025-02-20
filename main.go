package main

import (
	"log"

	"github.com/anujmritunjay/uber-backend/config"
	"github.com/anujmritunjay/uber-backend/middlewares"
	"github.com/anujmritunjay/uber-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	client, ctx, cancel := config.ConnectMongoDB()
	defer cancel()
	defer client.Disconnect(ctx)

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "Uber clone backend",
		})
	})
	router.Use(middlewares.ErrorMiddleware())

	routes.UserRoutes(router, client)
	// routes.SocketRoutes(router) // WebSocket routes

	log.Println("Server is running on port 8080")

	router.Run(":8080")
}
