package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	RPCPort  = "5001"
	mongoURL = "mongodb://mongo:2701"
	gRPCPort = "50001"
)

var client *mongo.Client

type App struct {
}

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connectToMongo() (*mongo.Client, error) {
	opts := options.Client().ApplyURI(mongoURL)
	opts.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	c, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println("Error connecting to MongoDB:", err)
		return nil, err
	}

	return c, nil
}
