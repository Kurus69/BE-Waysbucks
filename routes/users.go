package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/connection"
	"waysbucks/pkg/middleware"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	userReposetory := repositories.RepositoryUser(connection.DB)
	h := handlers.HandlerUser(userReposetory)

	r.HandleFunc("/users", middleware.Auth(h.FindAllUser)).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}
