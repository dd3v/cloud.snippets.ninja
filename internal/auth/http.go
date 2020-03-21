package auth

import (
	"context"
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type resource struct {
	service Service
}

func NewHTTPHandler(router *routing.RouteGroup, service Service) {
	r := resource{
		service: service,
	}
	router.Get("/me", r.me)
	router.Post("/login", r.login)
	router.Post("/refresh", r.refresh)
}

func (r resource) me(c *routing.Context) error {
	return c.Write("/me")
}

func (r resource) login(c *routing.Context) error {
	var request AuthRequest
	if err := c.Read(&request); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return c.WriteWithStatus(err, http.StatusBadRequest)
	}
	token, err := r.service.Login(context.TODO(), request)
	if err != nil {
		return c.WriteWithStatus(err.Error(), http.StatusBadRequest)
	}
	return c.Write(token)
}

func (r resource) refresh(c *routing.Context) error {

	var request RefreshRequest

	if err := c.Read(&request); err != nil {
		return c.WriteWithStatus(err, http.StatusBadRequest)
	}
	if err := request.Validate(); err != nil {
		return c.WriteWithStatus(err, http.StatusBadRequest)
	}
	token, err := r.service.Refresh(context.TODO(), request.RefreshToken)
	if err != nil {
		return c.WriteWithStatus(err, http.StatusForbidden)
	}

	return c.Write(token)
}
