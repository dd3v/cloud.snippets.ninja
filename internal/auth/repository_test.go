// +build integration

package auth

import (
	"context"
	"testing"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"github.com/dd3v/snippets.page.backend/pkg/dbcontext"
	"github.com/stretchr/testify/assert"
)

var db *dbcontext.DB
var r Repository
var table = "sessions"

func TestRepositoryMain(t *testing.T) {
	db = test.Database(t)
	test.TruncateTable(t, db, table)
	r = NewRepository(db)
}

func TestRepositoryCreateSession(t *testing.T) {
	session := entity.Session{
		UserID:       1,
		RefreshToken: "587e4ac6-8722-11ea-91bf-acde48001122",
		Exp:          time.Now().Add(time.Hour * 24),
		IP:           "127.0.0.1",
		UserAgent:    "Insomnia",
		CreatedAt:    time.Now(),
	}
	err := r.CreateSession(context.TODO(), session)
	assert.Nil(t, err)
}

func TestRepositoryFindSessionByRefreshToken(t *testing.T) {
	session, err := r.FindSessionByRefreshToken(context.TODO(), "587e4ac6-8722-11ea-91bf-acde48001122")
	assert.Equal(t, "587e4ac6-8722-11ea-91bf-acde48001122", session.RefreshToken)
	assert.Nil(t, err)
}

func TestDeleteSessionByRefreshToken(t *testing.T) {
	err := r.DeleteSessionByRefreshToken(context.TODO(), "587e4ac6-8722-11ea-91bf-acde48001122")
	assert.Nil(t, err)
	session, err := r.FindSessionByRefreshToken(context.TODO(), "587e4ac6-8722-11ea-91bf-acde48001122")
	assert.NotNil(t, err)
	assert.Empty(t, session.ID)
}
