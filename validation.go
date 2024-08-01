package main

import (
	"fmt"
	"strconv"
	"time"
)

func ValidateReceipt(receipt Receipt) error {
	if receipt.Retailer == "" {
		return fmt.Errorf("retailer is required")
	}
	if receipt.PurchaseDate == "" {
		return fmt.Errorf("purchaseDate is required")
	}
	if _, err := time.Parse("2006-01-02", receipt.PurchaseDate); err != nil {
		return fmt.Errorf("invalid purchaseDate format, expected YYYY-MM-DD")
	}
	if receipt.PurchaseTime == "" {
		return fmt.Errorf("purchaseTime is required")
	}
	if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		return fmt.Errorf("invalid purchaseTime format, expected HH:MM")
	}
	if receipt.Total == "" {
		return fmt.Errorf("total is required")
	}
	if _, err := strconv.ParseFloat(receipt.Total, 64); err != nil {
		return fmt.Errorf("invalid total format, expected a numeric value") // Ensure total amount is numeric
	}
	if len(receipt.Items) == 0 {
		return fmt.Errorf("at least one item is required")
	}
	for _, item := range receipt.Items {
		if item.ShortDescription == "" {
			return fmt.Errorf("item shortDescription is required")
		}
		if _, err := strconv.ParseFloat(item.Price, 64); err != nil {
			return fmt.Errorf("invalid total format, expected a numeric value") // Ensure total amount is numeric
		}
	}
	return nil
}
