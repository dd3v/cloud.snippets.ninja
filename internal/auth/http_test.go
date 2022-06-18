package auth

import (
	"context"
	"database/sql"
	"net/http"
	"testing"
	"time"

	"github.com/dd3v/cloud.snippets.ninja/internal/entity"
	"github.com/dd3v/cloud.snippets.ninja/internal/test"
	"github.com/dd3v/cloud.snippets.ninja/pkg/log"
)

func TestHTTP_Login(t *testing.T) {
	cases := []struct {
		name       string
		request    test.APITestCase
		repository Repository
	}{
		{
			name: "user can login",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/auth/login",
				Body:         `{"login":"dd3v", "password":"qwerty"}`,
				Header:       nil,
				WantStatus:   http.StatusOK,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetUserByLoginOrEmailFn: func(ctx context.Context, value string) (entity.User, error) {
					return entity.User{
						ID:        1,
						Password:  "$2a$10$ubN1SU6RUOjlbQiHObqy7.bgK08Gl/YNWxTSrqhkTsvtnsh1nFzDO",
						Login:     "dd3v",
						Email:     "test@test.com",
						CreatedAt: test.Time(2020),
						UpdatedAt: test.Time(2020),
					}, nil
				},
				DeleteSessionByUserIDAndUserAgentFn: func(ctx context.Context, userID int, userAgent string) error {
					return nil
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
			},
		},
		{
			name: "validation error",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/auth/login",
				Body:         `{"login":"dd3v", "password":"test"}`,
				Header:       nil,
				WantStatus:   http.StatusBadRequest,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetUserByLoginOrEmailFn: func(ctx context.Context, value string) (entity.User, error) {
					return entity.User{
						ID:        1,
						Password:  "$2a$10$ubN1SU6RUOjlbQiHObqy7.bgK08Gl/YNWxTSrqhkTsvtnsh1nFzDO",
						Login:     "dd3v",
						Email:     "test@test.com",
						CreatedAt: test.Time(2020),
						UpdatedAt: test.Time(2020),
					}, nil
				},
				DeleteSessionByUserIDAndUserAgentFn: func(ctx context.Context, userID int, userAgent string) error {
					return nil
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
			},
		},
		{
			name: "user repository error",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/auth/login",
				Body:         `{"login":"dd3v", "password":"qwerty"}`,
				Header:       nil,
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetUserByLoginOrEmailFn: func(ctx context.Context, value string) (entity.User, error) {
					return entity.User{}, repositoryMockErr
				},
				DeleteSessionByUserIDAndUserAgentFn: func(ctx context.Context, userID int, userAgent string) error {
					return nil
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
			},
		},
		{
			name: "session repository error",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/auth/login",
				Body:         `{"login":"dd3v", "password":"qwerty"}`,
				Header:       nil,
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetUserByLoginOrEmailFn: func(ctx context.Context, value string) (entity.User, error) {
					return entity.User{
						ID:        1,
						Password:  "$2a$10$ubN1SU6RUOjlbQiHObqy7.bgK08Gl/YNWxTSrqhkTsvtnsh1nFzDO",
						Login:     "dd3v",
						Email:     "test@test.com",
						CreatedAt: test.Time(2020),
						UpdatedAt: test.Time(2020),
					}, nil
				},
				DeleteSessionByUserIDAndUserAgentFn: func(ctx context.Context, userID int, userAgent string) error {
					return nil
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return repositoryMockErr
				},
			},
		},
	}

	for _, tc := range cases {
		logger, _ := log.NewForTests()
		service := NewService("jwt_test_key", tc.repository, logger)
		router := test.MockRouter()
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, tc.name, router, tc.request)
	}
}

func TestHTTP_RefreshToken(t *testing.T) {
	cases := []struct {
		name       string
		request    test.APITestCase
		repository Repository
	}{
		{
			name: "user can refresh token",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/auth/refresh",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       nil,
				WantStatus:   http.StatusOK,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) (entity.Session, error) {
					return entity.Session{
						ID:           1,
						UserID:       1,
						RefreshToken: "d5586222-c306-11eb-96c1-acde48001122",
						Exp:          time.Now().Add(time.Hour),
						IP:           "127.0.0.1",
						UserAgent:    "test",
						CreatedAt:    test.Time(2020),
					}, nil
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
				DeleteSessionByUserIDFn: func(ctx context.Context, userID int) (int64, error) {
					return 0, nil
				},
			},
		},
		{
			name: "refresh token expired",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/auth/refresh",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       nil,
				WantStatus:   http.StatusForbidden,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) (entity.Session, error) {
					return entity.Session{
						ID:           1,
						UserID:       1,
						RefreshToken: "d5586222-c306-11eb-96c1-acde48001122",
						Exp:          time.Now(),
						IP:           "127.0.0.1",
						UserAgent:    "test",
						CreatedAt:    test.Time(2020),
					}, nil
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
				DeleteSessionByUserIDFn: func(ctx context.Context, userID int) (int64, error) {
					return 0, nil
				},
			},
		},
		{
			name: "refresh by non-existent",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/auth/refresh",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       nil,
				WantStatus:   http.StatusForbidden,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) (entity.Session, error) {
					return entity.Session{}, sql.ErrNoRows
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
				DeleteSessionByUserIDFn: func(ctx context.Context, userID int) (int64, error) {
					return 0, nil
				},
			},
		},
		{
			name: "repository error",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/auth/refresh",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       nil,
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) (entity.Session, error) {
					return entity.Session{
						ID:           1,
						UserID:       1,
						RefreshToken: "d5586222-c306-11eb-96c1-acde48001122",
						Exp:          time.Now().Add(time.Hour),
						IP:           "127.0.0.1",
						UserAgent:    "test",
						CreatedAt:    test.Time(2020),
					}, nil
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return repositoryMockErr
				},
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
				DeleteSessionByUserIDFn: func(ctx context.Context, userID int) (int64, error) {
					return 0, nil
				},
			},
		},
	}

	for _, tc := range cases {
		logger, _ := log.NewForTests()
		service := NewService("jwt_test_key", tc.repository, logger)
		router := test.MockRouter()
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, tc.name, router, tc.request)
	}
}

func TestHTTP_Logout(t *testing.T) {
	cases := []struct {
		name       string
		request    test.APITestCase
		repository Repository
	}{
		{
			name: "user can logout",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/auth/logout",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusOK,
				WantResponse: "",
			},
			repository: RepositoryMock{
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
			},
		},
		{
			name: "unauthorized request",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/auth/logout",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       nil,
				WantStatus:   http.StatusUnauthorized,
				WantResponse: "",
			},
			repository: RepositoryMock{
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
			},
		},
	}

	for _, tc := range cases {
		logger, _ := log.NewForTests()
		service := NewService("jwt_test_key", tc.repository, logger)
		router := test.MockRouter()
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, tc.name, router, tc.request)
	}
}
