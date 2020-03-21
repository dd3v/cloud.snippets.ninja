package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/dd3v/snippets.page.backend/internal/auth"
	"github.com/dd3v/snippets.page.backend/internal/config"
	"github.com/dd3v/snippets.page.backend/internal/user"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../../config/app.toml", "path to config file")
}

func main() {
	config := config.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(db)

	router := routing.New()
	router.Use(
		content.TypeNegotiator(content.JSON),
	)
	apiGroup := router.Group("/api")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	user.NewHTTPHandler(apiGroup.Group("/v1"), userService)

	auth.NewHTTPHandler(apiGroup.Group("/v1"), auth.NewService(config.JWTSigningKey, auth.NewRepository(db)))

	address := fmt.Sprintf(":%v", config.BindAddr)
	httpServer := &http.Server{
		Addr:    address,
		Handler: router,
	}
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		os.Exit(-1)
	}

}

func newDB(databaseURL string) (*mongo.Database, error) {
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
	_, err = client.Database("snippets").Collection("users").Indexes().CreateOne(
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
	_, err = client.Database("snippets").Collection("users").Indexes().CreateOne(
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

	return client.Database("snippets"), nil
}
