package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/auth/mock"
	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/test"
)

func TestEndpoints(t *testing.T) {
	users := []entity.User{
		{
			ID:           100,
			PasswordHash: "$2a$10$Ln6XYtZOD.YfxJk/HFwVle7gFpE.dyWueCaLbsUhW6vtWbBGtFUyy",
			Login:        "user_100",
			Email:        "user_100@mail.com",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}
	sessions := []entity.Session{
		{
			ID:           1,
			UserID:       100,
			RefreshToken: "f49ac960-7cb5-11ea-aedc-acde48001122",
			Exp:          time.Now().Add(time.Hour * 1),
			IP:           "127.0.0.1",
			UserAgent:    "Insomnia",
			CreatedAt:    time.Now(),
		},
	}
	cases := []test.APITestCase{
		{
			Name:         "Login",
			Method:       http.MethodPost,
			URL:          "/auth/login",
			Body:         `{"login":"user_100", "password":"qwerty"}`,
			Header:       nil,
			WantStatus:   http.StatusOK,
			WantResponse: "",
		},
		{
			Name:         "Login - invalid login or password",
			Method:       http.MethodPost,
			URL:          "/auth/login",
			Body:         `{"login":"test", "password":"qwerty"}`,
			Header:       nil,
			WantStatus:   http.StatusInternalServerError,
			WantResponse: "",
		},
		{
			Name:         "Refresh JWT token",
			Method:       http.MethodPost,
			URL:          "/auth/refresh",
			Body:         `{"refresh_token":"f49ac960-7cb5-11ea-aedc-acde48001122"}`,
			Header:       test.MockAuthHeader(),
			WantStatus:   http.StatusOK,
			WantResponse: "",
		},
		{
			Name:         "Logout",
			Method:       http.MethodPost,
			URL:          "/auth/logout",
			Body:         `{"refresh_token":"f49ac960-7cb5-11ea-aedc-acde48001122"}`,
			Header:       test.MockAuthHeader(),
			WantStatus:   http.StatusOK,
			WantResponse: "",
		},
	}
	mockRepository := mock.NewRepository(users, sessions)
	service := NewService("test", mockRepository)
	router := test.MockRouter()
	NewHTTPHandler(router.Group(""), test.MockAuthHandler, service)
	for _, tc := range cases {
		test.Endpoint(t, router, tc)
	}
}
