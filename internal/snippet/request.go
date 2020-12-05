package snippet

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gopkg.in/guregu/null.v4"
)

//QuerySnippetsRequest - ...
type QuerySnippetsRequest struct {
	Limit    int       `form:"limit"`
	Offset   int       `form:"offset"`
	Favorite null.Bool `form:"favorite"`
	Public   null.Bool `form:"public"`
	Keywords string    `form:"keywords"`
}

// //NewQuerySnippetsRequest -
// func NewQuerySnippetsRequest() QuerySnippetsRequest {
// 	return QuerySnippetsRequest{Limit: 100, Offset: 0, Favorite: nil, Public: nil, Keywords: "qqq"}
// }

//Validate - ...
func (r QuerySnippetsRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Limit, validation.Min(1), validation.Max(250)),
		validation.Field(&r.Offset, validation.Min(0)),
		validation.Field(&r.Keywords, validation.Length(1, 50)),
	)
}
