package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

// Function to calculate points for a receipt
func calculatePoints(receipt Receipt) int {
	points := 0

	// Calculating points for retailer using alphanumeric characters constraint
	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points++
		}
	}

	// Calculating points for total constraints
	totalValue, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return points
	}

	if math.Mod(totalValue, 1) == 0 {
		points += 50
	}
	if math.Mod(totalValue, 0.25) == 0 {
		points += 25
	}

	// Calculating points for items constraints
	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			priceValue, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(priceValue * 0.2))
			}
		}
	}

	// Calculating points for date constraint
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Calculating points for time constraint
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() == 14 && purchaseTime.Minute() == 0 {
		// Skip adding points for exactly 2:00pm
	} else if err == nil && purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
