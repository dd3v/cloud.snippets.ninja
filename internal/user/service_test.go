package user

import (
	"context"
	"testing"

	"github.com/dd3v/snippets.page.backend/internal/user/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockRepository Repository
var testService Service

func TestMain(t *testing.T) {
	mockRepository := &mock.MemoryUserRepository{}
	testService = NewService(mockRepository)
}

func TestUserServiceCreate(t *testing.T) {
	createUserRequest := CreateRequest{
		Login:          "test",
		Email:          "test@mailservice.com",
		Password:       "qwerty",
		RepeatPassword: "qwerty",
	}
	user, err := testService.Create(context.TODO(), createUserRequest)
	assert.Equal(t, createUserRequest.Email, user.Email)
	assert.Nil(t, err)
	createUserRequest.Login = "error"
	_, err = testService.Create(context.TODO(), createUserRequest)
	assert.NotNil(t, err)
}

func TestUserServiceFind(t *testing.T) {
	users, err := testService.Find(context.TODO(), make(map[string]interface{}))
	assert.NotNil(t, users)
	assert.Nil(t, err)
}

func TestUserServiceCount(t *testing.T) {
	count, err := testService.Count(context.TODO())
	assert.Equal(t, count, 1)
	assert.Nil(t, err)
}

func TestUserServiceUpdate(t *testing.T) {
	createUserRequest := CreateRequest{
		Login:          "test",
		Email:          "test@mailservice.com",
		Password:       "qwerty",
		RepeatPassword: "qwerty",
	}
	user, err := testService.Create(context.TODO(), createUserRequest)
	assert.Equal(t, createUserRequest.Email, user.Email)
	assert.Nil(t, err)
	updateUserRequest := UpdateRequest{
		Website: "personalwebsite.com",
	}
	updatedUser, err := testService.Update(context.TODO(), user.ID.Hex(), updateUserRequest)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, updatedUser.ID)
}

func TestUserServiceFindByID(t *testing.T) {
	createUserRequest := CreateRequest{
		Login:          "test",
		Email:          "test@mailservice.com",
		Password:       "qwerty",
		RepeatPassword: "qwerty",
	}
	createdUser, err := testService.Create(context.TODO(), createUserRequest)
	result, err := testService.FindByID(context.TODO(), createdUser.ID.Hex())
	assert.Equal(t, createdUser.ID, result.ID)
	assert.Nil(t, err)
	_, err = testService.FindByID(context.TODO(), primitive.NewObjectID().Hex())
	assert.Equal(t, err, nil)
	_, err = testService.FindByID(context.TODO(), "error")
	assert.NotNil(t, err)
}

func TestUserServiceDelete(t *testing.T) {
	createUserRequest := CreateRequest{
		Login:          "test",
		Email:          "test@mailservice.com",
		Password:       "qwerty",
		RepeatPassword: "qwerty",
	}
	user, err := testService.Create(context.TODO(), createUserRequest)
	assert.Nil(t, err)
	err = testService.Delete(context.TODO(), user.ID.Hex())
	assert.Nil(t, err)
}
