package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitin1chandani/Go-proj-book-mgmt/pkg/routes"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
