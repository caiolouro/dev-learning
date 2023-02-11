package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type payload struct {
	Action string      `json:"action"`
	Auth   authPayload `json:"auth,omitempty"`
}

type authPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Broker(w http.ResponseWriter, r *http.Request) {
	payload := response{
		IsError: false,
		Message: "got here!",
	}

	_ = writeJSON(w, http.StatusOK, payload)
}

func HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var payload payload

	if err := readJSON(w, r, &payload); err != nil {
		errorJSON(w, err)
		return
	}

	switch payload.Action {
	case "auth":
		authenticate(w, payload.Auth)
	default:
		errorJSON(w, errors.New("unknown action"))
	}
}

func authenticate(w http.ResponseWriter, p authPayload) {
	body, _ := json.MarshalIndent(p, "", "\t")

	req, err := http.NewRequest("POST", "http://auth/authenticate", bytes.NewBuffer(body))
	if err != nil {
		errorJSON(w, err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		errorJSON(w, err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		errorJSON(w, errors.New("invalid credentials"))
		return
	} else if res.StatusCode != http.StatusAccepted {
		errorJSON(w, errors.New("error calling auth service"))
		return
	}

	var jsonRes response
	err = json.NewDecoder(res.Body).Decode(&jsonRes)
	if err != nil {
		errorJSON(w, err)
		return
	}

	if jsonRes.IsError {
		errorJSON(w, errors.New(jsonRes.Message), http.StatusUnauthorized)
		return
	}

	jsonRes.Message = "Authenticated!"
	writeJSON(w, http.StatusAccepted, jsonRes)
}
