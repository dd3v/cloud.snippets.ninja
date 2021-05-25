package snippet

import (
	"context"
	"database/sql"
	"github.com/dd3v/snippets.page.backend/internal/snippet/mock"
	"testing"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/rbac"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"github.com/dd3v/snippets.page.backend/pkg/query"
	"github.com/stretchr/testify/assert"
)


func TestService_GetByID(t *testing.T) {
	type args struct {
		id int
	}

	cases := []struct {
		name     string
		args     args
		rbac     test.RBACMock
		wantData entity.Snippet
		wantErr  error
	}{
		{
			"not found",
			args{
				id: 123,
			},
			test.RBACMock{
				CanViewSnippetFn: func(context.Context, entity.Snippet) error {
					return nil
				},
			},
			entity.Snippet{},
			sql.ErrNoRows,
		},
		{
			"success",
			args{
				id: 1,
			},
			test.RBACMock{
				CanViewSnippetFn: func(context.Context, entity.Snippet) error {
					return nil
				},
			},
			entity.Snippet{
				ID:                  1,
				UserID:              1,
				Favorite:            true,
				AccessLevel:         0,
				Title:               "PHP hello world",
				Content:             "<?php echo 'Hello world'; ?>",
				Language:            "php",
				CustomEditorOptions: entity.CustomEditorOptions{},
				CreatedAt:           test.Time(2020),
				UpdatedAt:           test.Time(2021),
			},
			nil,
		},
		{
			"forbidden",
			args{
				id: 2,
			},
			test.RBACMock{
				CanViewSnippetFn: func(context.Context, entity.Snippet) error {
					return rbac.AccessError
				},
			},
			entity.Snippet{},
			rbac.AccessError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := mock.NewRepository()
			s := NewService(mockRepository, tc.rbac)
			snippet, err := s.GetByID(context.Background(), tc.args.id)
			assert.Equal(t, tc.wantData, snippet)
			assert.Equal(t, tc.wantErr, err)
		})
	}

}

func TestService_Create(t *testing.T) {

	type args struct {
		snippet entity.Snippet
	}

	cases := []struct {
		name     string
		args     args
		wantData entity.Snippet
		wantErr  error
	}{
		{
			"success",
			args{
				snippet: entity.Snippet{},
			},
			entity.Snippet{},
			nil,
		},
		{
			"fail",
			args{
				snippet: entity.Snippet{
					ID:    1,
					Title: "error",
				},
			},
			entity.Snippet{},
			mock.ErrorRepository,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := mock.NewRepository()
			s := NewService(mockRepository, rbac.New())
			snippet, err := s.Create(context.Background(), tc.args.snippet)
			assert.Equal(t, tc.wantData, snippet)
			assert.Equal(t, tc.wantErr, err)
		})
	}

}

