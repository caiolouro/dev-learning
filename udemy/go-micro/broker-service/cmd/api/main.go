package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "80"

type App struct{}

func main() {
	app := App{}

	log.Printf("Starting broker service on port %s\n", port)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
