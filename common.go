package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// calculatePoints calculates the points based on the receipt details.
func calculatePoints(receipt Receipt) int {
	points := 0

	// 1. One point for every alphanumeric character in the retailer name
	alphaNumericCount := len(regexp.MustCompile(`[\\w]`).FindAllString(receipt.Retailer, -1))
	points += alphaNumericCount

	// 2. 50 points if the total is a round dollar amount with no cents
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}

	// 3. 25 points if the total is a multiple of 0.25
	totalFloat, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(totalFloat*100, 25) == 0 {
		points += 25
	}

	// 4. 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// 5. If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			priceFloat, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(priceFloat * 0.2))
		}
	}

	// 6. 6 points if the day in the purchase date is odd
	parsedDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if parsedDate.Day()%2 != 0 {
		points += 6
	}

	// 7. 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
