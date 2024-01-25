package models

import (
	"errors"
	"strings"
	"time"
)

// User represents a social media user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare calls methods to validate and format received user
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	
	user.format()
	return nil
}
func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("name field is required and cannot be empty")
	}

	if user.Nick== "" {
		return errors.New("nick field is required and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email field is required and cannot be empty")
	}

	if user.Password == "" {
		return errors.New("password field is required and cannot be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}