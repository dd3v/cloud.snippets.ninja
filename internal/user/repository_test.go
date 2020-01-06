package user

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database
var user entity.User
var userRepository Repository

func TestMain(t *testing.T) {
	db = test.Database(t)
	userRepository = NewRepository(db)
	faker.SetGenerateUniqueValues(true)

}

func TestUserRepository_Create(t *testing.T) {
	err := userRepository.Create(user)
	assert.NoError(t, err)
}

func TestUserRepository_FindByID(t *testing.T) {
	result, err := userRepository.FindByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, result.ID, user.ID, "they should be equal")
}

func TestUserRepository_Update(t *testing.T) {
	newUser := user
	newUser.Login = "test_login"
	newUser.Banned = true
	err := userRepository.Update(newUser)
	assert.NoError(t, err)
	assert.NotEqual(t, user.Login, newUser.Login, "they should not be equal")
	assert.NotEqual(t, user.Banned, newUser.Banned, "they should not be equal")
}

func TestUserRepository_Delete(t *testing.T) {
	err := userRepository.Delete(user.ID)
	assert.NoError(t, err)
	_, err = userRepository.FindByID(user.ID)
	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), "mongo: no documents in result")
	}
}

func TestUserRepository_Count(t *testing.T) {
	count, err := userRepository.Count()
	assert.NoError(t, err)
	assert.Equal(t, count, int64(0))
}
