package user

import (
	"context"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/dbcontext"
)

//Repository - ...
type Repository interface {
	List(ctx context.Context, limit int, offset int) ([]entity.User, error)
	GetByID(ctx context.Context, id int) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int, error)
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

func (r repository) List(ctx context.Context, limit int, offset int) ([]entity.User, error) {
	var users []entity.User
	err := r.db.With(ctx).Select().Limit(int64(limit)).Offset(int64(offset)).OrderBy("id").All(&users)
	return users, err
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

func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("users").Row(&count)
	return count, err
}
