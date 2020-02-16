package user

import (
	"context"
	"fmt"

	routing "github.com/go-ozzo/ozzo-routing"
)

type resource struct {
	service Service
}

func NewHTTPHandler(router *routing.RouteGroup, service Service) {
	res := resource{
		service: service,
	}
	router.Get("/me", res.me)
	router.Get("/users/<id>", res.view)
	router.Post("/users", res.create)
}

func (r resource) me(c *routing.Context) error {
	return nil
}

func (r resource) view(c *routing.Context) error {
	return nil
}

func (r resource) create(c *routing.Context) error {
	var request CreateRequest
	err := c.Read(&request)
	if err != nil {

	}
	err = request.Validate()
	if err != nil {
		return err
	}

	q, _ := r.service.CreateUser(context.TODO(), request)

	fmt.Println(q)
	fmt.Printf("%v", request)

	return nil
}
