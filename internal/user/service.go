package user

import (
	"context"
	"github.com/dd3v/snippets.ninja/internal/entity"
	"github.com/dd3v/snippets.ninja/pkg/security"
)

//Service - ...
type Service interface {
	GetByID(ctx context.Context, id int) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Exists(ctx context.Context, field string, value string) (bool, error)
}

//Repository - ...
type Repository interface {
	GetByID(ctx context.Context, id int) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, field string, value string) (bool, error)
}

type service struct {
	repository Repository
}

//NewService - ...
func NewService(repository Repository) Service {
	return service{
		repository: repository,
	}
}

func (s service) Exists(ctx context.Context, field string, value string) (bool, error) {
	return s.repository.Exists(ctx, field, value)
}

func (s service) GetByID(ctx context.Context, id int) (entity.User, error) {
	return s.repository.GetByID(ctx, id)
}

func (s service) Create(ctx context.Context, user entity.User) (entity.User, error) {

	passwordHash, err := security.GenerateHashFromPassword(user.Password)
	if err != nil {
		return entity.User{}, err
	}
	user.Password = passwordHash
	result, err := s.repository.Create(ctx, user)

	return result, err
}
