package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repos"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

// CreateUser is a handler function for creating a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewUsersRepo(db)
	user.ID, err = repo.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// SearchUsers is a handler function for searching users.
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching users"))
}

// SearchUser is a handler function for searching a specific user.
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching user"))
}

// UpdateUser is a handler function for updating user information.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}

// DeleteUser is a handler function for deleting a user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}
