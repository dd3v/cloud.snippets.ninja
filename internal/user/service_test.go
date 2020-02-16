package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dd3v/snippets.page.backend/internal/test/mock"
)

var mockRepository Repository
var testService Service

func TestMain(t *testing.T) {

	mockRepository = &mock.MockUserRepository{}
	testService = NewService(mockRepository)

	fmt.Println("service testing")

}

func TestCount(t *testing.T) {
	count, err := testService.Count(context.TODO())
	assert.Equal(t, count, 0)
	assert.Nil(t, err)
}

func TestCreateUser(t *testing.T) {

	createUserRequest := CreateRequest{
		Login:          "test",
		Email:          "test@mailservice.com",
		Password:       "qwerty",
		RepeatPassword: "qwerty",
	}

	user, err := testService.CreateUser(context.TODO(), createUserRequest)
	assert.Equal(t, createUserRequest.Email, user.Email)
	assert.Nil(t, err)

	createUserRequest.Login = "error"
	_, err = testService.CreateUser(context.TODO(), createUserRequest)

}

func TestFindByID(t *testing.T) {

	id := "234123424234"

	user, err := testService.FindByID(context.TODO(), id)

	assert.Equal(t, id, user.ID)
	assert.Nil(t, err)

}
