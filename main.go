package main

import (
	"log"
	"net/http"

	sql_db "github.com/alexburley/askconsult-api/db"
	api "github.com/alexburley/askconsult-api/internal"
	adapters "github.com/alexburley/askconsult-api/internal/adapters/repositories"
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
	defer db.Close()
	if err != nil {
		log.Fatal("db init failed:", err.Error())
	}

	server := api.NewServer(api.ServerDeps{UserRepository: adapters.NewUserRepository(db)})

	port := ":8080"
	log.Println("server is running on port", port)
	if err := http.ListenAndServe(port, server); err != nil {
		log.Fatal("server failed:", err)
	}
}
