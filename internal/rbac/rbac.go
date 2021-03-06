package rbac

import (
	"context"
	"errors"

	"github.com/dd3v/snippets.page.backend/internal/entity"
)

type RBAC struct{}

func New() RBAC {
	return RBAC{}
}

func (r RBAC) CanViewSnippet(ctx context.Context, snippet entity.Snippet) error {
	if snippet.UserID == r.getUserID(ctx) {
		return nil
	}
	return errors.New("Forbidden")
}

func (r RBAC) CanUpdateSnippet(ctx context.Context, snippet entity.Snippet) error {
	if snippet.UserID == r.getUserID(ctx) {
		return nil
	}
	return errors.New("Forbidden")
}

func (r RBAC) CanDeleteSnippet(ctx context.Context, snippet entity.Snippet) error {
	if snippet.UserID == r.getUserID(ctx) {
		return nil
	}
	return errors.New("Forbidden")
}

func (r RBAC) getUserID(ctx context.Context) int {
	identity := ctx.Value(entity.JWTContextKey).(entity.Identity)
	return identity.GetID()
}
