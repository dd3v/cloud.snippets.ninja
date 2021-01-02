package snippet

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gopkg.in/guregu/null.v4"
	"time"
)

//QuerySnippetsRequest - ...
type OwnSnippetsRequest struct {
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
	Favorite int    `form:"favorite"`
	Access   int    `form:"access"`
	Title    string `form:"title"`
}

//OwnSnippetsRequest -
func NewOwnSnippetsRequest() OwnSnippetsRequest {
	return OwnSnippetsRequest{Limit: 100, Offset: 0, Favorite: -1, Access: -1, Title: ""}
}

//Validate - ...
func (r OwnSnippetsRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Limit, validation.Min(1), validation.Max(250)),
		validation.Field(&r.Offset, validation.Min(0)),
		validation.Field(&r.Favorite, validation.Min(-1), validation.Max(1)),
		validation.Field(&r.Access, validation.Min(-1), validation.Max(1)),
		validation.Field(&r.Title, validation.Length(0, 50)),
	)
}

type CreateSnippetRequest struct {
	Favorite      bool      `from:"favorite"`
	Access        int       `from:"access"`
	Title         string    `from:"title"`
	Content       null.String    `from:"content"`
	FileExtension string    `from:"file_extension"`
	EditorOptions struct{}  `from:"editor_options"`
	CreatedAt     time.Time `from:"created_at"`
	UpdatedAt     time.Time `from:"updated_at"`

}

func NewCreateSnippetRequest() CreateSnippetRequest {
	return CreateSnippetRequest{Title: ""}
}

func (r CreateSnippetRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Favorite, validation.Min(0), validation.Max(1)),
		validation.Field(&r.Access, validation.Min(0), validation.Max(1)),
		validation.Field(&r.Title, validation.Length(0, 50)),
	)
}

type UpdateSnippetRequest struct {
	Favorite      bool      `from:"favorite"`
	Access        int       `from:"access"`
	Title         string    `from:"title"`
	Content       string    `from:"content"`
	FileExtension string    `from:"file_extension"`
	EditorOptions struct{}  `from:"editor_options"`
	CreatedAt     time.Time `from:"created_at"`
	UpdatedAt     time.Time `from:"updated_at"`

}

func NewUpdateSnippetRequest() UpdateSnippetRequest {
	return UpdateSnippetRequest{Title: ""}
}

func (r UpdateSnippetRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Favorite, validation.Min(0), validation.Max(1)),
		validation.Field(&r.Access, validation.Min(0), validation.Max(1)),
		validation.Field(&r.Title, validation.Length(0, 50)),
	)
}
