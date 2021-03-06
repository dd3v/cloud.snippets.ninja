package auth

import (
	"context"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/dbcontext"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

//Repository - ...
type Repository interface {
	FindUserByLoginOrEmail(ctx context.Context, login string) (entity.User, error)
	CreateSession(ctx context.Context, session entity.Session) error
	FindSessionByRefreshToken(ctx context.Context, refreshToken string) (entity.Session, error)
	DeleteSessionByRefreshToken(ctx context.Context, refreshToken string) error
}

type repository struct {
	db *dbcontext.DB
}

//NewRepository - ...
func NewRepository(db *dbcontext.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) FindUserByLoginOrEmail(ctx context.Context, value string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().From("users").Where(dbx.HashExp{"login": value}).OrWhere(dbx.HashExp{"email": value}).One(&user)
	return user, err
}

func (r repository) CreateSession(ctx context.Context, session entity.Session) error {
	return r.db.With(ctx).Model(&session).Insert()
}

func (r repository) FindSessionByRefreshToken(ctx context.Context, refreshToken string) (entity.Session, error) {
	var session entity.Session
	err := r.db.With(ctx).Select().Where(&dbx.HashExp{"refresh_token": refreshToken}).One(&session)
	return session, err
}

func (r repository) DeleteSessionByRefreshToken(ctx context.Context, refreshToken string) error {
	session, err := r.FindSessionByRefreshToken(ctx, refreshToken)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&session).Delete()
}
