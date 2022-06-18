package auth

import (
	"context"
	"errors"

	"github.com/dd3v/cloud.snippets.ninja/internal/entity"
)

var repositoryMockErr = errors.New("error repository")

type RepositoryMock struct {
	GetUserByLoginOrEmailFn             func(ctx context.Context, value string) (entity.User, error)
	CreateSessionFn                     func(ctx context.Context, session entity.Session) error
	GetSessionByRefreshTokenFn          func(ctx context.Context, refreshToken string) (entity.Session, error)
	DeleteSessionByRefreshTokenFn       func(ctx context.Context, refreshToken string) error
	DeleteSessionByUserIDAndUserAgentFn func(ctx context.Context, userID int, userAgent string) error
	DeleteSessionByUserIDFn             func(ctx context.Context, userID int) (int64, error)
}

func NewRepositoryMock() Repository {
	return RepositoryMock{}
}

func (r RepositoryMock) GetUserByLoginOrEmail(ctx context.Context, value string) (entity.User, error) {
	return r.GetUserByLoginOrEmailFn(ctx, value)
}

func (r RepositoryMock) CreateSession(ctx context.Context, session entity.Session) error {
	return r.CreateSessionFn(ctx, session)
}

func (r RepositoryMock) GetSessionByRefreshToken(ctx context.Context, refreshToken string) (entity.Session, error) {
	return r.GetSessionByRefreshTokenFn(ctx, refreshToken)
}

func (r RepositoryMock) DeleteSessionByRefreshToken(ctx context.Context, refreshToken string) error {
	return r.DeleteSessionByRefreshTokenFn(ctx, refreshToken)
}

func (r RepositoryMock) DeleteSessionByUserIDAndUserAgent(ctx context.Context, userID int, userAgent string) error {
	return r.DeleteSessionByUserIDAndUserAgentFn(ctx, userID, userAgent)
}

func (r RepositoryMock) DeleteSessionByUserID(ctx context.Context, userID int) (int64, error) {
	return r.DeleteSessionByUserIDFn(ctx, userID)
}
