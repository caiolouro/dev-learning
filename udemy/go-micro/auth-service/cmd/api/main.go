package main

import (
	"auth/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const port = "80"

var connAttempts = 0

const connAttemptsDelaySeconds = 2

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting auth service...")

	conn := connectToDB()
	if conn == nil {
		log.Panic("Couldn't connec to DB")
	}

	// app := Config{
	// 	DB:     conn,
	// 	Models: data.New(conn),
	// }

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("DB not yet ready...")
			connAttempts++
		} else {
			log.Println("Connected to DB")
			return conn
		}

		if connAttempts > 10 {
			log.Println(err)
			return nil
		}

		log.Printf("Backing off for %d seconds...\n", connAttemptsDelaySeconds)
		time.Sleep(connAttemptsDelaySeconds * time.Second)
		continue
	}
}
