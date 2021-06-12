package user

import (
	"context"
	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GetByID(t *testing.T) {
	type args struct {
		id int
	}
	cases := []struct {
		name       string
		repository Repository
		args       args
		wantData   entity.User
		wantErr    error
	}{
		{
			name: "get user by id",
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.User, error) {
					return entity.User{
						ID:        1,
						Password:  "$2a$10$ceGJobOZUCIVM72m9fMVZO.NjjQcaadIhJnQEE7Cdq/QuBze9yZAq",
						Login:     "dd3v",
						Email:     "dmitriy.d3v@gmail.com",
						CreatedAt: test.Time(2020),
						UpdatedAt: test.Time(2021),
					}, nil
				},
			},
			args: args{id: 1},
			wantData: entity.User{
				ID:        1,
				Password:  "$2a$10$ceGJobOZUCIVM72m9fMVZO.NjjQcaadIhJnQEE7Cdq/QuBze9yZAq",
				Login:     "dd3v",
				Email:     "dmitriy.d3v@gmail.com",
				CreatedAt: test.Time(2020),
				UpdatedAt: test.Time(2021),
			},
			wantErr: nil,
		},
		{
			name: "repository errpr",
			repository: RepositoryMock{
				GetByIDFn: func(ctx context.Context, id int) (entity.User, error) {
					return entity.User{}, repositoryMockErr
				},
			},
			args:     args{id: 1},
			wantData: entity.User{},
			wantErr:  repositoryMockErr,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(tc.repository)
			user, err := service.GetByID(context.Background(), tc.args.id)
			assert.Equal(t, tc.wantData, user)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestService_Create(t *testing.T) {
	type args struct {
		user entity.User
	}
	cases := []struct {
		name       string
		repository Repository
		args       args
		wantData   entity.User
		wantErr    error
	}{
		{
			name: "user can create new user",
			repository: RepositoryMock{
				CreateFn: func(ctx context.Context, user entity.User) (entity.User, error) {
					return entity.User{
						ID:        1,
						Password:  "$2a$10$ceGJobOZUCIVM72m9fMVZO.NjjQcaadIhJnQEE7Cdq/QuBze9yZAq",
						Login:     "dd3v",
						Email:     "dd3v@gmail.com",
						CreatedAt: test.Time(2021),
						UpdatedAt: test.Time(2021),
					}, nil
				},
				ExistsFn: func(ctx context.Context, field string, value string) (bool, error) {
					return false, nil
				},
			},
			args: args{user: entity.User{
				ID:        1,
				Password:  "qwerty",
				Login:     "dd3v",
				Email:     "dd3v@gmail.com",
				CreatedAt: test.Time(2021),
				UpdatedAt: test.Time(2021),
			}},
			wantData: entity.User{
				ID:        1,
				Password:  "$2a$10$ceGJobOZUCIVM72m9fMVZO.NjjQcaadIhJnQEE7Cdq/QuBze9yZAq",
				Login:     "dd3v",
				Email:     "dd3v@gmail.com",
				CreatedAt: test.Time(2021),
				UpdatedAt: test.Time(2021),
			},
			wantErr: nil,
		},
		{
			name: "repository error",
			repository: RepositoryMock{
				CreateFn: func(ctx context.Context, user entity.User) (entity.User, error) {
					return entity.User{}, repositoryMockErr
				},
				ExistsFn: func(ctx context.Context, field string, value string) (bool, error) {
					return false, nil
				},
			},
			args: args{user: entity.User{
				ID:        1,
				Password:  "qwerty",
				Login:     "dd3v",
				Email:     "dd3v@gmail.com",
				CreatedAt: test.Time(2021),
				UpdatedAt: test.Time(2021),
			}},
			wantData: entity.User{},
			wantErr:  repositoryMockErr,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(tc.repository)
			user, err := service.Create(context.Background(), tc.args.user)
			assert.Equal(t, tc.wantData, user)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestService_Exists(t *testing.T) {

	type args struct {
		field string
		value string
	}

	cases := []struct {
		name       string
		repository Repository
		args       args
		wantData   bool
		wantErr    error
	}{
		{
			name: "email is unique",
			repository: RepositoryMock{
				ExistsFn: func(ctx context.Context, field string, value string) (bool, error) {
					return false, nil
				},
			},
			args: args{
				field: "email",
				value: "test@gmail.com",
			},
			wantData: false,
			wantErr:  nil,
		},
		{
			name: "email is not unique",
			repository: RepositoryMock{
				ExistsFn: func(ctx context.Context, field string, value string) (bool, error) {
					return true, nil
				},
			},
			args: args{
				field: "email",
				value: "test@gmail.com",
			},
			wantData: true,
			wantErr:  nil,
		},
		{
			name: "repository error",
			repository: RepositoryMock{
				ExistsFn: func(ctx context.Context, field string, value string) (bool, error) {
					return false, repositoryMockErr
				},
			},
			args: args{
				field: "email",
				value: "test@gmail.com",
			},
			wantData: false,
			wantErr:  repositoryMockErr,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(tc.repository)
			exists, err := service.Exists(context.Background(), tc.args.field, tc.args.value)
			assert.Equal(t, tc.wantData, exists)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
