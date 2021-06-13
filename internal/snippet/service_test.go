package snippet

import (
	"context"
	"github.com/dd3v/snippets.ninja/internal/entity"
	"github.com/dd3v/snippets.ninja/internal/rbac"
	"github.com/dd3v/snippets.ninja/internal/test"
	"github.com/dd3v/snippets.ninja/pkg/query"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_GetByID(t *testing.T) {
	type args struct {
		id int
	}
	cases := []struct {
		name       string
		args       args
		rbac       test.RBACMock
		repository Repository
		wantData   entity.Snippet
		wantErr    error
	}{
		{
			name: "user can get snippet by ID",
			args: args{
				id: 1,
			},
			rbac: test.RBACMock{
				CanViewSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			repository: RepositoryMock{
				QueryByUserIDFn: nil,
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              1,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test snippet",
						Content:             "hello world",
						Language:            "go",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2020),
						UpdatedAt:           test.Time(2020),
					}, nil
				},
				CreateFn:        nil,
				UpdateFn:        nil,
				DeleteFn:        nil,
				CountByUserIDFn: nil,
			},
			wantData: entity.Snippet{
				ID:                  1,
				UserID:              1,
				Favorite:            false,
				AccessLevel:         0,
				Title:               "test snippet",
				Content:             "hello world",
				Language:            "go",
				CustomEditorOptions: entity.CustomEditorOptions{},
				CreatedAt:           test.Time(2020),
				UpdatedAt:           test.Time(2020),
			},
			wantErr: nil,
		},
		{
			name: "user does not have permission",
			args: args{
				id: 1,
			},
			rbac: test.RBACMock{
				CanViewSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return rbac.AccessError
				},
			},
			repository: RepositoryMock{
				QueryByUserIDFn: nil,
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              2,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test snippet",
						Content:             "hello world",
						Language:            "go",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2020),
						UpdatedAt:           test.Time(2020),
					}, nil
				},
				CreateFn:        nil,
				UpdateFn:        nil,
				DeleteFn:        nil,
				CountByUserIDFn: nil,
			},
			wantData: entity.Snippet{},
			wantErr:  rbac.AccessError,
		},
		{
			name: "repository error",
			args: args{
				id: 1,
			},
			rbac: test.RBACMock{
				CanViewSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			repository: RepositoryMock{
				QueryByUserIDFn: nil,
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{}, repositoryMockErr
				},
				CreateFn:        nil,
				UpdateFn:        nil,
				DeleteFn:        nil,
				CountByUserIDFn: nil,
			},
			wantData: entity.Snippet{},
			wantErr:  repositoryMockErr,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(tc.repository, tc.rbac)
			snippet, err := service.GetByID(context.Background(), tc.args.id)
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
		name       string
		args       args
		rbac       test.RBACMock
		repository Repository
		wantData   entity.Snippet
		wantErr    error
	}{
		{
			name: "user can create snippet",
			args: args{
				snippet: entity.Snippet{
					UserID:              1,
					Favorite:            false,
					AccessLevel:         0,
					Title:               "test",
					Content:             "test context",
					Language:            "php",
					CustomEditorOptions: entity.CustomEditorOptions{},
					CreatedAt:           test.Time(2020),
					UpdatedAt:           test.Time(2020),
				},
			},
			rbac: test.RBACMock{},
			repository: RepositoryMock{
				CreateFn: func(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              1,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test",
						Content:             "test context",
						Language:            "php",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2020),
						UpdatedAt:           test.Time(2020),
					}, nil
				},
			},
			wantData: entity.Snippet{
				ID:                  1,
				UserID:              1,
				Favorite:            false,
				AccessLevel:         0,
				Title:               "test",
				Content:             "test context",
				Language:            "php",
				CustomEditorOptions: entity.CustomEditorOptions{},
				CreatedAt:           test.Time(2020),
				UpdatedAt:           test.Time(2020),
			},
			wantErr: nil,
		},
		{
			name: "repository error",
			args: args{
				snippet: entity.Snippet{
					UserID:              1,
					Favorite:            false,
					AccessLevel:         0,
					Title:               "test",
					Content:             "test context",
					Language:            "php",
					CustomEditorOptions: entity.CustomEditorOptions{},
					CreatedAt:           test.Time(2020),
					UpdatedAt:           test.Time(2020),
				},
			},
			rbac: test.RBACMock{},
			repository: RepositoryMock{
				CreateFn: func(ctx context.Context, snippet entity.Snippet) (entity.Snippet, error) {
					return entity.Snippet{}, repositoryMockErr
				},
			},
			wantData: entity.Snippet{},
			wantErr:  repositoryMockErr,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(tc.repository, tc.rbac)
			snippet, err := service.Create(context.Background(), tc.args.snippet)
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
		name       string
		args       args
		rbac       test.RBACMock
		repository Repository
		wantData   entity.Snippet
		wantErr    error
	}{
		{
			name: "user can update snippet",
			args: args{
				snippet: entity.Snippet{
					ID:                  1,
					UserID:              1,
					Favorite:            false,
					AccessLevel:         0,
					Title:               "test",
					Content:             "test context",
					Language:            "php",
					CustomEditorOptions: entity.CustomEditorOptions{},
					CreatedAt:           test.Time(2020),
					UpdatedAt:           test.Time(2020),
				},
			},
			rbac: test.RBACMock{
				CanUpdateSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              1,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test",
						Content:             "test context",
						Language:            "php",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2020),
						UpdatedAt:           test.Time(2020),
					}, nil
				},
				UpdateFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			wantData: entity.Snippet{
				ID:                  1,
				UserID:              1,
				Favorite:            false,
				AccessLevel:         0,
				Title:               "test",
				Content:             "test context",
				Language:            "php",
				CustomEditorOptions: entity.CustomEditorOptions{},
				CreatedAt:           test.Time(2020),
				UpdatedAt:           test.Time(2020),
			},
			wantErr: nil,
		},
		{
			name: "repository error",
			args: args{
				snippet: entity.Snippet{
					ID:                  1,
					UserID:              1,
					Favorite:            false,
					AccessLevel:         0,
					Title:               "test",
					Content:             "test context",
					Language:            "php",
					CustomEditorOptions: entity.CustomEditorOptions{},
					CreatedAt:           test.Time(2020),
				},
			},
			rbac: test.RBACMock{
				CanUpdateSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.Snippet, error) {
					return entity.Snippet{
						ID:                  1,
						UserID:              1,
						Favorite:            false,
						AccessLevel:         0,
						Title:               "test",
						Content:             "test context",
						Language:            "php",
						CustomEditorOptions: entity.CustomEditorOptions{},
						CreatedAt:           test.Time(2020),
					}, nil
				},
				UpdateFn: func(ctx context.Context, snippet entity.Snippet) error {
					return repositoryMockErr
				},
			},
			wantData: entity.Snippet{},
			wantErr:  repositoryMockErr,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(tc.repository, tc.rbac)
			snippet, err := service.Update(context.Background(), tc.args.snippet)
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
		name       string
		args       args
		rbac       test.RBACMock
		repository RepositoryMock
		wantErr    error
	}{
		{
			name: "user can delete snippet",
			args: args{id: 1},
			rbac: test.RBACMock{
				CanDeleteSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
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
						CreatedAt:           time.Time{},
						UpdatedAt:           time.Time{},
					}, nil
				},
				DeleteFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			wantErr: nil,
		},
		{
			name: "repository error",
			args: args{id: 1},
			rbac: test.RBACMock{
				CanDeleteSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
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
						CreatedAt:           time.Time{},
						UpdatedAt:           time.Time{},
					}, nil
				},
				DeleteFn: func(ctx context.Context, snippet entity.Snippet) error {
					return repositoryMockErr
				},
			},
			wantErr: repositoryMockErr,
		},
		{
			name: "user does not have permissions",
			args: args{id: 1},
			rbac: test.RBACMock{
				CanDeleteSnippetFn: func(ctx context.Context, snippet entity.Snippet) error {
					return rbac.AccessError
				},
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
						CreatedAt:           time.Time{},
						UpdatedAt:           time.Time{},
					}, nil
				},
				DeleteFn: func(ctx context.Context, snippet entity.Snippet) error {
					return nil
				},
			},
			wantErr: rbac.AccessError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(tc.repository, tc.rbac)
			err := service.Delete(context.Background(), tc.args.id)
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
		name       string
		args       args
		repository Repository
		rbac       test.RBACMock
		wantData   int
		wantErr    error
	}{
		{
			name: "user name delete snippet",
			args: args{
				1, map[string]string{"sdfsdfsdf": "sdfsdfsd"},
			},
			repository: RepositoryMock{
				CountByUserIDFn: func(ctx context.Context, userID int, filter map[string]string) (int, error) {
					return 1, nil
				},
			},
			rbac:     test.RBACMock{},
			wantData: 1,
			wantErr:  nil,
		},
	}

	for _, tc := range cases {

		t.Run(tc.name, func(t *testing.T) {

			service := NewService(tc.repository, tc.rbac)
			count, err := service.CountByUserID(context.Background(), tc.args.userID, tc.args.filter)
			assert.Equal(t, tc.wantData, count)
			assert.Equal(t, tc.wantErr, err)

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
		name       string
		args       args
		repository RepositoryMock
		wantData   []entity.Snippet
		wantErr    error
	}{
		{
			name: "user can get snippets by conditions",
			args: args{
				userID: 1,
				filter: map[string]string{
					"favorite": "1",
				},
				sort:       query.NewSort("id", "asc"),
				pagination: query.NewPagination(1, 20),
			},
			repository: RepositoryMock{
				QueryByUserIDFn: func(ctx context.Context, userID int, filter map[string]string, sort query.Sort, pagination query.Pagination) ([]entity.Snippet, error) {
					return []entity.Snippet{
						{
							ID:                  1,
							UserID:              1,
							Favorite:            true,
							AccessLevel:         0,
							Title:               "test",
							Content:             "test",
							Language:            "test",
							CustomEditorOptions: entity.CustomEditorOptions{},
							CreatedAt:           test.Time(2020),
							UpdatedAt:           test.Time(2021),
						},
					}, nil
				},
			},
			wantData: []entity.Snippet{
				{
					ID:                  1,
					UserID:              1,
					Favorite:            true,
					AccessLevel:         0,
					Title:               "test",
					Content:             "test",
					Language:            "test",
					CustomEditorOptions: entity.CustomEditorOptions{},
					CreatedAt:           test.Time(2020),
					UpdatedAt:           test.Time(2021),
				},
			},
			wantErr: nil,
		},
		{
			name: "repository error",
			args: args{
				userID: 1,
				filter: map[string]string{
					"favorite": "1",
				},
				sort:       query.NewSort("id", "asc"),
				pagination: query.NewPagination(1, 20),
			},
			repository: RepositoryMock{
				QueryByUserIDFn: func(ctx context.Context, userID int, filter map[string]string, sort query.Sort, pagination query.Pagination) ([]entity.Snippet, error) {
					return []entity.Snippet{}, repositoryMockErr
				},
			},
			wantData: []entity.Snippet{},
			wantErr:  repositoryMockErr,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(tc.repository, test.RBACMock{})
			snippets, err := service.QueryByUserID(context.Background(), tc.args.userID, tc.args.filter, tc.args.sort, tc.args.pagination)
			assert.Equal(t, tc.wantData, snippets)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
