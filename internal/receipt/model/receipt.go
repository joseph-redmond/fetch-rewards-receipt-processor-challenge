package model

type Receipt struct {
	RetailerName string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	TotalAmount  string        `json:"total"`
	Items        []ReceiptItem `json:"items"`
}

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
