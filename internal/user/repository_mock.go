package user

import (
	"context"
	"errors"

	"github.com/dd3v/cloud.snippets.ninja/internal/entity"
)

var repositoryMockErr = errors.New("error repository")

type RepositoryMock struct {
	GetByIDFn func(ctx context.Context, id int) (entity.User, error)
	CreateFn  func(ctx context.Context, user entity.User) (entity.User, error)
	UpdateFn  func(ctx context.Context, user entity.User) error
	DeleteFn  func(ctx context.Context, id int) error
	ExistsFn  func(ctx context.Context, field string, value string) (bool, error)
}

func (r RepositoryMock) GetByID(ctx context.Context, id int) (entity.User, error) {
	return r.GetByIDFn(ctx, id)
}

func (r RepositoryMock) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return r.CreateFn(ctx, user)
}

func (r RepositoryMock) Update(ctx context.Context, user entity.User) error {
	return r.UpdateFn(ctx, user)
}

func (r RepositoryMock) Delete(ctx context.Context, id int) error {
	return r.DeleteFn(ctx, id)
}

func (r RepositoryMock) Exists(ctx context.Context, field string, value string) (bool, error) {
	return r.ExistsFn(ctx, field, value)
}
