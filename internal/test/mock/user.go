package mock

import (
	"context"
	"errors"

	"github.com/dd3v/snippets.page.backend/internal/entity"
)

var errorRepository = errors.New("error repository")

type MockUserRepository struct {
	items []entity.User
}

func (r *MockUserRepository) FindByID(context context.Context, id interface{}) (entity.User, error) {
	var user entity.User
	if id == "error" {
		return user, errorRepository
	}
	for i, item := range r.items {
		if item.ID == id {
			return r.items[i], nil
		}
	}
	return user, nil
}

func (r *MockUserRepository) Create(context context.Context, user entity.User) error {
	if user.Login == "error" {
		return errorRepository
	}
	r.items = append(r.items, user)
	return nil
}

func (r *MockUserRepository) Update(context context.Context, user entity.User) error {
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

func (r *MockUserRepository) Delete(context context.Context, id interface{}) error {
	for i, item := range r.items {
		if item.ID == id {
			r.items[i] = r.items[len(r.items)-1]
			r.items = r.items[:len(r.items)-1]
			break
		}
	}
	return nil
}

func (r *MockUserRepository) Count(context context.Context) (int, error) {
	return len(r.items), nil
}
