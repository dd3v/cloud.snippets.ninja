package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginRequest(t *testing.T) {
	cases := []struct {
		name    string
		request LoginRequest
		fail    bool
	}{
		{"success", LoginRequest{Login: "admin", Password: "qwerty"}, false},
		{"invalid login", LoginRequest{Login: "a", Password: "qwerty"}, true},
		{"invalid password", LoginRequest{Login: "user_100", Password: ""}, true},
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
		request RefreshRequest
		fail    bool
	}{
		{"success", RefreshRequest{RefreshToken: "sdfsdfsdf"}, false},
		{"empty refresh token", RefreshRequest{RefreshToken: ""}, true},
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
		name  string
		model LogoutRequest
		fail  bool
	}{
		{"success", LogoutRequest{RefreshToken: "refresh_token"}, false},
		{"fail", LogoutRequest{RefreshToken: ""}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.model.Validate()
			assert.Equal(t, tc.fail, err != nil)
		})
	}
}
