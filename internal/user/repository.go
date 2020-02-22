package user

import (
	"context"

	"github.com/dd3v/snippets.page.backend/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "users"

//Repository - ...
type Repository interface {
	FindByID(context context.Context, id string) (entity.User, error)
	Find(context context.Context, filter map[string]interface{}) ([]entity.User, error)
	Create(context context.Context, user entity.User) error
	Update(context context.Context, user entity.User) error
	Delete(context context.Context, id string) error
	Count(context context.Context) (int, error)
}

type repository struct {
	db *mongo.Database
}

//NewRepository - ...
func NewRepository(db *mongo.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Find(context context.Context, filter map[string]interface{}) ([]entity.User, error) {
	var users []entity.User
	filterCursor, err := r.db.Collection(collection).Find(context, filter)
	err = filterCursor.All(context, &users)

	return users, err
}

func (r *repository) FindByID(context context.Context, id string) (entity.User, error) {
	var user entity.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	err = r.db.Collection(collection).FindOne(context, bson.M{"_id": objectID}).Decode(&user)

	return user, err
}

func (r *repository) Create(context context.Context, user entity.User) error {
	if _, err := r.db.Collection(collection).InsertOne(context, user); err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(context context.Context, user entity.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	if _, err := r.db.Collection(collection).UpdateOne(context, filter, update); err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(context context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if _, err := r.db.Collection(collection).DeleteOne(context, bson.M{"_id": objectID}); err != nil {
		return err
	}
	return nil
}

func (r *repository) Count(context context.Context) (int, error) {
	count, err := r.db.Collection(collection).CountDocuments(context, bson.M{})

	return int(count), err
}
