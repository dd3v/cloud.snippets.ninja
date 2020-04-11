package user

import (
	"context"
	"strconv"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type resource struct {
	service Service
}

//NewHTTPHandler -
func NewHTTPHandler(router *routing.RouteGroup, service Service) {
	r := resource{
		service: service,
	}
	router.Get("/users", r.list)
	router.Get("/users/<id>", r.get)
	router.Post("/users", r.create)
	router.Put("/users/<id>", r.update)
	router.Delete("/users/<id>", r.delete)
}

func (r resource) list(c *routing.Context) error {

	cnt, err := r.service.Count(context.TODO())
	if err != nil {
		return err
	}

	return c.Write(cnt)
}

func (r resource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user, err := r.service.FindByID(context.TODO(), id)
	if err != nil {
		return err
	}
	return c.Write(user)
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
	user, err := r.service.Create(context.TODO(), request)
	if err != nil {
		return err
	}
	return c.Write(user)
}

func (r resource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	var request UpdateRequest
	if err := c.Read(&request); err != nil {
		return err
	}
	if err = request.Validate(); err != nil {
		return err
	}
	user, err := r.service.Update(context.TODO(), id, request)
	if err != nil {
		return err
	}
	return c.Write(user)
}

func (r resource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := r.service.Delete(context.TODO(), id); err != nil {
		return err
	}

	return c.Write(nil)
}
