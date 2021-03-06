package mock

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dd3v/snippets.page.backend/internal/entity"
)

var errorRepository = errors.New("error repository")

type UserMemoryRepository struct {
	items []entity.User
}

func NewRepository(items []entity.User) UserMemoryRepository {
	r := UserMemoryRepository{}
	r.items = items
	return r
}

func (r UserMemoryRepository) List(ctx context.Context, limit int, offset int) ([]entity.User, error) {
	return r.items, nil
}

func (r UserMemoryRepository) GetByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	for i, item := range r.items {
		if item.ID == id {
			return r.items[i], nil
		}
	}
	return user, sql.ErrNoRows
}

func (r UserMemoryRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	if user.Login == "error" {
		return entity.User{}, errorRepository
	}
	r.items = append(r.items, user)
	return user, nil
}

func (r UserMemoryRepository) Update(ctx context.Context, user entity.User) error {
	if user.Login == "error" {
		return errorRepository
	}
	for i, item := range r.items {
		if item.ID == user.ID {
			r.items[i] = user
			break
		}
	}
	return nil
}

func (r UserMemoryRepository) Delete(ctx context.Context, id int) error {
	for i, item := range r.items {
		if item.ID == id {
			r.items[i] = r.items[len(r.items)-1]
			r.items = r.items[:len(r.items)-1]
			break
		}
	}
	return nil
}

func (r UserMemoryRepository) Count(ctx context.Context) (int, error) {
	return len(r.items), nil
}
