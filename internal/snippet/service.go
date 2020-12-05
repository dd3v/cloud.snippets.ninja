package snippet

import (
	"context"

	"github.com/dd3v/snippets.page.backend/internal/entity"
)

type SearchRequest struct {
	Keyword  string
	Favorite bool
}

type Service interface {
	GetByUserID(context context.Context, userID int, request QuerySnippetsRequest) ([]entity.Snippet, error)
	List(context context.Context, public bool, limit int, offset int) ([]entity.Snippet, error)
	FindByID(context context.Context, id int) (entity.Snippet, error)
	//Create(context context.Context, request CreateRequest) (entity.User, error)
	//Update(context context.Context, id int, request UpdateRequest) (entity.User, error)
	Delete(context context.Context, id int) error
	Count(context context.Context) (int, error)
}

type service struct {
	repository Repository
}

//NewService - ...
func NewService(repository Repository) Service {
	return service{
		repository: repository,
	}
}

func (s service) GetByUserID(context context.Context, userID int, request QuerySnippetsRequest) ([]entity.Snippet, error) {

	conditions := map[string]interface{}{}

	conditions["favorite"] = request.Favorite
	conditions["public"] = request.Public
	if request.Keywords != "" {
		conditions["keywords"] = request.Keywords
	}

	snippets, err := s.repository.GetByUserID(context, userID, request.Limit, request.Offset, conditions)

	return snippets, err
}

func (s service) List(context context.Context, public bool, limit int, offset int) ([]entity.Snippet, error) {
	return s.repository.List(context, limit, offset)
}

func (s service) FindByID(context context.Context, id int) (entity.Snippet, error) {
	return s.repository.FindByID(context, id)
}

func (s service) Delete(context context.Context, id int) error {
	return s.repository.Delete(context, id)
}

func (s service) Count(context context.Context) (int, error) {
	return s.repository.Count(context)
}
