package test

import (
	"context"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/dd3v/snippets.page.backend/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var database *mongo.Database

func Database(t *testing.T) *mongo.Database {
	if database != nil {
		return database
	}
	config := config.NewConfig()
	_, err := toml.DecodeFile("../../config/app.toml", config)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DatabaseURL))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = client.Connect(context.TODO())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	database := client.Database(config.TestDatabaseName)
	database.Drop(context.TODO())

	return database
}
