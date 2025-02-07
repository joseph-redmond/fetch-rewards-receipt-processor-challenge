package routes

import (
	"github.com/gorilla/mux"
	"receipt-processor-challenge/internal/receipt/routes"
)

// Router Function that initializes the Main Router that merges all subrouters
func InitializeRouter() *mux.Router {
	mainRouter := mux.NewRouter()
	receiptRouter := routes.InitializeReceiptRouter()
	mainRouter.PathPrefix("/receipts").Handler(receiptRouter).Methods("POST", "GET")
	return mainRouter
}
