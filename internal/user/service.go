package user

import (
	"context"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/security"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Service - ...
type Service interface {
	Find(context context.Context, filter map[string]interface{}) ([]entity.User, error)
	FindByID(context context.Context, id string) (entity.User, error)
	Create(context context.Context, request CreateRequest) (entity.User, error)
	Update(context context.Context, id string, request UpdateRequest) (entity.User, error)
	Delete(context context.Context, id string) error
	Count(context context.Context) (int, error)
}

type service struct {
	repository Repository
}

//NewService - ...
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) Find(context context.Context, filter map[string]interface{}) ([]entity.User, error) {
	return s.repository.Find(context, filter)
}

func (s service) FindByID(context context.Context, id string) (entity.User, error) {
	return s.repository.FindByID(context, id)
}

func (s service) Create(context context.Context, request CreateRequest) (entity.User, error) {
	passwordHash, err := security.GenerateHashFromPassword(request.Password)
	if err != nil {
		return entity.User{}, err
	}
	user := entity.User{
		ID:            primitive.NewObjectID(),
		Login:         request.Login,
		Email:         request.Email,
		PasswordHash:  passwordHash,
		RefreshTokens: make([]entity.RefreshTokens, 0),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err = s.repository.Create(context, user)
	return user, err
}

func (s service) Update(context context.Context, id string, request UpdateRequest) (entity.User, error) {
	user, err := s.repository.FindByID(context, id)
	if err != nil {
		return user, err
	}
	user.Website = request.Website
	user.UpdatedAt = time.Now()

	if err := s.repository.Update(context, user); err != nil {
		return user, err
	}
	return user, nil
}

func (s service) Delete(context context.Context, id string) error {
	return s.repository.Delete(context, id)
}

func (s service) Count(context context.Context) (int, error) {
	return s.repository.Count(context)
}
