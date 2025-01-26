package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	adapters "github.com/alexburley/askconsult-api/internal/adapters/repositories"
	"github.com/alexburley/askconsult-api/internal/core/users"
	"github.com/gorilla/mux"
)

type ServerDeps struct {
	UserRepository *adapters.UserRepository
}

func NewServer(deps ServerDeps) *mux.Router {
	createUser := &CreateUserHandler{deps.UserRepository}
	listUsers := &ListUsersHandler{deps.UserRepository}

	// Set up router with dependency-injected handlers
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser.Handler).Methods("POST")
	r.HandleFunc("/users", listUsers.Handler).Methods("GET")

	return r
}

type ListUsersHandler struct {
	repository *adapters.UserRepository
}

// GetUsersHandler retrieves all users from the database
func (h *ListUsersHandler) Handler(w http.ResponseWriter, r *http.Request) {
	users, err := h.repository.List()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	// Return users in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

type CreateUserHandler struct {
	repository *adapters.UserRepository
}

// Handler handles the request to add a new user
func (h *CreateUserHandler) Handler(w http.ResponseWriter, r *http.Request) {
	var user users.User
	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userId, err := h.repository.Create(user)
	if err != nil {
		slog.Info(fmt.Sprintf("err %s", err.Error()))
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"user_id": userId,
	})
}
