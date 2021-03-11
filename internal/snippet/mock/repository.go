package mock

import (
	"errors"

	"github.com/dd3v/snippets.page.backend/internal/entity"
)

var errorRepository = errors.New("error repository")

type SnippetMemoryRepository struct {
	snippets []entity.Snippet
}
