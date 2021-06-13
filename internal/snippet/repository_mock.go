package snippet

import (
	"context"
	"errors"

	"github.com/dd3v/snippets.ninja/internal/entity"
	"github.com/dd3v/snippets.ninja/pkg/query"
)

var repositoryMockErr = errors.New("error repository")

type RepositoryMock struct {
	QueryByUserIDFn func(ctx context.Context, userID int, filter map[string]string, sort query.Sort, pagination query.Pagination) ([]entity.Snippet, error)
	GetByIDFn       func(ctx context.Context, id int) (entity.Snippet, error)
	CreateFn        func(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error)
	UpdateFn        func(ctx context.Context, snippet entity.Snippet) error
	DeleteFn        func(ctx context.Context, snippet entity.Snippet) error
	CountByUserIDFn func(ctx context.Context, userID int, filter map[string]string) (int, error)
}

func (r RepositoryMock) QueryByUserID(ctx context.Context, userID int, filter map[string]string, sort query.Sort, pagination query.Pagination) ([]entity.Snippet, error) {
	return r.QueryByUserIDFn(ctx, userID, filter, sort, pagination)
}

func (r RepositoryMock) GetByID(ctx context.Context, id int) (entity.Snippet, error) {
	return r.GetByIDFn(ctx, id)
}

func (r RepositoryMock) Create(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error) {
	return r.CreateFn(ctx, snippet)
}

func (r RepositoryMock) Update(ctx context.Context, snippet entity.Snippet) error {
	return r.UpdateFn(ctx, snippet)
}

func (r RepositoryMock) Delete(ctx context.Context, snippet entity.Snippet) error {
	return r.DeleteFn(ctx, snippet)
}

func (r RepositoryMock) CountByUserID(ctx context.Context, userID int, filter map[string]string) (int, error) {
	return r.CountByUserIDFn(ctx, userID, filter)
}
