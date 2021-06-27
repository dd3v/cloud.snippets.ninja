package snippet

import (
	"strconv"

	"github.com/dd3v/snippets.ninja/internal/entity"
	"github.com/dd3v/snippets.ninja/pkg/datatype"
	validation "github.com/go-ozzo/ozzo-validation"
)

type list struct {
	Favorite    bool   `form:"favorite"`
	AccessLevel int    `form:"access_level"`
	Title       string `form:"title"`
	Tags        string `form:"tags"`
	SortBy      string `form:"sort_by"`
	OrderBy     string `form:"order_by"`
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
}

func (l list) filterConditions() map[string]string {
	conditions := make(map[string]string)
	if l.Favorite != false {
		conditions["favorite"] = "1"
	}
	if l.AccessLevel != -1 {
		conditions["access_level"] = strconv.Itoa(l.AccessLevel)
	}
	if l.Title != "" {
		conditions["title"] = l.Title
	}
	if l.Tags != "" {
		conditions["tags"] = l.Tags
	}

	return conditions
}

func newList() list {
	return list{Favorite: false, AccessLevel: -1, Title: "", Tags: "", SortBy: "id", OrderBy: "desc", Page: 1, Limit: 50}
}

func (l list) validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Favorite, validation.In(0, 1, false, true)),
		validation.Field(&l.AccessLevel, validation.In(-1, 0, 1)),
		validation.Field(&l.Title, validation.Length(0, 100)),
		validation.Field(&l.SortBy, validation.In("id")),
		validation.Field(&l.OrderBy, validation.In("asc", "desc")),
		validation.Field(&l.Page, validation.Min(1)),
		validation.Field(&l.Limit, validation.Min(1), validation.Max(100)),
	)
}

type snippet struct {
	Favorite            datatype.FlexibleBool      `json:"favorite"`
	AccessLevel         int                        `json:"access_level"`
	Title               string                     `json:"title"`
	Content             string                     `json:"content"`
	Tags                entity.Tags                `json:"tags"`
	Language            string                     `json:"language"`
	CustomEditorOptions entity.CustomEditorOptions `json:"custom_editor_options"`
}

func (r snippet) validate() error {
	err := validation.Errors{
		"title":                      validation.Validate(r.Title, validation.Required, validation.Length(1, 500)),
		"access_level":               validation.Validate(r.AccessLevel, validation.In(0, 1)),
		"editor_options.theme":       validation.Validate(r.CustomEditorOptions.Theme, validation.In("default")),
		"editor_options.font_family": validation.Validate(r.CustomEditorOptions.FontFamily, validation.In("default")),
	}.Filter()
	return err
}
