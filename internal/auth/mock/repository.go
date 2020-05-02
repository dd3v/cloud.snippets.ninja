package mock

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dd3v/snippets.page.backend/internal/entity"
)

var errorRepository = errors.New("error repository")

//SessionMemoryRepository - ...
type SessionMemoryRepository struct {
	sessions []entity.Session
	users    []entity.User
}

//NewRepository - ...
func NewRepository(users []entity.User, sessions []entity.Session) SessionMemoryRepository {
	r := SessionMemoryRepository{}
	r.sessions = sessions
	r.users = users
	return r
}

//FindUser - ...
func (r SessionMemoryRepository) FindUser(context context.Context, login string) (entity.User, error) {
	var user entity.User
	for _, user := range r.users {
		if user.Login == login {
			return user, nil
		}
	}
	return user, sql.ErrNoRows
}

//CreateSession - ...
func (r SessionMemoryRepository) CreateSession(context context.Context, session entity.Session) error {
	r.sessions = append(r.sessions, session)
	return nil
}

//FindSessionByRefreshToken - ...
func (r SessionMemoryRepository) FindSessionByRefreshToken(context context.Context, refreshToken string) (entity.Session, error) {
	var session entity.Session
	for _, s := range r.sessions {
		if s.RefreshToken == refreshToken {
			return s, nil
		}
	}
	return session, sql.ErrNoRows
}

//DeleteSessionByRefreshToken - ...
func (r SessionMemoryRepository) DeleteSessionByRefreshToken(context context.Context, refreshToken string) error {
	for i, s := range r.sessions {
		if s.RefreshToken == refreshToken {
			r.sessions[i] = r.sessions[len(r.sessions)-1]
			r.sessions = r.sessions[:len(r.sessions)-1]
			return nil
		}
	}
	return sql.ErrNoRows
}
