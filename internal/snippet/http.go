package snippet

import (
	"fmt"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/errors"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type resource struct {
	service Service
}

//NewHTTPHandler - ...
func NewHTTPHandler(router *routing.RouteGroup, jwtAuthHandler routing.Handler, service Service) {
	r := resource{
		service: service,
	}

	router.Use(jwtAuthHandler)
	router.Get("/me/snippets", r.snippets)
	router.Post("/me/snippets", r.create)
	router.Put("/me/snippets/<id>", r.update)
	router.Delete("/me/snippets/<id>", r.delete)

}

func (r resource) snippets(c *routing.Context) error {
	identity := c.Request.Context().Value(entity.JWTContextKey).(entity.Identity)

	var request QuerySnippetsRequest

	if err := c.Read(&request); err != nil {
		return errors.BadRequest("")
	}

	fmt.Println("======BINDED REQUEST======")
	fmt.Println(request.Favorite.Value())
	fmt.Println(request.Public)
	fmt.Println("======BINDED REQUEST======")

	if err := request.Validate(); err != nil {
		fmt.Println(err)
		return err
	}

	snippets, err := r.service.GetByUserID(c.Request.Context(), identity.GetID(), request)
	if err != nil {
		return err
	}

	return c.Write(snippets)
}

func (r resource) create(c *routing.Context) error {
	return c.Write("create")
}

func (r resource) update(c *routing.Context) error {
	return c.Write("update")
}

func (r resource) delete(c *routing.Context) error {
	return c.Write("delete")
}
