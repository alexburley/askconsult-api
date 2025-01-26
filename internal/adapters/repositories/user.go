package adapters

import (
	"database/sql"
	"fmt"

	"github.com/alexburley/askconsult-api/internal/core"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) List() ([]core.User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, fmt.Errorf("reading users: %v", err)
	}
	defer rows.Close()

	var users []core.User
	for rows.Next() {
		var user core.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Create(user *core.User) (uuid.UUID, error) {
	// Insert user into the database
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	var userId uuid.UUID
	err := r.db.QueryRow(query, user.Name).Scan(&userId)
	if err != nil {
		return uuid.Nil, fmt.Errorf("inserting user: %v", err)
	}

	return userId, nil

}