func TestService_Update(t *testing.T) {

	type args struct {
		snippet entity.Snippet
	}

	cases := []struct {
		name     string
		args     args
		rbac     test.RBACMock
		wantData entity.Snippet
		wantErr  error
	}{
		{
			"success",
			args{
				snippet: entity.Snippet{
					ID:                  1,
					UserID:              1,
					Favorite:            false,
					AccessLevel:         0,
					Title:               "PHP hello world - updated",
					Content:             "<?php echo 'Hello world'; ?>",
					Language:            "php",
					CustomEditorOptions: entity.CustomEditorOptions{},
					CreatedAt:           test.Time(2020),
					UpdatedAt:           test.Time(2021),
				},
			},
			test.RBACMock{
				CanUpdateSnippetFn: func(context.Context, entity.Snippet) error {
					return nil
				},
			},
			entity.Snippet{
				ID:                  1,
				UserID:              1,
				Favorite:            false,
				AccessLevel:         0,
				Title:               "PHP hello world - updated",
				Content:             "<?php echo 'Hello world'; ?>",
				Language:            "php",
				CustomEditorOptions: entity.CustomEditorOptions{},
				CreatedAt:           test.Time(2020),
				UpdatedAt:           test.Time(2021),
			},
			nil,
		},
		{
			"repository fail",
			args{
				snippet: entity.Snippet{
					ID:    1,
					Title: "error",
				},
			},
			test.RBACMock{
				CanUpdateSnippetFn: func(context.Context, entity.Snippet) error {
					return nil
				},
			},
			entity.Snippet{},
			mock.ErrorRepository,
		},
		{
			"rbac fail",
			args{
				snippet: entity.Snippet{
					ID:                  1,
					UserID:              1,
					Favorite:            false,
					AccessLevel:         0,
					Title:               "PHP hello world - updated",
					Content:             "<?php echo 'Hello world'; ?>",
					Language:            "php",
					CustomEditorOptions: entity.CustomEditorOptions{},
					CreatedAt:           test.Time(2020),
					UpdatedAt:           test.Time(2021),
				},
			},
			test.RBACMock{
				CanUpdateSnippetFn: func(context.Context, entity.Snippet) error {
					return rbac.AccessError
				},
			},
			entity.Snippet{},
			rbac.AccessError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := mock.NewRepository()
			s := NewService(mockRepository, tc.rbac)
			snippet, err := s.Update(context.Background(), tc.args.snippet)
			assert.Equal(t, tc.wantData, snippet)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestService_Delete(t *testing.T) {

	type args struct {
		id int
	}

	cases := []struct {
		name    string
		args    args
		rbac    test.RBACMock
		wantErr error
	}{
		{
			"success",
			args{id: 1},
			test.RBACMock{
				CanDeleteSnippetFn: func(context.Context, entity.Snippet) error {
					return nil
				},
			},
			nil,
		},
		{
			"not found",
			args{id: 123},
			test.RBACMock{
				CanDeleteSnippetFn: func(context.Context, entity.Snippet) error {
					return nil
				},
			},
			sql.ErrNoRows,
		},
		{
			"forbidden",
			args{id: 2},
			test.RBACMock{
				CanDeleteSnippetFn: func(context.Context, entity.Snippet) error {
					return rbac.AccessError
				},
			},
			rbac.AccessError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := mock.NewRepository()
			s := NewService(mockRepository, tc.rbac)
			err := s.Delete(context.Background(), tc.args.id)
			assert.Equal(t, tc.wantErr, err)
		})
	}

}

func TestService_CountByUserID(t *testing.T) {

	type args struct {
		userID int
		filter map[string]string
	}

	cases := []struct {
		name     string
		args     args
		rbac     test.RBACMock
		wantData int
		wantErr  error
	}{
		{
			"success",
			args{userID: 1, filter: map[string]string{}},
			test.RBACMock{},
			3,
			nil,
		},
		{
			"repository error",
			args{userID: 0, filter: map[string]string{}},
			test.RBACMock{},
			0,
			mock.ErrorRepository,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := mock.NewRepository()
			s := NewService(mockRepository, tc.rbac)
			count, err := s.CountByUserID(context.Background(), tc.args.userID, tc.args.filter)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantData, count)
		})
	}

}

func TestService_QueryByUserID(t *testing.T) {

	type args struct {
		userID     int
		filter     map[string]string
		sort       query.Sort
		pagination query.Pagination
	}

	cases := []struct {
		name    string
		args    args
		rbac    test.RBACMock
		wantErr error
	}{
		{
			"success",
			args{
				userID:     1,
				filter:     map[string]string{},
				sort:       query.NewSort("id", "asc"),
				pagination: query.NewPagination(1, 10),
			},
			test.RBACMock{},
			nil,
		},
	}

	for _, tc := range cases {

		t.Run(tc.name, func(t *testing.T) {
			mockRepository := mock.NewRepository()
			s := NewService(mockRepository, tc.rbac)
			snippets, err := s.QueryByUserID(context.Background(), tc.args.userID, tc.args.filter, tc.args.sort, tc.args.pagination)
			assert.NotEmpty(t, snippets)
			assert.Equal(t, tc.wantErr, err)
		})

	}

}
