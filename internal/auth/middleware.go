package auth

import (
	"context"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dgrijalva/jwt-go"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/auth"
)

type Identity interface {
	GetID() int
	GetLogin() string
}

// GetJWTMiddleware returns a JWT-based authentication middleware. Ozzo routing
func GetJWTMiddleware(verificationKey string) routing.Handler {
	return auth.JWT(verificationKey, auth.JWTOptions{TokenHandler: handleToken})
}

func handleToken(c *routing.Context, token *jwt.Token) error {
	ctx := WithUser(
		c.Request.Context(),
		int(token.Claims.(jwt.MapClaims)["id"].(float64)),
		token.Claims.(jwt.MapClaims)["login"].(string),
	)
	c.Request = c.Request.WithContext(ctx)
	return nil
}

// WithUser returns a context that contains the user identity from the given JWT.
func WithUser(ctx context.Context, id int, login string) context.Context {
	return context.WithValue(ctx, entity.JWTCtxKey, entity.Identity{ID: id, Login: login})
}

// CurrentUser returns the user identity from the given context.
// Nil is returned if no user identity is found in the context.
func CurrentUser(ctx context.Context) Identity {
	if user, ok := ctx.Value(entity.JWTCtxKey).(entity.Identity); ok {
		return user
	}
	return nil
}
