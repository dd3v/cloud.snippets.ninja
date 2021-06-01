package user

import (
	"context"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/security"
)

//Service - ...
type Service interface {
	List(ctx context.Context, limit int, offset int) ([]entity.User, error)
	GetByID(ctx context.Context, id int) (entity.User, error)
	Create(ctx context.Context, request createRequest) (entity.User, error)
	Update(ctx context.Context, id int, request updateRequest) (entity.User, error)
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int, error)
}

//Repository - ...
type Repository interface {
	List(ctx context.Context, limit int, offset int) ([]entity.User, error)
	GetByID(ctx context.Context, id int) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int, error)
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

func (s service) List(ctx context.Context, limit int, offset int) ([]entity.User, error) {
	return s.repository.List(ctx, limit, offset)
}

func (s service) GetByID(ctx context.Context, id int) (entity.User, error) {
	return s.repository.GetByID(ctx, id)
}

func (s service) Create(ctx context.Context, request createRequest) (entity.User, error) {
	passwordHash, err := security.GenerateHashFromPassword(request.Password)
	if err != nil {
		return entity.User{}, err
	}
	user := entity.User{
		Login:        request.Login,
		Email:        request.Email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	result, err := s.repository.Create(ctx, user)
	return result, err
}

func (s service) Update(ctx context.Context, id int, request updateRequest) (entity.User, error) {
	user, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return user, err
	}
	user.UpdatedAt = time.Now()

	if err := s.repository.Update(ctx, user); err != nil {
		return user, err
	}
	return user, nil
}

func (s service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s service) Count(context context.Context) (int, error) {
	return s.repository.Count(context)
}
