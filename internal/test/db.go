package test

import (
	"context"
	"log"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/dd3v/snippets.page.backend/internal/config"
	"go.mongodb.org/mongo-driver/bson"
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

	_, err = client.Database("snippets_test").Collection("users").Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Database("snippets_test").Collection("users").Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"login": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return database
}
