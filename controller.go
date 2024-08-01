package main

import (
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// Controller to process a new receipt and store it
func processReceiptController(c *gin.Context) {
	var receipt Receipt

	// read receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Validate receipt fields
	if err := ValidateReceipt(receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new unique id for the receipt
	id := uuid.New().String()

    // Marshal the receipt into JSON
	receiptJSON, err := json.Marshal(receipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Store the receipt JSON in Redis with the generated UUID as the key
    err = redisClient.Set(ctx, id, receiptJSON, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store receipt"})
		return
	}

	// Return the generated UUID in the response
    c.JSON(http.StatusOK, gin.H{"id": id})
}

// Controller to get points for a receipt
func getPointsController(c *gin.Context) {
	// Get the receipt ID from the URL parameter
	id := c.Param("id")

	// Retrieve the receipt JSON from Redis using the ID
	receiptJSON, err := redisClient.Get(ctx, id).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve receipt"})
		return
	}

	var receipt Receipt

    // Unmarshal the JSON back into the receipt variable
	err = json.Unmarshal([]byte(receiptJSON), &receipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Calculate points based on the receipt
	points := calculatePoints(receipt)

    // Return the points as the response
	c.JSON(http.StatusOK, gin.H{"points": points})
}

