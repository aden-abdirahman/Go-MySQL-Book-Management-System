package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aden-abdirahman/Go-MySQL-Book-Management-System/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// creating new instance of router, registering routes, and starting server
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is started on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
