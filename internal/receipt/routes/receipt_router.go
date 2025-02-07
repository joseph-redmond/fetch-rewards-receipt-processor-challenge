package routes

import (
	"github.com/gorilla/mux"
	"receipt-processor-challenge/internal/receipt/handler"
	"receipt-processor-challenge/internal/receipt/repository"
	"receipt-processor-challenge/internal/receipt/service"
	"receipt-processor-challenge/pkg/logger"
	"receipt-processor-challenge/pkg/middleware"
	"time"
)

// Function to initialize the receipt router
func InitializeReceiptRouter() *mux.Router {
	log := logger.GetLogger()
	receiptRepo := repository.NewRepository(log)
	receiptService := service.NewService(receiptRepo, log)
	receiptHandler := handler.NewHandler(receiptService, log)
	router := mux.NewRouter().PathPrefix("/receipts").Subrouter()
	router.Use(middleware.WithRequestContext)
	router.Use(middleware.WithTimeout(5 * time.Second))
	router.HandleFunc("/{id}/points", receiptHandler.HandleReceiptFetchById).Methods("GET")
	router.HandleFunc("/process", receiptHandler.HandleReceiptProcessing).Methods("POST")

	return router
}
