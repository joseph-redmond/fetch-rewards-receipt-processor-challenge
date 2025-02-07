package validator

import (
	"math"
	"receipt-processor-challenge/internal/receipt/model"
	"regexp"
	"strconv"
	"time"
)

// IsValidReceipt Function to check if the receipt is valid according to the business rules
func IsValidReceipt(receipt model.Receipt) bool {
	validRetailer := isValidRetailer(receipt.RetailerName)
	validDate := isValidDate(receipt.PurchaseDate)
	validTime := isValidTime(receipt.PurchaseTime)
	validPrice := isValidPrice(receipt.TotalAmount)
	validItemsList := isValidItemsList(receipt.Items)
	priceMatchesItemsList := isPriceMatchingItemsListTotal(receipt.TotalAmount, receipt.Items)
	return validRetailer && validDate && validTime && validPrice && validItemsList && priceMatchesItemsList
}

// Function to check if a string is in a 24-hour time format
func isValidTime(timeStr string) bool {
	if len(timeStr) != 5 {
		return false
	}
	timeRegex := `^(2[0-3]|[01]?[0-9]):([0-5][0-9])$`
	re := regexp.MustCompile(timeRegex)
	return re.MatchString(timeStr)
}

// Function to check if a string is a valid date
func isValidDate(dateStr string) bool {
	if len(dateStr) != 10 {
		return false
	}
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}

// Function to check if the retailer name is valid
func isValidRetailer(retailerName string) bool {
	if retailerName == "" {
		return false
	}
	nameRegex := `^[\w\s\-&]+$`
	re := regexp.MustCompile(nameRegex)
	return re.MatchString(retailerName)
}

// Function to check if the price is valid
func isValidPrice(priceStr string) bool {
	if priceStr == "" {
		return false
	}
	priceRegex := `^\d+\.\d{2}$`
	re := regexp.MustCompile(priceRegex)
	return re.MatchString(priceStr)
}

// Function to check if the items list of the receipt is valid
func isValidItemsList(receiptItems []model.ReceiptItem) bool {
	if len(receiptItems) == 0 {
		return false
	}
	for _, item := range receiptItems {
		descriptionIsValid := isValidItemDescription(item.ShortDescription)
		priceIsValid := isValidPrice(item.Price)
		if !descriptionIsValid || !priceIsValid {
			return false
		}
	}
	return true
}

// Function to check if the total and items list pricing match up
func isPriceMatchingItemsListTotal(priceStr string, receiptItems []model.ReceiptItem) bool {
	if priceStr == "" {
		return false
	}
	if len(receiptItems) == 0 {
		return false
	}
	inputIsValid := isValidPrice(priceStr) && isValidItemsList(receiptItems)
	if !inputIsValid {
		return false
	}

	itemsListPriceTotal := 0.00
	reportedPrice, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return false
	}

	for _, receiptItem := range receiptItems {
		itemPriceAsFloat, err := strconv.ParseFloat(receiptItem.Price, 64)
		if err != nil {
			return false
		}
		itemsListPriceTotal += itemPriceAsFloat
	}

	truncatedItemsListPriceTotal := truncateFloat(itemsListPriceTotal, 2)
	truncatedReportedPrice := truncateFloat(reportedPrice, 2)
	return truncatedReportedPrice == truncatedItemsListPriceTotal
}

// Function to check if the item description is valid
func isValidItemDescription(itemDescription string) bool {
	if itemDescription == "" {
		return false
	}
	descriptionRegex := `^[\w\s\-]+$`
	re := regexp.MustCompile(descriptionRegex)
	return re.MatchString(itemDescription)
}

// Function that truncates a float to a given number of decimal points
func truncateFloat(f float64, decimals int) float64 {
	power := math.Pow(10, float64(decimals))
	return math.Floor(f*power) / power
}
