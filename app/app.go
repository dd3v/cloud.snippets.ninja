package app

import (
	"context"
	"net/http"

	storage "github.com/dd3v/snippets.page.backend/app/storage/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//New - init database connection, set up base configuration and return HTTP server
func New(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	storage := storage.New(db)
	server := newServer(storage)
	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(databaseURL string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
