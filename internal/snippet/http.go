package snippet

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/errors"
	"github.com/dd3v/snippets.page.backend/pkg/query"
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
	router.Get("/snippets/<id>", r.view)
	router.Post("/snippets", r.create)
	router.Put("/snippets/<id>", r.update)
	router.Delete("/snippets/<id>", r.delete)
	router.Get("/snippets", r.list)
}

func (r resource) view(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	snippet, err := r.service.GetByID(c.Request.Context(), id)
	if err != nil {
		return err
	}
	return c.Write(snippet)
}

func (r resource) create(c *routing.Context) error {
	identity := c.Request.Context().Value(entity.JWTCtxKey).(entity.Identity)
	request := snippet{}
	if err := c.Read(&request); err != nil {
		return err
	}
	if err := request.validate(); err != nil {
		return err
	}
	snippet := entity.Snippet{
		UserID:              identity.GetID(),
		Favorite:            request.Favorite.Value,
		AccessLevel:         request.AccessLevel,
		Title:               request.Title,
		Content:             request.Content,
		Language:            request.Language,
		CustomEditorOptions: request.CustomEditorOptions,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
	snippet, err := r.service.Create(c.Request.Context(), snippet)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(snippet, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	identity := c.Request.Context().Value(entity.JWTCtxKey).(entity.Identity)
	request := snippet{}
	if err := c.Read(&request); err != nil {
		return err
	}
	if err := request.validate(); err != nil {
		return err
	}
	snippet := entity.Snippet{
		ID:                  id,
		UserID:              identity.GetID(),
		Favorite:            request.Favorite.Value,
		AccessLevel:         request.AccessLevel,
		Title:               request.Title,
		Content:             request.Content,
		Language:            request.Language,
		CustomEditorOptions: request.CustomEditorOptions,
		UpdatedAt:           time.Now(),
	}
	response, err := r.service.Update(c.Request.Context(), snippet)
	if err != nil {
		return err
	}
	return c.Write(response)
}

func (r resource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = r.service.Delete(c.Request.Context(), id)
	if err != nil {
		return err
	}
	return c.Write("")
}

type listResponse struct {
	Items      []entity.Snippet `json:"items"`
	Page       int              `json:"page"`
	Limit      int              `json:"limit"`
	TotalItems int              `json:"total_items"`
	TotalPages int              `json:"total_pages"`
}

func (r resource) list(c *routing.Context) error {
	request := newList()
	identity := c.Request.Context().Value(entity.JWTCtxKey).(entity.Identity)
	if err := c.Read(&request); err != nil {
		return errors.BadRequest("")
	}
	if err := request.validate(); err != nil {
		return err
	}

	filter := request.filterConditions()
	total, err := r.service.CountByUserID(c.Request.Context(), identity.GetID(), filter)
	if err != nil {
		return err
	}
	pagination := query.NewPagination(request.Page, request.Limit)
	sort := query.NewSort(request.SortBy, request.OrderBy)
	snippets, err := r.service.QueryByUserID(c.Request.Context(), identity.GetID(), filter, sort, pagination)
	if err != nil {
		return err
	}
	return c.Write(listResponse{
		Items:      snippets,
		Page:       pagination.GetPage(),
		Limit:      pagination.GetLimit(),
		TotalItems: total,
		TotalPages: (total + pagination.GetLimit() - 1) / pagination.GetLimit(),
	})
}
