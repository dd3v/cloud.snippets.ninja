package auth

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type resource struct {
	service Service
}

//NewHTTPHandler - ...
func NewHTTPHandler(router *routing.RouteGroup, jwtAuthMiddleware routing.Handler, service Service) {
	r := resource{
		service: service,
	}
	router.Post("/auth/login", r.login)
	router.Use(jwtAuthMiddleware)
	router.Post("/auth/refresh", r.refresh)
	router.Post("/auth/logout", r.logout)
}

func (r resource) login(c *routing.Context) error {
	var request LoginRequest
	if err := c.Read(&request); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return err
	}
	token, err := r.service.Login(c.Request.Context(), request)
	if err != nil {
		return err
	}
	return c.Write(token)
}

func (r resource) refresh(c *routing.Context) error {
	var request RefreshRequest
	if err := c.Read(&request); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return err
	}
	token, err := r.service.Refresh(c.Request.Context(), request.RefreshToken)
	if err != nil {
		return err
	}
	return c.Write(token)
}

func (r resource) logout(c *routing.Context) error {
	var request LogoutRequest
	if err := c.Read(&request); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return err
	}
	err := r.service.Logout(c.Request.Context(), request.RefreshToken)
	return err
}
