// +build integration

package user

import (
	"context"
	"testing"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database
var r Repository

func TestUserRepositoryMain(t *testing.T) {
	db = test.Database(t)
	r = NewRepository(db)
}

func TestUserRepositoryCreate(t *testing.T) {
	id := primitive.NewObjectID()
	cases := []struct {
		name       string
		repository Repository
		entity     entity.User
		fail       bool
	}{
		{
			"success",
			r,
			entity.User{
				ID:           id,
				Login:        "test",
				Email:        "test@mailservice.com",
				PasswordHash: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9",
				Website:      "https://github.com",
				Token:        "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9",
				Banned:       false,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			false,
		},
		{
			"duplicate id",
			r,
			entity.User{
				ID:           id,
				Login:        "test2",
				Email:        "test2@mailservice.com",
				PasswordHash: "eyJ0eXAiOidfa1QiLCJhbGciOiJIUzI1NiJ9",
				Website:      "https://github.com",
				Token:        "eyJ0eXAiOfdad234fdasdg21wr424rggNiJ9",
				Banned:       false,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.repository.Create(context.TODO(), tt.entity)
			assert.Equal(t, tt.fail, err != nil)
		})
	}
}

func TestUserRepositoryCount(t *testing.T) {
	count, err := r.Count(context.TODO())
	assert.Nil(t, err)
	assert.NotEqual(t, count, 0)
}

func TestUserRepositoryFindByID(t *testing.T) {
	_, err := r.FindByID(context.TODO(), primitive.NewObjectID().Hex())
	assert.NotNil(t, err)
}
