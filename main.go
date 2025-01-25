package main

import (
	"fmt"
	"log"
	"net/http"

	sql_db "github.com/alexburley/askconsult-api/db"
	api "github.com/alexburley/askconsult-api/internal"
	"github.com/gorilla/mux"
)

// DBConfig creates and returns a Config instance with default values
func DBConfig() *sql_db.Config {
	return &sql_db.Config{
		DBHost:     "localhost",
		DBPort:     5432,
		DBUser:     "myuser",
		DBPassword: "mypassword",
		DBName:     "askconsult",
	}
}

func main() {

	db, err := sql_db.Init(*DBConfig())
	if err != nil {
		panic(fmt.Sprintf("db init failed: %s", err.Error()))
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
