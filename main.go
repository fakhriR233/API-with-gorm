package main

import (
	"fmt"
	"net/http"

	"asacoco/database"
	"asacoco/pkg/mysql"
	"asacoco/routes"

	"github.com/gorilla/mux"
)

func main() {
	//DB Initialize
	mysql.DatabaseInit()

	//run Auto Migration
	database.RunMigration()

	r:= mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("Server running on port:5000")
	http.ListenAndServe("localhost:5000", r)
}