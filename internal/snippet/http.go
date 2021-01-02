package snippet

import (
	"fmt"
	"strconv"

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
	request := NewOwnSnippetsRequest()
	if err := c.Read(&request); err != nil {
		return errors.BadRequest("")
	}
	//fmt.Println(fmt.Sprintf("%#v", request))
	if err := request.Validate(); err != nil {
		fmt.Println(err)
		return err
	}
	snippets, err := r.service.FindByUserID(c.Request.Context(), identity.GetID(), request)
	if err != nil {
		return err
	}
	return c.Write(snippets)
}

func (r resource) create(c *routing.Context) error {

	identity := c.Request.Context().Value(entity.JWTContextKey).(entity.Identity)
	request := NewCreateSnippetRequest()
	if err := c.Read(&request); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return err
	}
	snippet, err := r.service.Create(c.Request.Context(), identity.GetID(), request)
	if err != nil {
		return err
	}
	return c.Write(snippet)
}

func (r resource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	identity := c.Request.Context().Value(entity.JWTContextKey).(entity.Identity)

	request := NewUpdateSnippetRequest()
	if err := c.Read(&request); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return err
	}

	snippet, err := r.service.Update(c.Request.Context(), id, identity.GetID(), request)
	if err != nil {
		return err
	}

	return c.Write(snippet)
}

func (r resource) delete(c *routing.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	identity := c.Request.Context().Value(entity.JWTContextKey).(entity.Identity)

	snippet, err := r.service.Delete(c.Request.Context(), id, identity.GetID());
	if err != nil {
		return err
	}

	return c.Write(snippet)
}
