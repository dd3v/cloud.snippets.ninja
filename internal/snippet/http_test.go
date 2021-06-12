package snippet

import (
	"context"
	"database/sql"
	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/rbac"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"github.com/dd3v/snippets.page.backend/pkg/query"
	"net/http"
	"testing"
)

func TestHTTP_View(t *testing.T) {
	cases := []struct {
		name       string
		request    test.APITestCase
		repository Repository
		rbac       RBAC
	}{
		{
			name: "user can view own snippet",
			request: test.APITestCase{
				Method:       http.MethodGet,
				URL:          "/snippets/1",
				Body:         "",
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusOK,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              1,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test",
						Content:             "test",
						Language:            "go",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2020),
						UpdatedAt:           test.Time(2021),
					}, nil
				},
			},
			rbac: test.RBACMock{
				CanViewSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
		},
		{
			name: "user can't view snippet",
			request: test.APITestCase{
				Method:       http.MethodGet,
				URL:          "/snippets/1",
				Body:         "",
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusForbidden,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              2,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test",
						Content:             "test",
						Language:            "go",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2020),
						UpdatedAt:           test.Time(2021),
					}, nil
				},
			},
			rbac: test.RBACMock{
				CanViewSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return rbac.AccessError
				},
			},
		},
		{
			name: "repository error",
			request: test.APITestCase{
				Method:       http.MethodGet,
				URL:          "/snippets/1",
				Body:         "",
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{}, repositoryMockErr
				},
			},
			rbac: test.RBACMock{
				CanViewSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
		},
	}

	for _, tc := range cases {
		router := test.MockRouter()
		service := NewService(tc.repository, tc.rbac)
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, tc.name, router, tc.request)
	}
}

func TestHTTP_Create(t *testing.T) {
	cases := []struct {
		name       string
		request    test.APITestCase
		repository Repository
		rbac       RBAC
	}{
		{
			name: "user can create snippet",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/snippets",
				Body:         `{"favorite":false, "access_level":0,"title": "test"}`,
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusCreated,
				WantResponse: "*test*",
			},
			repository: RepositoryMock{
				CreateFn: func(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              1,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test",
						Content:             "test",
						Language:            "test",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2021),
						UpdatedAt:           test.Time(2021),
					}, nil
				},
			},
			rbac: test.RBACMock{},
		},
		{
			name: "request validation error",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/snippets",
				Body:         `{"favorite":false, "access_level":4,"title": "test"}`,
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusBadRequest,
				WantResponse: "",
			},
			repository: RepositoryMock{
				CreateFn: func(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              1,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test",
						Content:             "test",
						Language:            "test",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2021),
						UpdatedAt:           test.Time(2021),
					}, nil
				},
			},
			rbac: test.RBACMock{},
		},
		{
			name: "repository error",
			request: test.APITestCase{
				Method:       http.MethodPost,
				URL:          "/snippets",
				Body:         `{"favorite":false, "access_level":0,"title": "test"}`,
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repository: RepositoryMock{
				CreateFn: func(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error) {
					return entity.Snippet{}, repositoryMockErr
				},
			},
			rbac: test.RBACMock{},
		},
	}

	for _, tc := range cases {
		router := test.MockRouter()
		service := NewService(tc.repository, tc.rbac)
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, tc.name, router, tc.request)
	}
}

func TestHTTP_Update(t *testing.T) {
	cases := []struct {
		name       string
		request    test.APITestCase
		repository Repository
		rbac       RBAC
	}{
		{
			name: "user can update snippet",
			request: test.APITestCase{
				Method:       http.MethodPut,
				URL:          "/snippets/1",
				Body:         `{"favorite":false, "access_level":0,"title": "this is new test"}`,
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusOK,
				WantResponse: "*this is new test*",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              1,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test",
						Content:             "test",
						Language:            "test",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2020),
						UpdatedAt:           test.Time(2021),
					}, nil
				},

				UpdateFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			rbac: test.RBACMock{
				CanUpdateSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
		},
		{
			name: "user can't update snippet, permission error",
			request: test.APITestCase{
				Method:       http.MethodPut,
				URL:          "/snippets/1",
				Body:         `{"favorite":false, "access_level":0,"title": "this is new test"}`,
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusForbidden,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              1,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test",
						Content:             "test",
						Language:            "test",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2020),
						UpdatedAt:           test.Time(2021),
					}, nil
				},

				UpdateFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			rbac: test.RBACMock{
				CanUpdateSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return rbac.AccessError
				},
			},
		},
		{
			name: "404",
			request: test.APITestCase{
				Method:       http.MethodPut,
				URL:          "/snippets/1",
				Body:         `{"favorite":false, "access_level":0,"title": "this is new test"}`,
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusNotFound,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{}, sql.ErrNoRows
				},

				UpdateFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			rbac: test.RBACMock{
				CanUpdateSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
		},
		{
			name: "repository error",
			request: test.APITestCase{
				Method:       http.MethodPut,
				URL:          "/snippets/1",
				Body:         `{"favorite":false, "access_level":0,"title": "this is new test"}`,
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{}, repositoryMockErr
				},

				UpdateFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			rbac: test.RBACMock{
				CanUpdateSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
		},
	}
	for _, tc := range cases {
		router := test.MockRouter()
		service := NewService(tc.repository, tc.rbac)
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, tc.name, router, tc.request)
	}
}

