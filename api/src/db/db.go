package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Connect opens the DB connection and returns it
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DBConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}