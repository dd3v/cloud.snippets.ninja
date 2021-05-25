package mock

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"github.com/dd3v/snippets.page.backend/pkg/query"
)

var ErrorRepository = errors.New("error repository")

type SnippetMemoryRepository struct {
	snippets []entity.Snippet
}

type Repository interface {
	QueryByUserID(context.Context, int, map[string]string, query.Sort, query.Pagination) ([]entity.Snippet, error)
	GetByID(ctx context.Context, id int) (entity.Snippet, error)
	Create(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error)
	Update(ctx context.Context, snippet entity.Snippet) error
	Delete(ctx context.Context, snippet entity.Snippet) error
	CountByUserID(ctx context.Context, userID int, filter map[string]string) (int, error)
}

func NewRepository() SnippetMemoryRepository {
	r := SnippetMemoryRepository{}
	r.snippets = []entity.Snippet{
		{
			ID:                  1,
			UserID:              1,
			Favorite:            true,
			AccessLevel:         0,
			Title:               "PHP hello world",
			Content:             "<?php echo 'Hello world'; ?>",
			Language:            "php",
			CustomEditorOptions: entity.CustomEditorOptions{},
			CreatedAt:           test.Time(2020),
			UpdatedAt:           test.Time(2021),
		},
		{
			ID:                  2,
			UserID:              1,
			Favorite:            false,
			AccessLevel:         1,
			Title:               "Snippet 2",
			Content:             "test 2",
			Language:            "go",
			CustomEditorOptions: entity.CustomEditorOptions{Theme: "default"},
			CreatedAt:           test.Time(2020),
			UpdatedAt:           test.Time(2021),
		},
		{
			ID:                  3,
			UserID:              1,
			Favorite:            false,
			AccessLevel:         1,
			Title:               "Snippet 3",
			Content:             "test 3",
			Language:            "javascript",
			CustomEditorOptions: entity.CustomEditorOptions{Theme: "default"},
			CreatedAt:           test.Time(2020),
			UpdatedAt:           test.Time(2021),
		},
		{
			ID:                  4,
			UserID:              2,
			Favorite:            false,
			AccessLevel:         1,
			Title:               "Snippet 4",
			Content:             "test 4",
			Language:            "javascript",
			CustomEditorOptions: entity.CustomEditorOptions{Theme: "default"},
			CreatedAt:           test.Time(2020),
			UpdatedAt:           test.Time(2021),
		},
	}

	return r
}

func (r SnippetMemoryRepository) QueryByUserID(content context.Context, userID int, filter map[string]string, sort query.Sort, pagination query.Pagination) ([]entity.Snippet, error) {
	var snippets []entity.Snippet

	if userID == 0 {
		return snippets, ErrorRepository
	}

	return r.snippets, nil
}

func (r SnippetMemoryRepository) GetByID(ctx context.Context, id int) (entity.Snippet, error) {
	var snippet entity.Snippet
	for i, item := range r.snippets {
		if item.ID == id {
			return r.snippets[i], nil
		}
	}
	return snippet, sql.ErrNoRows
}

func (r SnippetMemoryRepository) Create(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error) {
	if snippet.Title == "error" {
		return entity.Snippet{}, ErrorRepository
	}
	r.snippets = append(r.snippets, snippet)
	return snippet, nil
}

func (r SnippetMemoryRepository) Update(ctx context.Context, snippet entity.Snippet) error {
	if snippet.Title == "error" {
		return ErrorRepository
	}
	for i, item := range r.snippets {
		if item.ID == snippet.ID {
			r.snippets[i] = snippet
			break
		}
	}
	return nil
}

func (r SnippetMemoryRepository) Delete(ctx context.Context, snippet entity.Snippet) error {
	for i, item := range r.snippets {
		if item.ID == snippet.ID {
			r.snippets[i] = r.snippets[len(r.snippets)-1]
			r.snippets = r.snippets[:len(r.snippets)-1]
			break
		}
	}
	return nil
}

func (r SnippetMemoryRepository) CountByUserID(ctx context.Context, userID int, filter map[string]string) (int, error) {

	if userID == 0 {
		return 0, ErrorRepository
	}

	count := 0
	for _, item := range r.snippets {
		if item.UserID == userID {
			count++
		}
	}
	return count, nil
}