func TestHTTP_Delete(t *testing.T) {
	cases := []struct {
		name       string
		request    test.APITestCase
		repository Repository
		rbac       RBAC
	}{
		{
			name: "user can delete",
			request: test.APITestCase{
				Method:       http.MethodDelete,
				URL:          "/snippets/1",
				Body:         "",
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusOK,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              1,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test",
						Content:             "test",
						Language:            "test",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2020),
						UpdatedAt:           test.Time(2021),
					}, nil
				},
				DeleteFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			rbac: test.RBACMock{
				CanDeleteSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
		},
		{
			name: "permission error",
			request: test.APITestCase{
				Method:       http.MethodDelete,
				URL:          "/snippets/1",
				Body:         "",
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusForbidden,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              1,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test",
						Content:             "test",
						Language:            "test",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2020),
						UpdatedAt:           test.Time(2021),
					}, nil
				},
				DeleteFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			rbac: test.RBACMock{
				CanDeleteSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return rbac.AccessError
				},
			},
		},
		{
			name: "repository error",
			request: test.APITestCase{
				Method:       http.MethodDelete,
				URL:          "/snippets/1",
				Body:         "",
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{}, repositoryMockErr
				},
				DeleteFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			rbac: test.RBACMock{
				CanDeleteSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return rbac.AccessError
				},
			},
		},
	}
	for _, tc := range cases {
		router := test.MockRouter()
		service := NewService(tc.repository, tc.rbac)
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, tc.name, router, tc.request)
	}
}

func TestHTTP_List(t *testing.T) {
	cases := []struct {
		name       string
		request    test.APITestCase
		repository Repository
		rbac       RBAC
	}{
		{
			name: "user cat filter snippet",
			request: test.APITestCase{
				Method:       http.MethodGet,
				URL:          "/snippets?favorite=true",
				Body:         "",
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusOK,
				WantResponse: "*test snippet*",
			},
			repository: RepositoryMock{
				QueryByUserIDFn: func(ctx context.Context, userID int, filter map[string]string, sort query.Sort, pagination query.Pagination) ([]entity.Snippet, error) {
					return []entity.Snippet{
						{
							ID:                  1,
							UserID:              1,
							Favorite:            false,
							AccessLevel:         0,
							Title:               "test snippet",
							Content:             "test",
							Language:            "php",
							CustomEditorOptions: entity.CustomEditorOptions{},
							CreatedAt:           test.Time(2020),
							UpdatedAt:           test.Time(2021),
						},
					}, nil
				},
				CountByUserIDFn: func(ctx context.Context, userID int, filter map[string]string) (int, error) {
					return 0, nil
				},
			},
			rbac: test.RBACMock{},
		},
		{
			name: "repository error",
			request: test.APITestCase{
				Method:       http.MethodGet,
				URL:          "/snippets?favorite=true",
				Body:         "",
				Header:       test.MockAuthHeader(),
				WantStatus:   http.StatusInternalServerError,
				WantResponse: "",
			},
			repository: RepositoryMock{
				QueryByUserIDFn: func(ctx context.Context, userID int, filter map[string]string, sort query.Sort, pagination query.Pagination) ([]entity.Snippet, error) {
					return []entity.Snippet{}, repositoryMockErr
				},
				CountByUserIDFn: func(ctx context.Context, userID int, filter map[string]string) (int, error) {
					return 0, nil
				},
			},
			rbac: test.RBACMock{},
		},
	}
	for _, tc := range cases {
		router := test.MockRouter()
		service := NewService(tc.repository, tc.rbac)
		NewHTTPHandler(router.Group(""), test.MockAuthMiddleware, service)
		test.Endpoint(t, tc.name, router, tc.request)
	}
}
