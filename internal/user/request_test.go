package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequestValidation(t *testing.T) {
	cases := []struct {
		name    string
		request createRequest
		fail    bool
	}{
		{
			"success",
			createRequest{
				Login:          "test",
				Email:          "test@mailservice.com",
				Password:       "qwerty",
				RepeatPassword: "qwerty",
			},
			false,
		},
		{
			"invalid email",
			createRequest{
				Login:          "test",
				Email:          "testmailservice.com",
				Password:       "qwerty",
				RepeatPassword: "qwerty",
			},
			true,
		},
		{
			"password confirmation",
			createRequest{
				Login:          "test",
				Email:          "testmail@service.com",
				Password:       "qwerty",
				RepeatPassword: "123456",
			},
			true,
		},
		{
			"length",
			createRequest{
				Login:          "sadfjk32149sadfmzkdrjk324sadfjk32149sadfmzkdrjk324",
				Email:          "testmail@service.com",
				Password:       "qwerty",
				RepeatPassword: "123456",
			},
			true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.Validate()
			assert.Equal(t, tc.fail, err != nil)
		})
	}
}

func TestUpdateRequestValidation(t *testing.T) {
	cases := []struct {
		name    string
		request updateRequest
		fail    bool
	}{
		{
			"success",
			updateRequest{
				Website: "github.com",
			},
			false,
		},
		{
			"invalid url",
			updateRequest{
				Website: "gith@@#ubcom",
			},
			true,
		},
		{
			"length",
			updateRequest{
				Website: "sadfjk32149sadfmzkdrjk324sadfjk32149sadfmzkdrjk324sadfjk32149sadfmzkdrjk324sadfjk32149sadfmzkdrjk3243s",
			},
			true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.Validate()
			assert.Equal(t, tc.fail, err != nil)
		})
	}
}
