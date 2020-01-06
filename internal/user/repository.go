package user

import (
	"context"

	"github.com/dd3v/snippets.page.backend/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "users"

type Repository interface {
	FindByID(id primitive.ObjectID) (entity.User, error)
	Create(user entity.User) error
	Update(user entity.User) error
	Delete(id primitive.ObjectID) error
	Count() (int64, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindByID(id primitive.ObjectID) (entity.User, error) {
	var user entity.User
	err := r.db.Collection(collection).FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	return user, err
}

func (r *repository) Create(user entity.User) error {
	if _, err := r.db.Collection(collection).InsertOne(context.TODO(), user); err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(user entity.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	if _, err := r.db.Collection(collection).UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(id primitive.ObjectID) error {
	if _, err := r.db.Collection(collection).DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return err
	}
	return nil
}

func (r *repository) Count() (int64, error) {
	return r.db.Collection(collection).CountDocuments(context.TODO(), bson.M{})
}
