package main

import (
	"encoding/json"
	"net/http"
)

type response struct {
	IsError bool   `json:"is_error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *App) Broker(w http.ResponseWriter, r *http.Request) {
	payload := response{
		IsError: false,
		Message: "got here!",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}
