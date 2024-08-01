package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var ctx = context.Background()

// Entry point of the application
func main() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})

	// Ping Redis to check if the connection is successful
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	// Create a new Gin router
	router := gin.Default()

	// Middleware for logging and recovery
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Setting up routes
	router.POST("/receipts/process", processReceiptController)
	router.GET("/receipts/:id/points", getPointsController)

	log.Println("Starting server on :8080")
	router.Run(":8080")
}

