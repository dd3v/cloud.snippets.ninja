package auth

import (
	"context"
	"errors"
	"time"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/security"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

//Service - ...
type Service interface {
	Login(ctx context.Context, request loginRequest) (entity.TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (entity.TokenPair, error)
	Logout(ctx context.Context, refreshToken string) error
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
	return service{
		jwtSigningKey: JWTSigningKey,
		repository:    repository,
	}
}

func (s service) Login(ctx context.Context, request loginRequest) (entity.TokenPair, error) {
	user, err := s.repository.FindUserByLoginOrEmail(ctx, request.Login)
	if err != nil {
		return entity.TokenPair{}, errors.New("Invalid login or password")
	}
	if security.CompareHashAndPassword(user.PasswordHash, request.Password) == true {
		accessToken, err := s.generateAccessToken(user.ID)
		if err != nil {
			return entity.TokenPair{}, err
		}
		refreshToken, err := s.generateRefreshToken()
		if err != nil {
			return entity.TokenPair{}, err
		}
		session := entity.Session{
			UserID:       user.ID,
			RefreshToken: refreshToken,
			Exp:          time.Now().Add(time.Minute * 100),
			IP:           "127.0.0.1",
			UserAgent:    "local",
			CreatedAt:    time.Now(),
		}
		if err = s.repository.CreateSession(ctx, session); err != nil {
			return entity.TokenPair{}, err
		}
		return entity.TokenPair{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}
	return entity.TokenPair{}, errors.New("Invalid login or password")
}

func (s service) Refresh(ctx context.Context, refreshToken string) (entity.TokenPair, error) {
	session, err := s.repository.FindSessionByRefreshToken(ctx, refreshToken)
	if err != nil {
		return entity.TokenPair{}, err
	}
	if err := s.repository.DeleteSessionByRefreshToken(ctx, session.RefreshToken); err != nil {
		return entity.TokenPair{}, err
	}
	if session.Exp.After(time.Now()) == true {
		accessToken, err := s.generateAccessToken(session.UserID)
		if err != nil {
			return entity.TokenPair{}, err
		}
		refreshToken, err := s.generateRefreshToken()
		if err != nil {
			return entity.TokenPair{}, err
		}
		session := entity.Session{
			UserID:       session.UserID,
			RefreshToken: refreshToken,
			Exp:          time.Now().Add(time.Minute * 100),
			IP:           "127.0.0.1",
			UserAgent:    "local",
			CreatedAt:    time.Now(),
		}
		if err = s.repository.CreateSession(ctx, session); err != nil {
			return entity.TokenPair{}, errors.New("session error expired")
		}
		return entity.TokenPair{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}
	return entity.TokenPair{}, errors.New("session already expired")
}

func (s service) Logout(ctx context.Context, token string) error {
	return s.repository.DeleteSessionByRefreshToken(ctx, token)
}

func (s service) generateAccessToken(userID int) (string, error) {
	jwtClaims := &userClaims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2315).Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims).SignedString([]byte(s.jwtSigningKey))
}

func (s service) generateRefreshToken() (string, error) {
	token, err := uuid.NewUUID()
	return token.String(), err
}
