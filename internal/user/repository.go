package user

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"

	"github.com/dd3v/cloud.snippets.ninja/internal/entity"
	"github.com/dd3v/cloud.snippets.ninja/pkg/dbcontext"
)

type repository struct {
	db *dbcontext.DB
}

func NewRepository(db *dbcontext.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) GetByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().Model(id, &user)
	return user, err
}

func (r repository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.db.With(ctx).Model(&user).Insert()
	return user, err
}

func (r repository) Update(ctx context.Context, user entity.User) error {
	return r.db.With(ctx).Model(&user).Update()
}

func (r repository) Delete(ctx context.Context, id int) error {
	user, err := r.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&user).Delete()
}

func (r repository) Exists(ctx context.Context, field string, value string) (bool, error) {
	var count int
	var exists bool
	err := r.db.With(ctx).Select("COUNT(*)").From("users").Where(dbx.HashExp{field: value}).Row(&count)
	if count == 0 {
		exists = false
	} else {
		exists = true
	}
	return exists, err
}
