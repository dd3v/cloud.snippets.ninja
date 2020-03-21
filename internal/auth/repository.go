package auth

import (
	"context"
	"fmt"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = "users"

//Repository - ...
type Repository interface {
	FindUser(context context.Context, login string) (entity.User, error)
	SaveRefreshToken(context context.Context, userID string, refreshTokens entity.RefreshTokens) error
	ReplaceRefreshToken(context context.Context, refreshToken string) error
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

func (r repository) ReplaceRefreshToken(context context.Context, refreshToken string) error {

	// z := entity.RefreshTokens{
	// 	ID:        primitive.NewObjectID(),
	// 	Token:     "fuck",
	// 	Exp:       time.Now().Add(time.Hour * 100),
	// 	CreatedAt: time.Now(),
	// }

	var a entity.RefreshTokens

	err := r.db.Collection(collection).FindOne(context,
		bson.M{"refreshTokens": bson.M{"$elemMatch": bson.M{"token": refreshToken}}},
	).Decode(&a)

	if err != nil {
		fmt.Println("part of arrat error")
		fmt.Println(err)
	}

	fmt.Println(r)

	singleResult := r.db.Collection(collection).FindOneAndUpdate(
		context,
		bson.M{
			"refreshTokens.token": refreshToken,
		},
		bson.M{
			"$pull": bson.M{"refreshTokens.token": refreshToken},
		},
	)

	fmt.Println(singleResult)

	return singleResult.Err()
}

func (r repository) FindAccessToken(context context.Context, refreshToken string) (string, error) {
	return "", nil
}

func (r repository) FindUser(context context.Context, login string) (entity.User, error) {
	var user entity.User
	err := r.db.Collection(collection).FindOne(
		context,
		bson.M{
			"$or": []bson.M{
				bson.M{"login": login},
				bson.M{"email": login},
			},
		},
	).Decode(&user)
	return user, err
}

func (r repository) SaveRefreshToken(context context.Context, userID string, refreshTokens entity.RefreshTokens) error {

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = r.db.Collection(collection).UpdateOne(
		context,
		bson.M{
			"_id": id,
		},
		bson.M{
			"$push": bson.M{"refreshTokens": refreshTokens},
		},
	)

	return err
}
