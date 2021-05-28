package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginRequest(t *testing.T) {
	cases := []struct {
		name    string
		request loginRequest
		fail    bool
	}{
		{"success", loginRequest{Login: "admin", Password: "qwerty"}, false},
		{"invalid login", loginRequest{Login: "a", Password: "qwerty"}, true},
		{"invalid password", loginRequest{Login: "user_100", Password: ""}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.Validate()
			assert.Equal(t, tc.fail, err != nil)
		})
	}
}

func TestRefreshRequest(t *testing.T) {
	cases := []struct {
		name    string
		request refreshRequest
		fail    bool
	}{
		{"success", refreshRequest{RefreshToken: "sdfsdfsdf"}, false},
		{"empty refresh token", refreshRequest{RefreshToken: ""}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.Validate()
			assert.Equal(t, tc.fail, err != nil)
		})
	}
}

func testLogoutRequest(t *testing.T) {
	cases := []struct {
		name    string
		request logoutRequest
		fail    bool
	}{
		{"success", logoutRequest{RefreshToken: "refresh_token"}, false},
		{"fail", logoutRequest{RefreshToken: ""}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.Validate()
			assert.Equal(t, tc.fail, err != nil)
		})
	}
}
