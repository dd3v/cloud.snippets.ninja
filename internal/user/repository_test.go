package user

import (
	"github.com/dd3v/snippets.page.backend/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database
var user entity.User
var userRepository Repository

// func TestMain(t *testing.T) {
// 	db = test.Database(t)
// 	userRepository = NewRepository(db)
// 	faker.SetGenerateUniqueValues(true)

// }

// func TestUserRepository_Create(t *testing.T) {
// 	err := userRepository.Create(user)
// 	assert.NoError(t, err)
// }

// func TestUserRepository_FindByID(t *testing.T) {
// 	result, err := userRepository.FindByID(user.ID)
// 	assert.NoError(t, err)

// 	faker.SetGenerateUniqueValues(true)
// }

// func TestCreate(t *testing.T) {

// 	user := entity.User{
// 		ID:    primitive.NewObjectID(),
// 		Login: "sdfsdfsdf",
// 	}

// 	err := userRepository.Create(context.TODO(), user)
// 	assert.NoError(t, err)
// }

// func TestFinByID(t *testing.T) {
// 	result, err := userRepository.FindByID(context.TODO(), user.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, result.ID, user.ID, "user ID should be equal")
// }

// func TestUpdate(t *testing.T) {
// 	newUser := user
// 	newUser.Login = "test_login"
// 	newUser.Banned = true
// 	err := userRepository.Update(context.TODO(), newUser)
// 	assert.NoError(t, err)
// 	assert.NotEqual(t, user.Login, newUser.Login, "they should not be equal")
// 	assert.NotEqual(t, user.Banned, newUser.Banned, "they should not be equal")
// }

// func TestDelete(t *testing.T) {
// 	err := userRepository.Delete(context.TODO(), user.ID)
// 	assert.NoError(t, err)
// 	_, err = userRepository.FindByID(context.TODO(), user.ID)
// 	if assert.NotNil(t, err) {
// 		assert.Equal(t, err.Error(), "mongo: no documents in result")
// 	}
// }

// func TestCount(t *testing.T) {
// 	count, err := userRepository.Count(context.TODO(), map[string]string{})
// 	assert.NoError(t, err)
// 	assert.Equal(t, count, int64(0))
// }
