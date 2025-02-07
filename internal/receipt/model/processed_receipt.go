package model

type ProcessedReceipt struct {
	receiptId string
	receipt   *Receipt
	points    int
}

type ProcessedReceiptResponse struct {
	ID string `json:"id"`
}

type PointsTotalResponse struct {
	Points int
}

// Function to create a new PointsTotalResponse
func NewPointsTotalResponse(points int) *PointsTotalResponse {
	return &PointsTotalResponse{
		Points: points,
	}
}

// Function to create a new ProcessedReceipt
func NewProcessedReceipt(receiptId string, receipt *Receipt, points int) *ProcessedReceipt {
	return &ProcessedReceipt{
		receiptId: receiptId,
		receipt:   receipt,
		points:    points,
	}
}

// Function to create a new ProcessedReceiptResponse
func NewProcessedReceiptResponse(receiptId string) *ProcessedReceiptResponse {
	return &ProcessedReceiptResponse{
		ID: receiptId,
	}
}

func (r ProcessedReceipt) ID() string {
	return r.receiptId
}

func (r ProcessedReceipt) Receipt() *Receipt {
	return r.receipt
}

func (r ProcessedReceipt) Points() int {
	return r.points
}

func (r ProcessedReceipt) HasId() bool {
	return r.receiptId != ""
}
