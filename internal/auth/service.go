package auth

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/security"
	"github.com/dgrijalva/jwt-go"
)

//Service - ...
type Service interface {
	Login(context context.Context, request AuthRequest) (map[string]string, error)
	Refresh(context context.Context, refreshToken string) (map[string]string, error)
	Logout(context context.Context, refreshToken string) error
}

type userClaims struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	jwt.StandardClaims
}

type service struct {
	jwtSigningKey string
	repository    Repository
}

//NewService - ...
func NewService(JWTSigningKey string, repository Repository) Service {
	return &service{
		jwtSigningKey: JWTSigningKey,
		repository:    repository,
	}
}

func (s service) Login(context context.Context, request AuthRequest) (map[string]string, error) {
	user, err := s.repository.FindUser(context, request.Login)
	if err != nil {
		return nil, errors.New("Invalid login or password")
	}
	if security.CompareHashAndPassword(user.PasswordHash, request.Password) == true {

		accessToken, err := s.generateAccessToken(user.ID)
		if err != nil {
			return nil, err
		}

		refreshToken, err := s.generateRefreshToken()
		if err != nil {
			return nil, err
		}

		session := entity.Session{
			UserID:       user.ID,
			RefreshToken: refreshToken,
			Exp:          time.Now().Add(time.Minute * 100),
			IP:           "127.0.0.1",
			UserAgent:    "local",
			CreatedAt:    time.Now(),
		}

		if err = s.repository.CreateSession(context, session); err != nil {
			return nil, err
		}

		return map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		}, nil
	}

	return nil, errors.New("Invalid login or password")
}

func (s service) Refresh(context context.Context, refreshToken string) (map[string]string, error) {

	session, err := s.repository.FindSessionByRefreshToken(context, refreshToken)
	if err != nil {
		return nil, err
	}
	if err := s.repository.DeleteSessionByRefreshToken(context, session.RefreshToken); err != nil {
		return nil, err
	}

	if session.Exp.After(time.Now()) == true {
		accessToken, err := s.generateAccessToken(session.UserID)
		if err != nil {
			return nil, err
		}

		refreshToken, err := s.generateRefreshToken()
		if err != nil {
			return nil, err
		}

		session := entity.Session{
			UserID:       session.UserID,
			RefreshToken: refreshToken,
			Exp:          time.Now().Add(time.Minute * 100),
			IP:           "127.0.0.1",
			UserAgent:    "local",
			CreatedAt:    time.Now(),
		}

		if err = s.repository.CreateSession(context, session); err != nil {
			return nil, errors.New("session error expired")
		}

		return map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		}, nil
	}
	return nil, errors.New("session already expired")
}

func (s service) Logout(context context.Context, token string) error {
	return s.repository.DeleteSessionByRefreshToken(context, token)
}

func (s service) generateAccessToken(userID int) (string, error) {
	jwtClaims := &userClaims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims).SignedString([]byte(s.jwtSigningKey))
}

func (s service) generateRefreshToken() (string, error) {
	token, err := uuid.NewUUID()
	return token.String(), err
}
