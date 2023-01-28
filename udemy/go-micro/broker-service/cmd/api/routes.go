package main

import (
	"net/http"

	ola "github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := ola.NewRouter()
}
