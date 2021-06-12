package user

import (
	"net/http"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/errors"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type resource struct {
	service Service
}

//NewHTTPHandler -
func NewHTTPHandler(router *routing.RouteGroup, jwtAuthMiddleware routing.Handler, service Service) {
	r := resource{
		service: service,
	}
	router.Post("/users", r.create)
	router.Use(jwtAuthMiddleware)
	router.Get("/users/me", r.me)
}

func (r resource) create(c *routing.Context) error {
	var request createRequest
	if err := c.Read(&request); err != nil {
		return err
	}
	err := request.Validate()
	if err != nil {
		return err
	}

	if exists, err := r.service.Exists(c.Request.Context(), "email", request.Email); err != nil {
		return errors.InternalServerError(err.Error())
	} else if exists {
		return errors.BadRequest("email should be unique")
	}

	if exists, err := r.service.Exists(c.Request.Context(), "login", request.Login); err != nil {
		return errors.InternalServerError(err.Error())
	} else if exists {
		return errors.BadRequest("login should be unique")
	}

	user := entity.User{
		Password:  request.Password,
		Login:     request.Login,
		Email:     request.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err = r.service.Create(c.Request.Context(), user)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(user, http.StatusCreated)
}

func (r resource) me(c *routing.Context) error {
	identity := c.Request.Context().Value(entity.JWTCtxKey).(entity.Identity)
	me, err := r.service.GetByID(c.Request.Context(), identity.GetID())
	if err != nil {
		return err
	}
	return c.Write(me)
}
