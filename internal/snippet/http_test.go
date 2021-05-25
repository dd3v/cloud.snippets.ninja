package snippet

import (
	"net/http"
	"testing"

	"github.com/dd3v/snippets.page.backend/internal/snippet/mock"
	"github.com/dd3v/snippets.page.backend/internal/test"
)

func TestSnippetEndpoint(t *testing.T) {

	cases := []test.APITestCase{
		{
			Name:         "Get snippets (unauthorized)",
			Method:       http.MethodGet,
			URL:          "/snippets",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusUnauthorized,
			WantResponse: "",
		},
		{
			Name:         "Get Snippets (authorized)",
			Method:       http.MethodGet,
			URL:          "/snippets",
			Body:         "",
			Header:       test.MockAuthHeader(),
			WantStatus:   http.StatusOK,
			WantResponse: "",
		},
		{
			Name:         "Create snippet (invalid request)",
			Method:       http.MethodPost,
			URL:          "/snippets",
			Body:        `{"title":""}`,
			Header:       test.MockAuthHeader(),
			WantStatus:   http.StatusBadRequest,
			WantResponse: "",
		},
		{
			Name:         "Create snippet (success)",
			Method:       http.MethodPost,
			URL:          "/snippets",
			Body:        `{"title":"hello"}`,
			Header:       test.MockAuthHeader(),
			WantStatus:   http.StatusCreated,
			WantResponse: "",
		},
	}

	router := test.MockRouter()
	service := NewService(mock.NewRepository(), test.RBACMock{})

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
			test.Endpoint(t, router, tc)
		})
	}
}
