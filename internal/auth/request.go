package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

//AuthRequest - ...
type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (r AuthRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Login, validation.Required, validation.Length(2, 50)),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 50)),
	)
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (r RefreshRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.RefreshToken, validation.Required),
	)
}
