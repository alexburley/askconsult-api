package main

import (
	"askconsult/routes"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/signup", routes.CreateUser)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

