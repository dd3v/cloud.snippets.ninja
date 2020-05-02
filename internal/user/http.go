package user

import (
	"net/http"
	"strconv"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/errors"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type resource struct {
	service Service
}

//NewHTTPHandler -
func NewHTTPHandler(router *routing.RouteGroup, jwtAuthHandler routing.Handler, service Service) {
	r := resource{
		service: service,
	}
	router.Post("/users", r.create)
	router.Use(jwtAuthHandler)
	router.Get("/users/me", r.me)
	router.Get("/users/<id>", r.get)
	router.Put("/users/me", r.update)
}

func (r resource) create(c *routing.Context) error {
	var request CreateRequest
	if err := c.Read(&request); err != nil {
		return err
	}
	err := request.Validate()
	if err != nil {
		return err
	}
	user, err := r.service.Create(c.Request.Context(), request)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(user, http.StatusCreated)
}

func (r resource) me(c *routing.Context) error {
	identity := c.Request.Context().Value(entity.JWTContextKey).(entity.Identity)
	me, err := r.service.FindByID(c.Request.Context(), identity.GetID())
	if err != nil {
		return err
	}
	return c.Write(me)
}

func (r resource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user, err := r.service.FindByID(c.Request.Context(), id)
	if err != nil {
		return err
	}
	return c.Write(user.GetPublicProfile())
}

func (r resource) update(c *routing.Context) error {
	identity := c.Request.Context().Value(entity.JWTContextKey).(entity.Identity)
	var request UpdateRequest
	if err := c.Read(&request); err != nil {
		return errors.BadRequest("")
	}
	if err := request.Validate(); err != nil {
		return err
	}
	user, err := r.service.Update(c.Request.Context(), identity.GetID(), request)
	if err != nil {
		return err
	}
	return c.Write(user)
}
