package main

import (
	"fmt"
	"net/http"
	"waysbucks/database"
	"waysbucks/pkg/connection"
	"waysbucks/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	connection.Database()
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/waysbucks").Subrouter())

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
