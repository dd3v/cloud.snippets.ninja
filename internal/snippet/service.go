package snippet

import (
	"context"

	"github.com/dd3v/snippets.ninja/internal/entity"
	"github.com/dd3v/snippets.ninja/pkg/query"
)

type Service interface {
	GetByID(ctx context.Context, id int) (entity.Snippet, error)
	QueryByUserID(context.Context, int, map[string]string, query.Sort, query.Pagination) ([]entity.Snippet, error)
	Create(context context.Context, snippet entity.Snippet) (entity.Snippet, error)
	Update(context context.Context, snippet entity.Snippet) (entity.Snippet, error)
	Delete(context context.Context, id int) error
	GetTags(ctx context.Context, userID int) (entity.Tags, error)
	CountByUserID(context context.Context, userID int, filter map[string]string) (int, error)
}

type Repository interface {
	QueryByUserID(ctx context.Context, userID int, filter map[string]string, sort query.Sort, pagination query.Pagination) ([]entity.Snippet, error)
	GetByID(ctx context.Context, id int) (entity.Snippet, error)
	Create(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error)
	Update(ctx context.Context, snippet entity.Snippet) error
	Delete(ctx context.Context, snippet entity.Snippet) error
	GetTags(ctx context.Context, userID int) (entity.Tags, error)
	CountByUserID(ctx context.Context, userID int, filter map[string]string) (int, error)
}

type RBAC interface {
	CanViewSnippet(ctx context.Context, snippet entity.Snippet) error
	CanDeleteSnippet(ctx context.Context, snippet entity.Snippet) error
	CanUpdateSnippet(ctx context.Context, snippet entity.Snippet) error
}

type service struct {
	repository Repository
	rbac       RBAC
}

//NewService - ...
func NewService(repository Repository, rbac RBAC) Service {
	return service{
		repository: repository,
		rbac:       rbac,
	}
}

func (s service) GetByID(ctx context.Context, id int) (entity.Snippet, error) {
	snippet, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return entity.Snippet{}, err
	}
	if err := s.rbac.CanViewSnippet(ctx, snippet); err != nil {
		return entity.Snippet{}, err
	}
	return snippet, err
}

func (s service) QueryByUserID(ctx context.Context, userID int, filter map[string]string, sort query.Sort, pagination query.Pagination) ([]entity.Snippet, error) {
	snippets, err := s.repository.QueryByUserID(ctx, userID, filter, sort, pagination)
	return snippets, err
}

func (s service) Create(context context.Context, snippet entity.Snippet) (entity.Snippet, error) {
	return s.repository.Create(context, snippet)
}

func (s service) Update(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error) {
	_, err := s.repository.GetByID(ctx, snippet.ID)
	if err != nil {
		return entity.Snippet{}, err
	}
	if err := s.rbac.CanUpdateSnippet(ctx, snippet); err != nil {
		return entity.Snippet{}, err
	}

	err = s.repository.Update(ctx, snippet)
	if err != nil {
		return entity.Snippet{}, err
	}
	return snippet, err
}

func (s service) Delete(ctx context.Context, id int) error {
	snippet, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if err := s.rbac.CanDeleteSnippet(ctx, snippet); err != nil {
		return err
	}
	return s.repository.Delete(ctx, snippet)
}

func (s service) CountByUserID(ctx context.Context, userID int, filter map[string]string) (int, error) {
	return s.repository.CountByUserID(ctx, userID, filter)
}

func (s service) GetTags(ctx context.Context, userID int) (entity.Tags, error) {


	return s.repository.GetTags(ctx, 1)
}
