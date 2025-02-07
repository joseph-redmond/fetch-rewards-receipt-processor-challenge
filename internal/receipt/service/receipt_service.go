package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"receipt-processor-challenge/internal/receipt/model"
	"receipt-processor-challenge/internal/receipt/processor"
	"receipt-processor-challenge/internal/receipt/repository"
)

type Service struct {
	repo   *repository.Repository
	logger *logrus.Logger
}

// Function to create a new Receipt Service
func NewService(repo *repository.Repository, logger *logrus.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

// Function to process a receipt
func (receiptService *Service) ProcessReceipt(ctx context.Context, receipt *model.Receipt) (*model.ProcessedReceipt, error) {
	logger := receiptService.logger

	logger.Infoln("Processing receipt")
	processedReceipt := processor.ProcessReceipt(receipt)
	savedProcessedReceipt, err := receiptService.saveProcessedReceipt(ctx, processedReceipt)
	if err != nil {
		logger.Errorf("Error saving processed receipt: %v", err)
		return &model.ProcessedReceipt{}, err
	}
	return &savedProcessedReceipt, nil
}

// Function to find a receipt by it's id
func (receiptService *Service) FindReceiptById(ctx context.Context, receiptId string) (model.ProcessedReceipt, error) {
	logger := receiptService.logger
	logger.Infof("Calling service to find receipt")
	parsedReceiptId, parseError := uuid.Parse(receiptId)
	if parseError != nil {
		logger.Errorf("Error parsing receipt id: %v", parseError)
		return model.ProcessedReceipt{}, parseError
	}

	receipt, err := receiptService.repo.FindById(ctx, parsedReceiptId)
	if err != nil {
		logger.Errorf("Error finding receipt: %v", err)
		return model.ProcessedReceipt{}, err
	}
	return receipt, err
}

// Function to delete a receipt by it's id
func (receiptService *Service) DeleteReceiptById(ctx context.Context, receiptId string) error {
	logger := receiptService.logger
	logger.Infof("Calling service to delete receipt")

	parsedReceiptId, parseError := uuid.Parse(receiptId)
	if parseError != nil {
		logger.Errorf("Error parsing receipt id: %v", parseError)
		return parseError
	}

	err := receiptService.repo.DeleteById(ctx, parsedReceiptId)
	if err != nil {
		logger.Errorf("Error deleting receipt: %v", err)
	}
	return err
}

// Function to save a processed receipt to the dataset for persistence
func (receiptService *Service) saveProcessedReceipt(ctx context.Context, receipt *model.ProcessedReceipt) (model.ProcessedReceipt, error) {
	logger := receiptService.logger

	logger.Infof("Calling service to create receipt")
	savedReceipt, err := receiptService.repo.Save(ctx, receipt)
	if err != nil {
		logger.Errorf("Error creating receipt: %v", err)
	}
	return savedReceipt, err
}
