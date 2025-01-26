package core

import "github.com/google/uuid"

// User struct to represent user data
type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
