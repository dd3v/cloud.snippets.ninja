package rbac

import (
	"context"
	"errors"
	"github.com/dd3v/snippets.page.backend/internal/entity"
)

type RBAC struct{}

var AccessError = errors.New("access denied")

func New() RBAC {
	return RBAC{}
}

func (r RBAC) CanViewSnippet(ctx context.Context, snippet entity.Snippet) error {
	return r.isOwner(r.GetUserID(ctx), snippet.GetOwnerID(), snippet.IsPublic())
}

func (r RBAC) CanUpdateSnippet(ctx context.Context, snippet entity.Snippet) error {
	return r.isOwner(r.GetUserID(ctx), snippet.GetOwnerID(), snippet.IsPublic())

}

func (r RBAC) CanDeleteSnippet(ctx context.Context, snippet entity.Snippet) error {
	return r.isOwner(r.GetUserID(ctx), snippet.GetOwnerID(), snippet.IsPublic())
}

func (r RBAC) isOwner(userID int, ownerID int, public bool) error {
	if userID == ownerID || public == true {
		return nil
	}
	return AccessError
}

func (r RBAC) GetUserID(ctx context.Context) int {
	identity := ctx.Value(entity.JWTCtxKey).(entity.Identity)
	return identity.GetID()
}
