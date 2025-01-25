package api

import (
	"encoding/json"
	"net/http"
)

// User struct to represent user data
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Handler function to list users
func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}