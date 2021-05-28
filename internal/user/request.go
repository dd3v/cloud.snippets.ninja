package user

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v3/is"
)

//createRequest - ...
type createRequest struct {
	Login          string `json:"login"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

func stringEquals(str string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if s != str {
			return errors.New("password and repeat password should be equal")
		}
		return nil
	}
}

//Validate - ...
func (r createRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Login, validation.Required, validation.Length(2, 50)),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 50)),
		validation.Field(&r.RepeatPassword, validation.Required, validation.Length(6, 50)),
		validation.Field(&r.Password, validation.By(stringEquals(r.RepeatPassword))),
	)
}

//updateRequest - ...
type updateRequest struct {
	Website string `json:"website"`
}

//Validate - ...
func (u updateRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Website, validation.Length(5, 100), is.URL),
	)
}
