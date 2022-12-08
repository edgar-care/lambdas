package database

import (
	"context"
	"log"
	"time"

	"github.com/edgar-care/graphql/lib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

type ErrorBody struct {
	Message string `json:"message"`
}

func Connect(dbUrl string) *DB {
	log.Println("Connecting to " + dbUrl)
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	lib.CheckError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	lib.CheckError(err)

	return &DB{client: client}
}
