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

	// reading receipt and error handling
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Validating receipt fields and error handling
	if err := ValidateReceipt(receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generating a new unique id for the receipt
	id := uuid.New().String()
	receiptJSON, err := json.Marshal(receipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	err = redisClient.Set(ctx, id, receiptJSON, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store receipt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Controller to get points for a receipt
func getPointsController(c *gin.Context) {
	// Parsing id from the end point
	id := c.Param("id")

	// checking if we have the id in our In-memeory receipts map
	receiptJSON, err := redisClient.Get(ctx, id).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve receipt"})
		return
	}
	var receipt Receipt
	err = json.Unmarshal([]byte(receiptJSON), &receipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Calculating points based on the receipt
	points := calculatePoints(receipt)

	c.JSON(http.StatusOK, gin.H{"points": points})
}
