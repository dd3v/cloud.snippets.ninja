package mongo

import "go.mongodb.org/mongo-driver/mongo"

type Storage struct {
	db *mongo.Client
}

func New(db *mongo.Client) *Storage {
	return &Storage{
		db: db,
	}
}
