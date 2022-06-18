package auth

import (
	"context"
	"fmt"

	"github.com/dd3v/cloud.snippets.ninja/internal/entity"
	"github.com/dd3v/cloud.snippets.ninja/pkg/dbcontext"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type repository struct {
	db *dbcontext.DB
}

//NewRepository - ...
func NewRepository(db *dbcontext.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) GetUserByLoginOrEmail(ctx context.Context, value string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().From("users").Where(dbx.HashExp{"login": value}).OrWhere(dbx.HashExp{"email": value}).One(&user)
	return user, err
}

func (r repository) CreateSession(ctx context.Context, session entity.Session) error {
	return r.db.With(ctx).Model(&session).Insert()
}

func (r repository) GetSessionByRefreshToken(ctx context.Context, refreshToken string) (entity.Session, error) {
	var session entity.Session
	err := r.db.With(ctx).Select().Where(&dbx.HashExp{"refresh_token": refreshToken}).One(&session)
	return session, err
}

func (r repository) DeleteSessionByUserIDAndUserAgent(ctx context.Context, userID int, userAgent string) error {
	result, err := r.db.With(ctx).Delete("sessions", dbx.HashExp{"user_id": userID, "user_agent": userAgent}).Execute()
	fmt.Println(result)
	return err
}

func (r repository) DeleteSessionByRefreshToken(ctx context.Context, refreshToken string) error {
	session, err := r.GetSessionByRefreshToken(ctx, refreshToken)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&session).Delete()
}

func (r repository) DeleteSessionByUserID(ctx context.Context, userID int) (int64, error) {
	result, err := r.db.With(ctx).Delete("sessions", dbx.HashExp{"user_id": userID}).Execute()
	if err != nil {
		return 0, err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRows, nil

}
