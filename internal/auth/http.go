package auth

import (
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
	router.Get("/refresh/<token>", r.refresh)
}

func (r resource) me(c *routing.Context) error {
	return c.Write("/me")
}

func (r resource) login(c *routing.Context) error {
	return c.Write("/logn")
}

func (r resource) refresh(c *routing.Context) error {
	return c.Write("/refresh")
}
