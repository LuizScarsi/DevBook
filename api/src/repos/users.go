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
func (repo Users) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}