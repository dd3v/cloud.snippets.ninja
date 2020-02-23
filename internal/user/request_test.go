package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequestValidation(t *testing.T) {
	cases := []struct {
		name  string
		model CreateRequest
		fail  bool
	}{
		{
			"success",
			CreateRequest{
				Login:          "test",
				Email:          "test@mailservice.com",
				Password:       "qwerty",
				RepeatPassword: "qwerty",
			},
			false,
		},
		{
			"invalid email",
			CreateRequest{
				Login:          "test",
				Email:          "testmailservice.com",
				Password:       "qwerty",
				RepeatPassword: "qwerty",
			},
			true,
		},
		{
			"password confirmation",
			CreateRequest{
				Login:          "test",
				Email:          "testmail@service.com",
				Password:       "qwerty",
				RepeatPassword: "123456",
			},
			true,
		},
		{
			"length",
			CreateRequest{
				Login:          "sadfjk32149sadfmzkdrjk324sadfjk32149sadfmzkdrjk324",
				Email:          "testmail@service.com",
				Password:       "qwerty",
				RepeatPassword: "123456",
			},
			true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.fail, err != nil)
		})
	}
}

func TestUpdateRequestValidation(t *testing.T) {
	cases := []struct {
		name  string
		model UpdateRequest
		fail  bool
	}{
		{
			"success",
			UpdateRequest{
				Website: "github.com",
			},
			false,
		},
		{
			"invalid url",
			UpdateRequest{
				Website: "gith@@#ubcom",
			},
			true,
		},
		{
			"length",
			UpdateRequest{
				Website: "sadfjk32149sadfmzkdrjk324sadfjk32149sadfmzkdrjk324sadfjk32149sadfmzkdrjk324sadfjk32149sadfmzkdrjk3243s",
			},
			true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.fail, err != nil)
		})
	}
}
