package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/connection"
	"waysbucks/pkg/middleware"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepoTransaction(connection.DB)
	h := handlers.HandlerTransaction(transactionRepository)
	r.HandleFunc("/transaction", middleware.Auth(h.Checkout)).Methods("POST")
}
