package snippet

import (
	"context"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
)

type SearchRequest struct {
	Keyword  string
	Favorite bool
}

type Service interface {
	FindByUserID(context context.Context, userID int, request OwnSnippetsRequest) ([]entity.Snippet, error)
	FindByID(context context.Context, id int) (entity.Snippet, error)
	Create(context context.Context, userID int, request CreateSnippetRequest) (entity.Snippet, error)
	Update(context context.Context, id int, userID int, request UpdateSnippetRequest) (entity.Snippet, error)
	Delete(context context.Context, id int, userID int) (entity.Snippet, error)
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

func (s service) FindByUserID(context context.Context, userID int, request OwnSnippetsRequest) ([]entity.Snippet, error) {
	conditions := map[string]interface{}{}
	if request.Access != -1 {
		conditions["access"] = request.Access
	}
	if request.Title != "" {
		conditions["title"] = request.Title
	}
	if request.Favorite != -1 {
		conditions["favorite"] = request.Favorite
	}
	snippets, err := s.repository.GetByUserID(context, userID, request.Limit, request.Offset, conditions)
	return snippets, err
}

func (s service) Create(context context.Context, userID int, request CreateSnippetRequest) (entity.Snippet, error) {
	return s.repository.Create(context, entity.Snippet{
		UserID:        userID,
		Favorite:      request.Favorite,
		Access:        request.Access,
		Title:         request.Title,
		Content:       request.Content,
		FileExtension: request.FileExtension,
		EditorOptions: request.EditorOptions,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	})
}

func (s service) Update(context context.Context, id int, userID int, request UpdateSnippetRequest) (entity.Snippet, error) {
	snippet, err := s.repository.GetByIDAndUserID(context, id, userID)
	if err != nil {
		return snippet,err
	}
	snippet.Favorite = request.Favorite
	snippet.Access = request.Access
	snippet.Title = request.Title
	snippet.Content.String = request.Content
	snippet.FileExtension = request.FileExtension
	snippet.EditorOptions = request.EditorOptions
	snippet.UpdatedAt = time.Now()

	return s.repository.Update(context, snippet)
}

func (s service) FindByID(context context.Context, id int) (entity.Snippet, error) {
	return s.repository.FindByID(context, id)
}

func (s service) Delete(context context.Context, id int, userID int) (entity.Snippet, error) {

	snippet, err := s.repository.GetByIDAndUserID(context, id, userID)
	if err != nil {
		return snippet,err
	}

	return s.repository.Delete(context, snippet)
}

func (s service) Count(context context.Context) (int, error) {
	return s.repository.Count(context)
}
