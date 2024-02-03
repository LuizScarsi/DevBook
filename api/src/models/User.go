package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}
func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name field is required and cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("nick field is required and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email field is required and cannot be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid e-mail")
	}

	if step == "register" && user.Password == "" {
		return errors.New("password field is required and cannot be empty")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashedPasswd, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPasswd)
	}
	return nil
}
