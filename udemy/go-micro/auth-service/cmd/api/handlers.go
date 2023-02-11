package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *App) Authenticate(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := readJSON(w, r, &payload); err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(payload.Email)
	if err != nil {
		errorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	valid, err := user.PasswordMatches(payload.Password)
	if err != nil || !valid {
		errorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	response := response{
		IsError: false,
		Message: fmt.Sprintf("Logged as user %s", user.Email),
		Data:    user,
	}

	writeJSON(w, http.StatusAccepted, response)
}
