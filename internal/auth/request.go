package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

//AuthRequest - ...
type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (r loginRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Login, validation.Required, validation.Length(2, 50)),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 50)),
	)
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (r refreshRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.RefreshToken, validation.Required),
	)
}

type logoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (r logoutRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.RefreshToken, validation.Required),
	)
}
