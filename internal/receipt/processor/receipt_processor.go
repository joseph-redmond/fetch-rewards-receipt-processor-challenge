package processor

import (
	"github.com/google/uuid"
	"math"
	"receipt-processor-challenge/internal/receipt/model"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Function to process a new receipt
func ProcessReceipt(receipt *model.Receipt) *model.ProcessedReceipt {
	points := getPoints(receipt)
	processedReceipt := model.NewProcessedReceipt(uuid.New().String(), receipt, points)
	return processedReceipt
}

// Function to calculate the points from a receipt
func getPoints(receipt *model.Receipt) int {
	points := 0
	points += calculatePointFromRetailerName(receipt.RetailerName)
	points += calculatePointsFromRoundTotalAmount(receipt.TotalAmount)
	points += calculatePointsFromMultipleOfQuarterAmount(receipt.TotalAmount)
	points += calculatePointsFromNumberItemsOnReceipt(receipt.Items)
	points += calculatePointsForItemDescriptionLengthIsMultipleOfThree(receipt.Items)
	points += calculatePointsFromPurchaseDayBeingOdd(receipt.PurchaseDate)
	points += calculatePointsFromPurchaseTimeBeingBetweenTwoAndFourPM(receipt.PurchaseTime)
	return points
}

// Function to calculate the points from the retailer name according to the business rules
func calculatePointFromRetailerName(retailerName string) int {
	alphanumericRegex := regexp.MustCompile(`[a-zA-Z0-9]`)
	return len(alphanumericRegex.FindAllString(retailerName, -1))
}

// Function to calculate the points from the total regarding round numbers according to the business rules
func calculatePointsFromRoundTotalAmount(totalAmount string) int {
	if isRoundDollar(totalAmount) {
		return 50
	}
	return 0
}

// Function to calculate the points from the total regarding multiple of 0.25 according to the business rules
func calculatePointsFromMultipleOfQuarterAmount(totalAmount string) int {
	if isMultipleOfQuarter(totalAmount) {
		return 25
	}
	return 0
}

// Function to calculate the points from the number of items on the receipt according to the business rules
func calculatePointsFromNumberItemsOnReceipt(items []model.ReceiptItem) int {
	pairs := len(items) / 2
	return pairs * 5
}

// Function to calculate the points from the item description length being multiple of three according to the business rules
func calculatePointsForItemDescriptionLengthIsMultipleOfThree(items []model.ReceiptItem) int {
	pointTotal := 0
	for _, item := range items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				continue
			}
			points := math.Ceil(price * 0.2)
			pointTotal += int(points)
		}
	}
	return pointTotal
}

// Function to calculate the points from the purchase day being odd according to the business rules
func calculatePointsFromPurchaseDayBeingOdd(purchaseDate string) int {
	date, err := time.Parse("2006-01-02", purchaseDate)
	if err == nil && date.Day()%2 != 0 {
		return 6
	}
	return 0
}

// Function to calculate the points from the purchase time being between 2 PM and 4 PM according to the business rules
func calculatePointsFromPurchaseTimeBeingBetweenTwoAndFourPM(purchaseTime string) int {
	dateTime := "2025-01-01 " + purchaseTime
	parsedTime, err := time.Parse("2006-01-02 15:04", dateTime)

	if err == nil && parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		return 10
	}
	return 0
}

// Function to check if a value is a multiple of 0.25
func isMultipleOfQuarter(value string) bool {
	convertedFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	return int(convertedFloat*100)%25 == 0
}

// Function to check if a value is a round dollar value ending in .00
func isRoundDollar(value string) bool {
	convertedFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	return int(convertedFloat*100)%100 == 0
}
