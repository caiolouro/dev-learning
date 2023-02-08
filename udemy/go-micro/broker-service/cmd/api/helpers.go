package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type response struct {
	IsError bool   `json:"is_error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	d := json.NewDecoder(r.Body)
	if err := d.Decode(data); err != nil {
		return err
	}

	if err := d.Decode(&struct{}{}); err != nil {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 { // using variadics to emulate this last parameter optional
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write(out); err != nil {
		return err
	}

	return nil
}

func errorJSON(w http.ResponseWriter, err error, status ...int) error {
	code := http.StatusBadRequest
	if len(status) > 0 {
		code = status[0]
	}

	var payload response
	payload.IsError = true
	payload.Message = err.Error()

	return writeJSON(w, code, payload)
}
