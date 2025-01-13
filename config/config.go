package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AppContext struct {
	DB  *mongo.Client
	Ctx context.Context
}

func ConnectMongoDB() (*mongo.Client, context.Context, context.CancelFunc) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI not set in the .env file")
	}

	clientOptions := options.Client().ApplyURI(uri).
		SetMaxPoolSize(100).
		SetConnectTimeout(10 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Database connected successfully.")
	return client, ctx, cancel
}
