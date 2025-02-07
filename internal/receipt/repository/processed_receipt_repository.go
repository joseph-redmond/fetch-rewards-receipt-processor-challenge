package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"receipt-processor-challenge/internal/receipt/model"
	"receipt-processor-challenge/pkg/db"
)

type Repository struct {
	Store  *db.Store[model.ProcessedReceipt]
	Logger *logrus.Logger
}

// Function to create a new Processed Receipt Repository
func NewRepository(logger *logrus.Logger) *Repository {
	return &Repository{
		Store:  db.NewStore[model.ProcessedReceipt](),
		Logger: logger,
	}
}

// Function to save a new processed receipt to the dataset for persistence
func (receiptRepository *Repository) Save(ctx context.Context, receipt *model.ProcessedReceipt) (model.ProcessedReceipt, error) {
	logger := receiptRepository.Logger
	logger.Infof("saving processed receipt with id %v to the database", receipt.ID())
	savedEntity, err := receiptRepository.Store.Save(*receipt)
	if err != nil {
		logger.Errorf("failed to save receipt with id %v to the database: %v", receipt.ID, err)
	}
	return savedEntity, err
}

// Function to find a processed receipt in the dataset by it's id
func (receiptRepository *Repository) FindById(ctx context.Context, id uuid.UUID) (model.ProcessedReceipt, error) {
	log := receiptRepository.Logger

	log.Infof("fetching receipt with id %v from the database", id)

	receipt, err := receiptRepository.Store.FindById(id.String())
	if err != nil {
		log.Errorf("failed to fetch receipt with id %v from the database: %v", id, err)
	}
	return receipt, err
}

// Function to delete a processed receipt from the dataset by it's id
func (receiptRepository *Repository) DeleteById(ctx context.Context, id uuid.UUID) error {
	log := receiptRepository.Logger
	log.Infof("deleting receipt with id %v from the database", id)

	err := receiptRepository.Store.DeleteById(id.String())
	if err != nil {
		log.Errorf("failed to delete receipt with id %v from the database: %v", id, err)
	}
	return err
}
