package user

import (
	"context"
	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"net/http"
	"testing"
)

func TestHTTP_Create(t *testing.T) {
	cases := []struct {
		name       string
		request    test.APITestCase
		repository Repository
	}{
		{
			name: "create success",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/users",
				Body:         `{"login": "dd3v", "email": "dd3v@gmail.com", "password": "qwerty", "repeat_password": "qwerty"}`,
				Header:       nil,
				WantStatus:   http.StatusCreated,
				WantResponse: "*dd3v*",
			},
			repository: RepositoryMock{
				CreateFn: func(ctx context.Context, user entity.User) (entity.User, error) {
					return entity.User{
						ID:        1,
						Password:  "$2a$10$ceGJobOZUCIVM72m9fMVZO.NjjQcaadIhJnQEE7Cdq/QuBze9yZAq",
						Login:     "dd3v",
						Email:     "dd3v@gmail.com",
						CreatedAt: test.Time(2021),
						UpdatedAt: test.Time(2021),
					}, nil
				},
				ExistsFn: func(ctx context.Context, field string, value string) (bool, error) {
					return false, nil
				},
			},
		},
		{
			name: "validation error",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/users",
				Body:         `{"email": "dd3v@gmail.com", "password": "qwerty", "repeat_password": "qwerty"}`,
				Header:       nil,
				WantStatus:   http.StatusBadRequest,
				WantResponse: "",
			},
			repository: RepositoryMock{
				CreateFn: func(ctx context.Context, user entity.User) (entity.User, error) {
					return entity.User{}, nil
				},
				ExistsFn: func(ctx context.Context, field string, value string) (bool, error) {
					return false, nil
				},
			},
		},
		{
			name: "validation error, email or login",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/users",
				Body:         `{"email": "dd3v@gmail.com", "password": "qwerty", "repeat_password": "qwerty"}`,
				Header:       nil,
				WantStatus:   http.StatusBadRequest,
				WantResponse: "",
			},
			repository: RepositoryMock{
				ExistsFn: func(ctx context.Context, field string, value string) (bool, error) {
					return true, nil
				},
			},
		},
		{
			name: "repository error",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/users",
				Body:         `{"login":"dd3v", "email": "dd3v@gmail.com", "password": "qwerty", "repeat_password": "qwerty"}`,
				Header:       nil,
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repository: RepositoryMock{
				ExistsFn: func(ctx context.Context, field string, value string) (bool, error) {
					return true, repositoryMockErr
				},
			},
		},
	}
	for _, tc := range cases {
		router := test.MockRouter()
		service := NewService(tc.repository)
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, tc.name, router, tc.request)
	}
}

func TestHTTP_Me(t *testing.T) {
	var cases = []struct {
		name       string
		request    test.APITestCase
		repository Repository
	}{
		{
			name: "unauthorized",
			request: test.APITestCase{
				Method:       http.MethodGet,
				URL:          "/users/me",
				Body:         "",
				Header:       nil,
				WantStatus:   http.StatusUnauthorized,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.User, error) {
					return entity.User{
						ID:        1,
						Password:  "$2a$10$ceGJobOZUCIVM72m9fMVZO.NjjQcaadIhJnQEE7Cdq/QuBze9yZAq",
						Login:     "dd3v",
						Email:     "dd3v@gmail.com",
						CreatedAt: test.Time(2021),
						UpdatedAt: test.Time(2021),
					}, nil
				},
			},
		},
		{
			name: "success",
			request: test.APITestCase{
				Method:       http.MethodGet,
				URL:          "/users/me",
				Body:         "",
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusOK,
				WantResponse: "*dd3v*",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.User, error) {
					return entity.User{
						ID:        1,
						Password:  "$2a$10$ceGJobOZUCIVM72m9fMVZO.NjjQcaadIhJnQEE7Cdq/QuBze9yZAq",
						Login:     "dd3v",
						Email:     "dd3v@gmail.com",
						CreatedAt: test.Time(2021),
						UpdatedAt: test.Time(2021),
					}, nil
				},
			},
		},
	}
	for _, tc := range cases {
		router := test.MockRouter()
		service := NewService(tc.repository)
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, tc.name, router, tc.request)
	}
}
