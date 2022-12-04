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
	r.HandleFunc("/transaction", middleware.Auth(h.AddTransaction)).Methods("POST") //user Order
	r.HandleFunc("/transactions", middleware.Auth(h.FindTrans)).Methods("GET")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
	r.HandleFunc("/history", middleware.Auth(h.HistoryTransUser)).Methods("GET")             //user profile
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.UpdateTransaction)).Methods("PATCH") //user checkout
	r.HandleFunc("/canceltrans/{id}", middleware.Auth(h.CancelTransaction)).Methods("PATCH") //admin
	r.HandleFunc("/accepttrans/{id}", middleware.Auth(h.AcceptTransaction)).Methods("PATCH") //admin
}
