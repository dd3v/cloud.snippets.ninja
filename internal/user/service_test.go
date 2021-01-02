package user

import (
	"context"
	"testing"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/user/mock"
	"github.com/stretchr/testify/assert"
)

var mockRepository Repository
var testService Service

var users = []entity.User{
	{
		ID:           100,
		PasswordHash: "hash_100",
		Login:        "user_100",
		Email:        "user_100@mail.com",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
	{
		ID:           200,
		PasswordHash: "hash_100",
		Login:        "user_200",
		Email:        "user_200@mail.com",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
	{
		ID:           300,
		PasswordHash: "hash_100",
		Login:        "user_300",
		Email:        "user_300@mail.com",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
}

func TestService(t *testing.T) {
	mockRepository := mock.NewRepository(users)
	testService = NewService(mockRepository)
}

func TestCount(t *testing.T) {
	count, err := testService.Count(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, count, len(users))
}

func TestCreate(t *testing.T) {
	request := CreateRequest{
		Login:          "test_user",
		Email:          "test_user@mailservice.com",
		Password:       "qwerty",
		RepeatPassword: "qwerty",
	}
	user, err := testService.Create(context.TODO(), request)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestUpdate(t *testing.T) {
	request := UpdateRequest{
		Website: "new_test_100.com",
	}
	_, err := testService.Update(context.TODO(), 100, request)
	assert.Nil(t, err)
}

func TestFindById(t *testing.T) {
	id := 100000
	user, err := testService.FindByID(context.TODO(), id)
	assert.NotNil(t, err)
	assert.NotEqual(t, user.ID, id)

	user, err = testService.FindByID(context.TODO(), 100)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, 100)
}

func TestDelete(t *testing.T) {
	err := testService.Delete(context.TODO(), 100)
	assert.Nil(t, err)
	_, err = testService.FindByID(context.TODO(), 100)
	assert.NotNil(t, err)
}
