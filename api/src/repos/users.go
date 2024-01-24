package repos

import (
	"api/src/models"
	"database/sql"
)

// Users represents a users repository
type Users struct {
	db *sql.DB
}

// NewUsersRepo creates a users repository
func NewUsersRepo(db *sql.DB) *Users {
	return &Users{db}
}

// Create inserts an user into the database
func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}