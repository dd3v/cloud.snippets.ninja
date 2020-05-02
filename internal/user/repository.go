package user

import (
	"context"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/dbcontext"
)

//Repository - ...
type Repository interface {
	List(context context.Context, limit int, offset int) ([]entity.User, error)
	FindByID(context context.Context, id int) (entity.User, error)
	Create(context context.Context, user entity.User) (entity.User, error)
	Update(context context.Context, user entity.User) error
	Delete(context context.Context, id int) error
	Count(context context.Context) (int, error)
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

func (r repository) List(context context.Context, limit int, offset int) ([]entity.User, error) {
	var users []entity.User
	err := r.db.With(context).Select().Limit(int64(limit)).Offset(int64(offset)).OrderBy("id").All(&users)
	return users, err
}

func (r repository) FindByID(context context.Context, id int) (entity.User, error) {
	var user entity.User
	err := r.db.With(context).Select().Model(id, &user)
	return user, err
}

func (r repository) Create(context context.Context, user entity.User) (entity.User, error) {
	err := r.db.With(context).Model(&user).Insert()
	return user, err
}

func (r repository) Update(context context.Context, user entity.User) error {
	return r.db.With(context).Model(&user).Update()
}

func (r repository) Delete(context context.Context, id int) error {
	user, err := r.FindByID(context, id)
	if err != nil {
		return err
	}
	return r.db.With(context).Model(&user).Delete()
}

func (r repository) Count(context context.Context) (int, error) {
	var count int
	err := r.db.With(context).Select("COUNT(*)").From("users").Row(&count)
	return count, err
}
