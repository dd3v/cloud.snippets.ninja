package auth

import (
	"context"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
)

func TestAuthService_Login(t *testing.T) {
	type args struct {
		auth authCredentials
	}

	cases := []struct {
		name       string
		args       args
		repository RepositoryMock
		wantData   bool
		wantErr    error
	}{
		{
			"user can successfully login",
			args{
				authCredentials{
					User:      "test",
					Password:  "qwerty",
					UserAgent: "test",
					IP:        "127.0.0.1",
				},
			},
			RepositoryMock{
				GetUserByLoginOrEmailFn: func(ctx context.Context, login string) (entity.User, error) {
					return entity.User{
						ID:           1,
						PasswordHash: "$2a$10$ubN1SU6RUOjlbQiHObqy7.bgK08Gl/YNWxTSrqhkTsvtnsh1nFzDO",
						Login:        "test",
						Email:        "",
						CreatedAt:    test.Time(2020),
						UpdatedAt:    test.Time(2020),
					}, nil
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
				DeleteSessionByUserIDAndUserAgentFn: func(ctx context.Context, userID int, userAgent string) error {
					return nil
				},
			},
			true,
			nil,
		},
		{
			"invalid login or password",
			args{
				authCredentials{
					User:      "test",
					Password:  "123123",
					UserAgent: "test",
					IP:        "127.0.0.1",
				},
			},
			RepositoryMock{
				GetUserByLoginOrEmailFn: func(ctx context.Context, login string) (entity.User, error) {
					return entity.User{
						ID:           0,
						PasswordHash: "$2a$10$ubN1SU6RUOjlbQiHObqy7.bgK08Gl/YNWxTSrqhkTsvtnsh1nFzDO",
						Login:        "test",
						Email:        "",
						CreatedAt:    test.Time(2020),
						UpdatedAt:    test.Time(2020),
					}, nil
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
			},
			false,
			authErr,
		},
		{
			"error when try to find user by login or password",
			args{
				authCredentials{
					User:      "test",
					Password:  "qwerty",
					UserAgent: "test",
					IP:        "127.0.0.1",
				},
			},
			RepositoryMock{
				GetUserByLoginOrEmailFn: func(ctx context.Context, login string) (entity.User, error) {
					return entity.User{}, repositoryMockErr
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
			},
			false,
			repositoryMockErr,
		},
		{
			"session could not be created",
			args{
				authCredentials{
					User:      "test",
					Password:  "qwerty",
					UserAgent: "test",
					IP:        "127.0.0.1",
				},
			},
			RepositoryMock{
				GetUserByLoginOrEmailFn: func(ctx context.Context, login string) (entity.User, error) {
					return entity.User{
						ID:           1,
						PasswordHash: "$2a$10$ubN1SU6RUOjlbQiHObqy7.bgK08Gl/YNWxTSrqhkTsvtnsh1nFzDO",
						Login:        "test",
						Email:        "",
						CreatedAt:    test.Time(2020),
						UpdatedAt:    test.Time(2020),
					}, nil
				},
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return repositoryMockErr
				},
				DeleteSessionByUserIDAndUserAgentFn: func(ctx context.Context, userID int, userAgent string) error {
					return nil
				},
			},
			false,
			repositoryMockErr,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService("jwt_test_key", tc.repository)
			tokenPair, err := service.Login(context.Background(), tc.args.auth)
			assert.Equal(t, tc.wantData, tokenPair.AccessToken != "")
			assert.Equal(t, tc.wantData, tokenPair.RefreshToken != "")
			assert.IsType(t, tc.wantErr, err)
		})
	}
}
func TestAuthService_RefreshAccessToken(t *testing.T) {
	type args struct {
		refreshCredentials refreshCredentials
	}
	cases := []struct {
		name       string
		repository Repository
		args       args
		wantData   bool
		wantErr    error
	}{
		{
			"user can successfully refresh token and get new token pair",
			RepositoryMock{
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
				GetSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) (entity.Session, error) {
					return entity.Session{
						ID:           1,
						UserID:       1,
						RefreshToken: "07c40c34-c07d-11eb-a218-acde48001122",
						Exp:          time.Now().Add(time.Hour),
						IP:           "127.0.0.1",
						UserAgent:    "Insomnia",
						CreatedAt:    test.Time(2020),
					}, nil
				},
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
			},
			args{
				refreshCredentials{
					RefreshToken: "07c40c34-c07d-11eb-a218-acde48001122",
					UserAgent:    "Insomnia",
					IP:           "127.0.0.1",
				}},
			true,
			nil,
		},
		{
			"session already expired",
			RepositoryMock{
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
				GetSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) (entity.Session, error) {
					return entity.Session{
						ID:           1,
						UserID:       1,
						RefreshToken: "07c40c34-c07d-11eb-a218-acde48001122",
						Exp:          time.Now().Add(time.Duration(-10) * time.Minute),
						IP:           "127.0.0.1",
						UserAgent:    "Insomnia",
						CreatedAt:    test.Time(2020),
					}, nil
				},
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
				DeleteSessionByUserIDFn: func(ctx context.Context, userID int) (int64, error) {
					return 10, nil
				},
			},
			args{
				refreshCredentials{
					RefreshToken: "07c40c34-c07d-11eb-a218-acde48001122",
					UserAgent:    "Insomnia",
					IP:           "127.0.0.1",
				}},
			false,
			expiredSessionErr,
		},
		{
			"session could not be created",
			RepositoryMock{
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return createSessionErr
				},
				GetSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) (entity.Session, error) {
					return entity.Session{
						ID:           1,
						UserID:       1,
						RefreshToken: "07c40c34-c07d-11eb-a218-acde48001122",
						Exp:          time.Now().Add(time.Hour),
						IP:           "127.0.0.1",
						UserAgent:    "Insomnia",
						CreatedAt:    test.Time(2020),
					}, nil
				},
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
			},
			args{
				refreshCredentials{
					RefreshToken: "07c40c34-c07d-11eb-a218-acde48001122",
					UserAgent:    "Insomnia",
					IP:           "127.0.0.1",
				}},
			false,
			createSessionErr,
		},
		{
			"refresh token already expired",
			RepositoryMock{
				CreateSessionFn: func(ctx context.Context, session entity.Session) error {
					return nil
				},
				GetSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) (entity.Session, error) {
					return entity.Session{
						ID:           1,
						UserID:       1,
						RefreshToken: "07c40c34-c07d-11eb-a218-acde48001122",
						Exp:          time.Now(),
						IP:           "127.0.0.1",
						UserAgent:    "Insomnia",
						CreatedAt:    test.Time(2020),
					}, nil
				},
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
			},
			args{
				refreshCredentials{
					RefreshToken: "07c40c34-c07d-11eb-a218-acde48001122",
					UserAgent:    "Insomnia",
					IP:           "127.0.0.1",
				}},
			false,
			expiredSessionErr,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService("jwt_test_key", tc.repository)
			token, err := service.Refresh(context.Background(), tc.args.refreshCredentials)
			assert.Equal(t, tc.wantData, token.AccessToken != "")
			assert.Equal(t, tc.wantData, token.RefreshToken != "")
			assert.IsType(t, tc.wantErr, err)
		})
	}
}

func TestAuthService_Logout(t *testing.T) {
	type args struct {
	}

	cases := []struct {
		name       string
		args       args
		repository Repository
		wantErr    error
	}{
		{
			"user can successfully logout",
			args{},
			RepositoryMock{
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return nil
				},
			},
			nil,
		},
		{
			"repository error",
			args{},
			RepositoryMock{
				DeleteSessionByRefreshTokenFn: func(ctx context.Context, refreshToken string) error {
					return repositoryMockErr
				},
			},
			nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService("jwt_test", tc.repository)
			err := service.Logout(context.Background(), "d5586222-c306-11eb-96c1-acde48001122")
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
