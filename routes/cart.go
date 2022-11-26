package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/connection"
	"waysbucks/pkg/middleware"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(connection.DB)
	h := handlers.HandlerCart(cartRepository)
	r.HandleFunc("/cart/{id}", middleware.Auth(h.AddCart)).Methods("POST")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
}
