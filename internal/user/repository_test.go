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
				ID:           primitive.NewObjectID(),
				Login:        "test_case_0",
				Email:        "test_case_0@mailservice.com",
				PasswordHash: "$2a$10$BTAOpHA5j62f56UlEWY3MuaUmY967Pm1kQm3nMCer0wEN2YQGBL8S",
				Website:      "https://github.com",
				Token:        "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9",
				Banned:       false,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			false,
		},
		{
			"duplicate email",
			r,
			entity.User{
				ID:           primitive.NewObjectID(),
				Login:        "test_case_2",
				Email:        "test_case_0@mailservice.com",
				PasswordHash: "$2a$10$BTAOpHA5j62f56UlEWY3MuaUmY967Pm1kQm3nMCer0wEN2YQGBL8S",
				Website:      "https://github.com",
				Token:        "eyJ0eXAiOfdad234fdasdg21wr424rggNiJ9",
				Banned:       false,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			true,
		},
		{
			"duplicate login",
			r,
			entity.User{
				ID:           primitive.NewObjectID(),
				Login:        "test_case_0",
				Email:        "test_case_4@mailservice.com",
				PasswordHash: "$2a$10$BTAOpHA5j62f56UlEWY3MuaUmY967Pm1kQm3nMCer0wEN2YQGBL8S",
				Website:      "https://github.com",
				Token:        "eyJ0eXAiOfdad234fdasdg21wr424rggNiJ9",
				Banned:       false,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			true,
		},
		{
			"duplicate login or email",
			r,
			entity.User{
				ID:           primitive.NewObjectID(),
				Login:        "test_case_0",
				Email:        "test_case_0@mailservice.com",
				PasswordHash: "$2a$10$BTAOpHA5j62f56UlEWY3MuaUmY967Pm1kQm3nMCer0wEN2YQGBL8S",
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
	user := entity.User{
		ID:           primitive.NewObjectID(),
		Login:        "test_case_create_5",
		Email:        "test_case_create_5@mailservice.com",
		PasswordHash: "$2a$10$BTAOpHA5j62f56UlEWY3MuaUmY967Pm1kQm3nMCer0wEN2YQGBL8S",
		Website:      "https://github.com",
		Token:        "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9",
		Banned:       false,
		CreatedAt:    time.Now().UTC().Round(time.Second),
		UpdatedAt:    time.Now().UTC().Round(time.Second),
	}
	err = r.Create(context.TODO(), user)
	assert.Nil(t, err)
	result, err := r.FindByID(context.TODO(), user.ID.Hex())
	assert.Equal(t, user, result)
}

func TestUserRepositoryDelete(t *testing.T) {
	user := entity.User{
		ID:           primitive.NewObjectID(),
		Login:        "test_case_create_6",
		Email:        "test_case_create_6@mailservice.com",
		PasswordHash: "$2a$10$BTAOpHA5j62f56UlEWY3MuaUmY967Pm1kQm3nMCer0wEN2YQGBL8S",
		Website:      "https://github.com",
		Token:        "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9",
		Banned:       false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := r.Create(context.TODO(), user)
	assert.Nil(t, err)
	err = r.Delete(context.TODO(), user.ID.Hex())
	assert.Nil(t, err)
	_, err = r.FindByID(context.TODO(), user.ID.Hex())
	assert.NotNil(t, err)
}

func TestUserRepositoryUpdate(t *testing.T) {
	user := entity.User{
		ID:           primitive.NewObjectID(),
		Login:        "test_case_create_7",
		Email:        "test_case_create_7@mailservice.com",
		PasswordHash: "$2a$10$BTAOpHA5j62f56UlEWY3MuaUmY967Pm1kQm3nMCer0wEN2YQGBL8S",
		Website:      "https://github.com",
		Token:        "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9",
		Banned:       false,
		CreatedAt:    time.Now().UTC().Round(time.Second),
		UpdatedAt:    time.Now().UTC().Round(time.Second),
	}
	err := r.Create(context.TODO(), user)
	assert.Nil(t, err)
	user.Website = "http://facebook.com"
	user.Banned = true
	user.UpdatedAt = time.Now().UTC().Round(time.Second)
	err = r.Update(context.TODO(), user)
	assert.Nil(t, err)
	result, err := r.FindByID(context.TODO(), user.ID.Hex())
	assert.Nil(t, err)
	assert.Equal(t, user, result)
}
