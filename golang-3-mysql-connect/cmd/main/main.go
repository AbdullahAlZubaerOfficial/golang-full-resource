package main

import (
	"fmt" // <- import fmt
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/abdullahalzubaerofficial/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	fmt.Println("Server running on http://localhost:9010") // <- debug message

	// Serve the router
	log.Fatal(http.ListenAndServe(":9010", r))
}
