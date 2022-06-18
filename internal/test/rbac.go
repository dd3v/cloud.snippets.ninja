package test

import (
	"context"

	"github.com/dd3v/cloud.snippets.ninja/internal/entity"
)

type RBACMock struct {
	CanViewSnippetFn   func(ctx context.Context, snippet entity.Snippet) error
	CanUpdateSnippetFn func(ctx context.Context, snippet entity.Snippet) error
	CanDeleteSnippetFn func(ctx context.Context, snippet entity.Snippet) error
	GetUserIDFn        func(ctx context.Context) int
}

func (r RBACMock) CanViewSnippet(ctx context.Context, snippet entity.Snippet) error {
	return r.CanViewSnippetFn(ctx, snippet)
}

func (r RBACMock) CanUpdateSnippet(ctx context.Context, snippet entity.Snippet) error {
	return r.CanUpdateSnippetFn(ctx, snippet)
}

func (r RBACMock) CanDeleteSnippet(ctx context.Context, snippet entity.Snippet) error {
	return r.CanDeleteSnippetFn(ctx, snippet)
}

func (r RBACMock) GetUserID(ctx context.Context) int {
	identity := ctx.Value(entity.JWTCtxKey).(entity.Identity)
	return identity.GetID()
}
