package user

import (
	"context"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

//Service - ...
type Service interface {
	FindByID(context context.Context, id interface{}) (entity.User, error)
	CreateUser(context context.Context, request CreateRequest) (entity.User, error)
	Count(context context.Context) (int, error)
}

type service struct {
	repository Repository
}

func generatePasswortHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash[:]), err
}

//NewService - ...
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) FindByID(context context.Context, id interface{}) (entity.User, error) {
	var user entity.User

	return user, nil
}

func (s service) CreateUser(context context.Context, request CreateRequest) (entity.User, error) {
	user := entity.User{
		Login: request.Login,
		Email: request.Email,
	}
	err := s.repository.Create(context, user)
	return user, err
}

func (s service) Count(context context.Context) (int, error) {
	return s.repository.Count(context)
}
