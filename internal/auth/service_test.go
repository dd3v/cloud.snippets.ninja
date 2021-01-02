package auth

import (
	"context"
	"testing"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/auth/mock"
	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

var mockRepository Repository
var testService Service
var sessions = []entity.Session{
	{
		ID:           1,
		UserID:       100,
		RefreshToken: "f49ac960-7cb5-11ea-aedc-acde48001122",
		Exp:          time.Now(),
		IP:           "127.0.0.1",
		UserAgent:    "Insomnia",
		CreatedAt:    time.Now(),
	},
}
var users = []entity.User{
	{
		ID:           100,
		PasswordHash: "$2a$10$Ln6XYtZOD.YfxJk/HFwVle7gFpE.dyWueCaLbsUhW6vtWbBGtFUyy",
		Login:        "user_100",
		Email:        "user_100@mail.com",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
}

func TestMain(t *testing.T) {
	mockRepository = mock.NewRepository(users, sessions)
	testService = NewService("sdfsdfsdf", mockRepository)
}

func TestRefreshTokenSessionAlreadyExpired(t *testing.T) {
	request := RefreshRequest{
		RefreshToken: "f49ac960-7cb5-11ea-aedc-acde48001122",
	}
	tokenPair, err := testService.Refresh(context.TODO(), request.RefreshToken)
	assert.Equal(t, entity.TokenPair{}, tokenPair)
	assert.NotNil(t, err)
}

func TestLogin(t *testing.T) {
	request := LoginRequest{
		Login:    "user_100",
		Password: "qwerty",
	}
	tokenPair, err := testService.Login(context.TODO(), request)
	assert.Nil(t, err)
	assert.NotEmpty(t, tokenPair.AccessToken)
	assert.NotEmpty(t, tokenPair.RefreshToken)
}

func TestLogout(t *testing.T) {
	err := testService.Logout(context.TODO(), "f49ac960-7cb5-11ea-aedc-acde48001122")
	assert.Nil(t, err)
}
