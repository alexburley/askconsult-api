package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	api "github.com/alexburley/askconsult-api/internal"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "myuser"
	dbPassword = "mypassword"
	dbName     = "askconsult"
)

func main() {
	
	// Create a database connection
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
	
	// Create users table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT NOT NULL)`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	
	createUser := &api.CreateUserHandler{DB: db}
	listUsers := &api.ListUsersHandler{DB: db}
	
	// Set up router with dependency-injected handlers
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser.Handler).Methods("POST")
	r.HandleFunc("/users", listUsers.Handler).Methods("GET")

	// Start the server
	port := ":8080"
	log.Println("Server is running on port", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Server failed:", err)
	}
}