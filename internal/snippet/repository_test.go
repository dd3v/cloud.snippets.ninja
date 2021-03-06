// +build integration

package snippet

import (
	"context"
	"testing"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"github.com/dd3v/snippets.page.backend/pkg/dbcontext"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

var db *dbcontext.DB
var r Repository
var table = "snippets"

func TestSnippetRepositoryMain(t *testing.T) {
	db = test.Database(t)
	test.TruncateTable(t, db, table)
	r = NewRepository(db)
}

func TestCreate(t *testing.T) {
	snippet := entity.Snippet{
		ID:            1,
		UserID:        1,
		Favorite:      false,
		Access:        0,
		Title:         "Test Snippet",
		Content:       null.NewString("Hello world", true),
		Language:      "txt",
		EditorOptions: entity.EditorOptions{},
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	response, err := r.Create(context.TODO(), snippet)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, snippet, response)
}

func TestCount(t *testing.T) {

}
