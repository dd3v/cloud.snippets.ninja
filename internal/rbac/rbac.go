package rbac

import (
	"context"
	"errors"

	"github.com/dd3v/cloud.snippets.ninja/internal/entity"
)

type RBAC struct{}

var AccessError = errors.New("accesslog denied")

func New() RBAC {
	return RBAC{}
}

func (r RBAC) CanViewSnippet(ctx context.Context, snippet entity.Snippet) error {
	return r.isOwner(r.GetUserID(ctx), snippet.GetOwnerID())
}

func (r RBAC) CanUpdateSnippet(ctx context.Context, snippet entity.Snippet) error {
	return r.isOwner(r.GetUserID(ctx), snippet.GetOwnerID())

}

func (r RBAC) CanDeleteSnippet(ctx context.Context, snippet entity.Snippet) error {
	return r.isOwner(r.GetUserID(ctx), snippet.GetOwnerID())
}

func (r RBAC) isOwner(userID int, ownerID int) error {
	if userID == ownerID {
		return nil
	}
	return AccessError
}

func (r RBAC) GetUserID(ctx context.Context) int {
	identity := ctx.Value(entity.JWTCtxKey).(entity.Identity)
	return identity.GetID()
}
