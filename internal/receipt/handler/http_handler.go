package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"receipt-processor-challenge/internal/receipt/model"
	"receipt-processor-challenge/internal/receipt/service"
	"receipt-processor-challenge/internal/receipt/validator"
)

type Handler struct {
	service *service.Service
	logger  *logrus.Logger
}

// Function for creating a new receipt handler
func NewHandler(service *service.Service, logger *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

// Function for handling the processing of receipts received through a http request
func (receiptHandler *Handler) HandleReceiptProcessing(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	ctx := request.Context()
	log := receiptHandler.logger.WithContext(ctx)

	var receipt model.Receipt
	if err := json.NewDecoder(request.Body).Decode(&receipt); err != nil {
		log.WithError(err).Error("failed to decode request body")
		http.Error(responseWriter, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	if isValidReceipt := validator.IsValidReceipt(receipt); !isValidReceipt {
		log.WithField("receipt", receipt).Error("invalid receipt")
		http.Error(responseWriter, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	savedProcessedReceipt, err := receiptHandler.service.ProcessReceipt(ctx, &receipt)
	if err != nil {
		http.Error(responseWriter, "The receipt is invalid.", http.StatusBadRequest)
		return
	}
	processedResponseId := savedProcessedReceipt.ID()
	processedReceiptResponse := model.NewProcessedReceiptResponse(processedResponseId)

	log.WithFields(logrus.Fields{"receipt_id": savedProcessedReceipt.ID()}).Info("receipt created successfully")
	responseWriter.WriteHeader(http.StatusOK)
	encodingFailed := json.NewEncoder(responseWriter).Encode(processedReceiptResponse)
	if encodingFailed != nil {
		log.WithError(err).Error("failed to encode response body")
		http.Error(responseWriter, "The receipt is invalid.", http.StatusBadRequest)
		return
	}
}

// Function for handling the processing of receipts received through a http request
func (receiptHandler *Handler) HandleReceiptFetchById(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	ctx := request.Context()
	requestId := ctx.Value("request_id").(string)
	log := receiptHandler.logger.WithContext(ctx).WithFields(logrus.Fields{"request_id": requestId})
	receiptId := mux.Vars(request)["id"]
	if receiptId == "" {
		log.Error("invalid receipt id")
		http.Error(responseWriter, "The receipt is invalid.", http.StatusBadRequest)
		return
	}
	receipt, err := receiptHandler.service.FindReceiptById(ctx, receiptId)
	if err != nil {
		log.WithError(err).Error("No receipt found for that ID:" + receiptId)
		http.Error(responseWriter, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	log.Info("receipt fetched successfully")
	responseWriter.WriteHeader(http.StatusOK)
	log.Infof("Response Body: %v", receipt)

	pointsTotalResponse := model.NewPointsTotalResponse(receipt.Points())

	err = json.NewEncoder(responseWriter).Encode(pointsTotalResponse)
	if err != nil {
		log.WithError(err).Error("failed to encode receipt")
	}
}
