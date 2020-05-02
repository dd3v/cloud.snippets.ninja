package auth

import (
	"context"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/dbcontext"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

//Repository - ...
type Repository interface {
	FindUser(context context.Context, login string) (entity.User, error)
	CreateSession(context context.Context, session entity.Session) error
	FindSessionByRefreshToken(context context.Context, refreshToken string) (entity.Session, error)
	DeleteSessionByRefreshToken(context context.Context, refreshToken string) error
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

func (r repository) FindUser(context context.Context, login string) (entity.User, error) {
	var user entity.User
	err := r.db.With(context).Select().From("users").Where(dbx.HashExp{"login": login}).OrWhere(dbx.HashExp{"email": login}).One(&user)
	return user, err
}

func (r repository) CreateSession(context context.Context, session entity.Session) error {
	return r.db.With(context).Model(&session).Insert()
}

func (r repository) FindSessionByRefreshToken(context context.Context, refreshToken string) (entity.Session, error) {
	var session entity.Session
	err := r.db.With(context).Select().Where(&dbx.HashExp{"refresh_token": refreshToken}).One(&session)
	return session, err
}

func (r repository) DeleteSessionByRefreshToken(context context.Context, refreshToken string) error {
	session, err := r.FindSessionByRefreshToken(context, refreshToken)
	if err != nil {
		return err
	}
	return r.db.With(context).Model(&session).Delete()
}
