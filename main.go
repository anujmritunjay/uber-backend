package main

import (
	"log"

	"github.com/anujmritunjay/uber-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	client, ctx, cancel := config.ConnectMongoDB()
	defer cancel()
	defer client.Disconnect(ctx)

	router := gin.Default()

	log.Println("Server is running on port 8080")

	router.Run(":8080")
}
