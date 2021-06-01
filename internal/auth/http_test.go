package auth

import (
	"context"
	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"net/http"
	"testing"
	"time"
)

func TestHTTP_Login(t *testing.T) {
	cases := []struct {
		request        test.APITestCase
		repositoryMock Repository
	}{
		{
			request: test.APITestCase{
				Name:         "user can login",
				Method:       http.MethodPost,
				URL:          "/auth/login",
				Body:         `{"login":"dd3v", "password":"qwerty"}`,
				Header:       nil,
				WantStatus:   http.StatusOK,
				WantResponse: "",
			},
			repositoryMock: RepositoryMock{
				GetUserByLoginOrEmailFn: func(ctx context.Context, value string) (entity.User, error) {
					return entity.User{
						ID:           1,
						PasswordHash: "$2a$10$ubN1SU6RUOjlbQiHObqy7.bgK08Gl/YNWxTSrqhkTsvtnsh1nFzDO",
						Login:        "dd3v",
						Email:        "test@test.com",
						CreatedAt:    test.Time(2020),
						UpdatedAt:    test.Time(2020),
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
			request: test.APITestCase{
				Name:         "validation error",
				Method:       http.MethodPost,
				URL:          "/auth/login",
				Body:         `{"login":"dd3v", "password":"test"}`,
				Header:       nil,
				WantStatus:   http.StatusBadRequest,
				WantResponse: "",
			},
			repositoryMock: RepositoryMock{
				GetUserByLoginOrEmailFn: func(ctx context.Context, value string) (entity.User, error) {
					return entity.User{
						ID:           1,
						PasswordHash: "$2a$10$ubN1SU6RUOjlbQiHObqy7.bgK08Gl/YNWxTSrqhkTsvtnsh1nFzDO",
						Login:        "dd3v",
						Email:        "test@test.com",
						CreatedAt:    test.Time(2020),
						UpdatedAt:    test.Time(2020),
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
			request: test.APITestCase{
				Name:         "user repository error",
				Method:       http.MethodPost,
				URL:          "/auth/login",
				Body:         `{"login":"dd3v", "password":"qwerty"}`,
				Header:       nil,
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repositoryMock: RepositoryMock{
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
			request: test.APITestCase{
				Name:         "session repository error",
				Method:       http.MethodPost,
				URL:          "/auth/login",
				Body:         `{"login":"dd3v", "password":"qwerty"}`,
				Header:       nil,
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repositoryMock: RepositoryMock{
				GetUserByLoginOrEmailFn: func(ctx context.Context, value string) (entity.User, error) {
					return entity.User{
						ID:           1,
						PasswordHash: "$2a$10$ubN1SU6RUOjlbQiHObqy7.bgK08Gl/YNWxTSrqhkTsvtnsh1nFzDO",
						Login:        "dd3v",
						Email:        "test@test.com",
						CreatedAt:    test.Time(2020),
						UpdatedAt:    test.Time(2020),
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
		service := NewService("jwt_test_key", tc.repositoryMock)
		router := test.MockRouter()
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, router, tc.request)
	}
}

func TestHTTP_RefreshToken(t *testing.T) {
	cases := []struct {
		request        test.APITestCase
		repositoryMock Repository
	}{
		{
			request: test.APITestCase{
				Name:         "user can refresh token",
				Method:       http.MethodPost,
				URL:          "/auth/refresh",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       nil,
				WantStatus:   http.StatusOK,
				WantResponse: "",
			},
			repositoryMock: RepositoryMock{
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
			request: test.APITestCase{
				Name:         "refresh token expired",
				Method:       http.MethodPost,
				URL:          "/auth/refresh",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       nil,
				WantStatus:   http.StatusForbidden,
				WantResponse: "",
			},
			repositoryMock: RepositoryMock{
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
			request: test.APITestCase{
				Name:         "refresh by non-existent",
				Method:       http.MethodPost,
				URL:          "/auth/refresh",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       nil,
				WantStatus:   http.StatusForbidden,
				WantResponse: "",
			},
			repositoryMock: RepositoryMock{
				GetSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) (entity.Session, error) {
					return entity.Session{}, repositoryMockErr
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
			request: test.APITestCase{
				Name:         "repository error",
				Method:       http.MethodPost,
				URL:          "/auth/refresh",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       nil,
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repositoryMock: RepositoryMock{
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
		service := NewService("jwt_test_key", tc.repositoryMock)
		router := test.MockRouter()
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, router, tc.request)
	}
}

func TestHTTP_Logout(t *testing.T) {
	cases := []struct {
		request        test.APITestCase
		repositoryMock Repository
	}{
		{
			request: test.APITestCase{
				Name:         "user can logout",
				Method:       http.MethodPost,
				URL:          "/auth/logout",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusOK,
				WantResponse: "",
			},
			repositoryMock: RepositoryMock{
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
			},
		},
		{
			request: test.APITestCase{
				Name:         "unauthorized request",
				Method:       http.MethodPost,
				URL:          "/auth/logout",
				Body:         `{"refresh_token":"d5586222-c306-11eb-96c1-acde48001122"}`,
				Header:       nil,
				WantStatus:   http.StatusUnauthorized,
				WantResponse: "",
			},
			repositoryMock: RepositoryMock{
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
			},
		},
	}

	for _, tc := range cases {
		service := NewService("jwt_test_key", tc.repositoryMock)
		router := test.MockRouter()
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, router, tc.request)
	}
}
