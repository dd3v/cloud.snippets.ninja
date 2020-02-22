package mock

import (
	"context"
	"errors"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var errorRepository = errors.New("error repository")

type MemoryUserRepository struct {
	items []entity.User
}

func (r *MemoryUserRepository) Find(context context.Context, condition map[string]interface{}) ([]entity.User, error) {
	return r.items, nil
}

func (r *MemoryUserRepository) FindByID(context context.Context, id string) (entity.User, error) {
	var user entity.User
	if id == "error" {
		return user, errorRepository
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	for i, item := range r.items {
		if item.ID == objectID {
			return r.items[i], nil
		}
	}
	return user, nil
}

func (r *MemoryUserRepository) Create(context context.Context, user entity.User) error {
	if user.Login == "error" {
		return errorRepository
	}
	r.items = append(r.items, user)
	return nil
}

func (r *MemoryUserRepository) Update(context context.Context, user entity.User) error {
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

func (r *MemoryUserRepository) Delete(context context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	for i, item := range r.items {
		if item.ID == objectID {
			r.items[i] = r.items[len(r.items)-1]
			r.items = r.items[:len(r.items)-1]
			break
		}
	}
	return nil
}

func (r *MemoryUserRepository) Count(context context.Context) (int, error) {
	return len(r.items), nil
}
