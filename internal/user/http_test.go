package user

import (
	"net/http"
	"testing"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"github.com/dd3v/snippets.page.backend/internal/user/mock"
)

func TestUserEndpoint(t *testing.T) {

	users := []entity.User{{
		ID:           100,
		PasswordHash: "hash",
		Login:        "test_100",
		Email:        "test_100@gmail.com",
		Website:      "homepage.com",
		Banned:       false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}}
	cases := []test.APITestCase{
		//public
		{
			Name:         "create new user",
			Method:       http.MethodPost,
			URL:          "/users",
			Body:         `{"login":"test", "email": "test@gmail.com", "password": "qwerty", "repeat_password": "qwerty"}`,
			Header:       nil,
			WantStatus:   http.StatusCreated,
			WantResponse: `*test*`,
		},
		{
			Name:         "create new user - validation error",
			Method:       http.MethodPost,
			URL:          "/users",
			Body:         `{"login":"a", "email": "test@gmail.com", "password": "qwerty", "repeat_password": "qwerty"}`,
			Header:       nil,
			WantStatus:   http.StatusBadRequest,
			WantResponse: `*Validation error*`,
		},
		//auth token
		{
			Name:         "get user by id",
			Method:       http.MethodGet,
			URL:          "/users/100",
			Body:         "",
			Header:       test.MockAuthHeader(),
			WantStatus:   http.StatusOK,
			WantResponse: "",
		},
		{
			Name:         "get user by id - not found",
			Method:       http.MethodGet,
			URL:          "/users/101",
			Body:         "",
			Header:       test.MockAuthHeader(),
			WantStatus:   http.StatusNotFound,
			WantResponse: "",
		},
		{
			Name:         "update user",
			Method:       http.MethodPut,
			URL:          "/users/me",
			Body:         `{"website":"http://github.com"}`,
			Header:       test.MockAuthHeader(),
			WantStatus:   http.StatusOK,
			WantResponse: `*github.com*`,
		},
		{
			Name:         "user by id - unauthorized",
			Method:       http.MethodGet,
			URL:          "/users/100",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusUnauthorized,
			WantResponse: "",
		},
	}
	service := NewService(mock.NewRepository(users))
	router := test.MockRouter()
	NewHTTPHandler(router.Group(""), test.MockAuthHandler, service)
	for _, tc := range cases {
		test.Endpoint(t, router, tc)
	}
}
