package app

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dd3v/snippets.page.backend/internal/config"
	"github.com/dd3v/snippets.page.backend/internal/user"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func New(config *config.Config) error {

	fmt.Println("app init...")

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		fmt.Println(err)
	}
	router := mux.NewRouter().StrictSlash(true)
	v1 := router.PathPrefix("/v1").Subrouter()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.NewHTTPHandler(userService)

	v1.HandleFunc("/me", userHandler.Me).Methods("GET")

	address := fmt.Sprintf(":%v", config.BindAddr)
	httpServer := &http.Server{
		Addr:    address,
		Handler: router,
	}
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		os.Exit(-1)
	}

	return nil
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
	return client.Database("snippets"), nil
}
