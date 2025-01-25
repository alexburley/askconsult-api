package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

// User struct to represent user data
type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ListUsersHandler struct {
	DB *sql.DB
}

// GetUsersHandler retrieves all users from the database
func (h *ListUsersHandler) Handler(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			slog.Info(fmt.Sprintf("err %s", err.Error()))
			http.Error(w, "Error scanning user data", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Return users in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

type CreateUserHandler struct {
	DB *sql.DB
}

// Handler handles the request to add a new user
func (h *CreateUserHandler) Handler(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert user into the database
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	var userID uuid.UUID
	err := h.DB.QueryRow(query, user.Name).Scan(&userID)
	if err != nil {
		slog.Info(fmt.Sprintf("err %s", err.Error()))
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"user_id": userID,
	})
}
