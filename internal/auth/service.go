package auth

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dd3v/snippets.page.backend/internal/entity"
	"github.com/dd3v/snippets.page.backend/pkg/security"
	"github.com/dgrijalva/jwt-go"
)

//Service - ...
type Service interface {
	Login(context context.Context, request AuthRequest) (map[string]string, error)
	Refresh(context context.Context, refreshToken string) (string, error)
}

type userClaims struct {
	ID    string `json:"id"`
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

		accessToken, err := s.generateAccessToken(user)
		if err != nil {
			return nil, err
		}

		refreshToken, err := s.generateRefreshToken()
		if err != nil {
			return nil, err
		}

		refreshTokenClaims := entity.RefreshTokens{
			ID:        primitive.NewObjectID(),
			Token:     refreshToken,
			Exp:       time.Now().Add(time.Hour * 100),
			CreatedAt: time.Now(),
		}

		err = s.repository.SaveRefreshToken(context, user.ID.Hex(), refreshTokenClaims)
		if err != nil {
			return nil, err
		}

		return map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		}, nil
	}

	return nil, errors.New("Invalid login or password")
}

func (s service) Refresh(context context.Context, refreshToken string) (string, error) {

	err := s.repository.ReplaceRefreshToken(context, refreshToken)

	if err != nil {
		return "", err
	}
	// token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(s.jwtSigningKey), nil
	// })
	// if err != nil {
	// 	return "", err
	// }
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	fmt.Println(ok)
	// 	fmt.Println(claims["id"])
	// }
	return "", nil
}

func (s service) generateAccessToken(user entity.User) (string, error) {
	jwtClaims := &userClaims{
		ID:    user.ID.Hex(),
		Login: user.Login,
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

// func (s service) generateTokenPair(user entity.User) (map[string]string, error) {

// 	refreshTokenClaims := entity.RefreshTokens{
// 		ID:        primitive.NewObjectID(),
// 		Token:     refreshToken.String(),
// 		Exp:       time.Now().Add(time.Hour * 100),
// 		CreatedAt: time.Now(),
// 	}
// 	_ = s.repository.SaveRefreshToken(context.TODO(), user.ID.Hex(), refreshTokenClaims)

// }
