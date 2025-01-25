package main

import (
	"log"
	"net/http"

	api "github.com/alexburley/askconsult-api/internal"
)


func main() {
	http.HandleFunc("/users", api.ListUsersHandler)

	port := ":8080"
	log.Println("Server starting on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}