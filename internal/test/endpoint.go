package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/stretchr/testify/assert"
)

//APITestCase - ...
type APITestCase struct {
	Method       string
	URL          string
	Body         string
	Header       http.Header
	WantStatus   int
	WantResponse string
}

//Endpoint - ...
func Endpoint(t *testing.T, name string, router *routing.Router, tc APITestCase) {
	t.Run(name, func(t *testing.T) {
		req, err := http.NewRequest(tc.Method, tc.URL, bytes.NewBufferString(tc.Body))
		if err != nil {
			//TODO: probably t.Error() will be enough
			t.Error(err)
			t.FailNow()
		}
		if tc.Header != nil {
			req.Header = tc.Header
		}
		if req.Header.Get("Content-Type") == "" {
			req.Header.Set("Content-Type", "application/json")
		}
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, tc.WantStatus, res.Code)
		if tc.WantResponse != "" {
			pattern := strings.Trim(tc.WantResponse, "*")
			if pattern != tc.WantResponse {
				assert.Contains(t, res.Body.String(), pattern, "response mismatch")
			} else {
				assert.JSONEq(t, tc.WantResponse, res.Body.String(), "response mismatch")
			}
		}
	})
}
