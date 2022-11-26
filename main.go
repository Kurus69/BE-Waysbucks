package main

import (
	"fmt"
	"net/http"
	"waysbucks/database"
	"waysbucks/pkg/connection"
	"waysbucks/routes"

	"github.com/gorilla/mux"
)

func main() {
	connection.Database()
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/waysbucks").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
