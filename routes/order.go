package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/connection"
	"waysbucks/pkg/middleware"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryOrder(connection.DB)
	h := handlers.HandlerOrder(orderRepository)

	r.HandleFunc("/detailProduct/{id}", middleware.Auth(h.AddOrder)).Methods("POST")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.DeleteOrder)).Methods("DELETE")
	r.HandleFunc("/myOrder", middleware.Auth(h.GetOrderUser)).Methods("GET")
}
