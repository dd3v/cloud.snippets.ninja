// +build integration

package snippet

import (
	"context"
	"testing"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"github.com/dd3v/snippets.page.backend/pkg/dbcontext"
	"github.com/dd3v/snippets.page.backend/pkg/query"
	"github.com/stretchr/testify/assert"
)

var db *dbcontext.DB
var r Repository
var table = "snippets"

func TestRepository_Main(t *testing.T) {
	t.Logf("\033[35m" + "Testing Snippet Repository" + "\033[0m")
	db = test.Database(t)
	test.TruncateTable(t, db, table)
	r = NewRepository(db)
}

func TestRepository_Create(t *testing.T) {
	snippet := entity.Snippet{
		ID:                  1,
		UserID:              1,
		Favorite:            false,
		AccessLevel:         0,
		Title:               "Test Snippet",
		Content:             "",
		Language:            "txt",
		CustomEditorOptions: entity.CustomEditorOptions{},
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
	response, err := r.Create(context.Background(), snippet)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, snippet, response)
}

func TestRepository_GetByID(t *testing.T) {
	snippet, err := r.GetByID(context.Background(), 1)
	assert.NotEmpty(t, snippet)
	assert.Nil(t, err)
}

func TestRepository_Update(t *testing.T) {
	snippet := entity.Snippet{
		ID:          1,
		UserID:      1,
		Favorite:    true,
		AccessLevel: 1,
		Title:       "New title",
		Content:     "New text",
		Language:    "php",
		CustomEditorOptions: entity.CustomEditorOptions{
			Theme: "default",
		},
		UpdatedAt: time.Now(),
	}
	err := r.Update(context.Background(), snippet)
	assert.Nil(t, err)
	updated, err := r.GetByID(context.Background(), 1)
	assert.Nil(t, err)
	assert.Equal(t, snippet.ID, updated.ID)
	assert.Equal(t, snippet.UserID, updated.UserID)
	assert.Equal(t, snippet.Favorite, updated.Favorite)
	assert.Equal(t, snippet.AccessLevel, updated.AccessLevel)
	assert.Equal(t, snippet.Title, updated.Title)
	assert.Equal(t, snippet.Content, updated.Content)
	assert.Equal(t, snippet.Language, updated.Language)
	assert.Equal(t, snippet.CustomEditorOptions, updated.CustomEditorOptions)
	//assert.Equal(t, snippet.UpdatedAt, updated.UpdatedAt)

}

func TestRepository_Count(t *testing.T) {
	count, err := r.Count(context.Background(), 1, map[string]string{})
	assert.Nil(t, err)
	assert.Equal(t, count, 1)
}

func TestRepository_Delete(t *testing.T) {
	snippet, err := r.GetByID(context.Background(), 1)
	assert.Nil(t, err)
	err = r.Delete(context.Background(), snippet)
	assert.Nil(t, err)
	count, err := r.Count(context.Background(), 1, map[string]string{})
	assert.Nil(t, err)
	assert.Equal(t, count, 0)
}

func TestRepository_List(t *testing.T) {
	snippets := []entity.Snippet{
		entity.Snippet{
			UserID:              1,
			Favorite:            true,
			AccessLevel:         0,
			Title:               "Binary Search",
			Content:             "",
			Language:            "php",
			CustomEditorOptions: entity.CustomEditorOptions{},
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		},
		entity.Snippet{
			UserID:              1,
			Favorite:            true,
			AccessLevel:         0,
			Title:               "Linear search",
			Content:             "",
			Language:            "js",
			CustomEditorOptions: entity.CustomEditorOptions{},
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		},
		entity.Snippet{
			UserID:              1,
			Favorite:            false,
			AccessLevel:         1,
			Title:               "Bubble sort",
			Content:             "",
			Language:            "go",
			CustomEditorOptions: entity.CustomEditorOptions{},
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		},
	}

	for _, snippet := range snippets {
		snippet, err := r.Create(context.Background(), snippet)
		if err != nil {
			t.Fail()
		}
		t.Logf("Saved test snippet. ID: %d", snippet.ID)
	}

	filter := map[string]string{
		"favorite": "1",
		"language": "php",
	}
	sort := query.NewSort("id", "asc")
	pagination := query.NewPagination(1, 10)
	snippets, err := r.List(context.Background(), 1, filter, sort, pagination)
	assert.Nil(t, err)
	assert.True(t, len(snippets) > 0)
	for _, snippet := range snippets {
		assert.True(t, snippet.Favorite)
		assert.Equal(t, snippet.Language, "php")
	}
}
