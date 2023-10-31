package main

import (
	"golang-scyllacloud/db"
	api "golang-scyllacloud/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Scyllaconnect()
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", api.HelloHandler)

	// Start the server
	http.ListenAndServe(":8080", r)
}
