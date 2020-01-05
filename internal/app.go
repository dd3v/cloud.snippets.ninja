package app

import (
	"fmt"

	"github.com/dd3v/snippets.page.backend/internal/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	logger *logrus.Logger
	router *mux.Router
}

func New(config *config.Config) error {

	fmt.Println("app init...")

	return nil

}

// func newDB(databaseURL string) (*mongo.Client, error) {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = client.Connect(context.TODO())
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = client.Ping(context.TODO(), readpref.Primary())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return client, nil
// }
