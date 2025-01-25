package sql_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

var InitFile = "db/init.sql"

// Init reads SQL statements from a file and executes them
func Init(config Config) (db *sql.DB, err error) {

	// Database connection string
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName,
	)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	content, err := os.ReadFile(InitFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL file: %w", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL file: %w", err)
	}

	log.Println("SQL initialization completed successfully")

	return db, nil
}
