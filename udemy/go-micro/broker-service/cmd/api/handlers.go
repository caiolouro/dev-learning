package main

import (
	"net/http"
)

func Broker(w http.ResponseWriter, r *http.Request) {
	payload := response{
		IsError: false,
		Message: "got here!",
	}

	_ = writeJSON(w, http.StatusOK, payload)
}
