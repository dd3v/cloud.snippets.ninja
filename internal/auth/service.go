package auth

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dd3v/snippets.ninja/pkg/log"

	"time"

	"github.com/dd3v/snippets.ninja/internal/entity"
	"github.com/dd3v/snippets.ninja/internal/errors"
	"github.com/dd3v/snippets.ninja/pkg/security"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	authErr            = errors.Unauthorized("auth: invalid login or password")
	createSessionErr   = errors.InternalServerError("auth: session create error")
	expiredSessionErr  = errors.Forbidden("auth: session already expired")
	sessionNotFoundErr = errors.Forbidden("auth: accesslog denied")
)

//Service - ...
type Service interface {
	Login(ctx context.Context, credentials authCredentials) (entity.TokenPair, error)
	Refresh(ctx context.Context, credentials refreshCredentials) (entity.TokenPair, error)
	Logout(ctx context.Context, refreshToken string) error
}

//Repository - ...
type Repository interface {
	GetUserByLoginOrEmail(ctx context.Context, value string) (entity.User, error)
	CreateSession(ctx context.Context, session entity.Session) error
	GetSessionByRefreshToken(ctx context.Context, refreshToken string) (entity.Session, error)
	DeleteSessionByRefreshToken(ctx context.Context, refreshToken string) error
	DeleteSessionByUserIDAndUserAgent(ctx context.Context, userID int, userAgent string) error
	DeleteSessionByUserID(ctx context.Context, userID int) (int64, error)
}

type authCredentials struct {
	User      string
	Password  string
	UserAgent string
	IP        string
}

type refreshCredentials struct {
	RefreshToken string
	UserAgent    string
	IP           string
}

type userClaims struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	jwt.StandardClaims
}

type service struct {
	jwtSigningKey string
	repository    Repository
	logger        log.Logger
}

//NewService - ...
func NewService(JWTSigningKey string, repository Repository, logger log.Logger) Service {
	return service{
		jwtSigningKey: JWTSigningKey,
		repository:    repository,
		logger:        logger,
	}
}

//Login
//1. Try to get user by login or email
//2. Check if it exists
//3. Compare password hash
//4.Generate token pair
//5. Remove useless sessions from db by user id and user-agent
//6.Upsert new fresh session
func (s service) Login(ctx context.Context, credentials authCredentials) (entity.TokenPair, error) {
	user, err := s.repository.GetUserByLoginOrEmail(ctx, credentials.User)
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.With(ctx).Info("Security alert. User not found")
			return entity.TokenPair{}, authErr

		} else {
			return entity.TokenPair{}, err
		}
		fmt.Println("FUUUK")

	}
	if security.CompareHashAndPassword(user.Password, credentials.Password) == true {
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
			IP:           credentials.IP,
			UserAgent:    credentials.UserAgent,
			CreatedAt:    time.Now(),
		}

		err = s.repository.DeleteSessionByUserIDAndUserAgent(ctx, user.ID, credentials.UserAgent)
		if err != nil {
			return entity.TokenPair{}, err
		}
		if err = s.repository.CreateSession(ctx, session); err != nil {
			return entity.TokenPair{}, err
		}

		return entity.TokenPair{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	} else {
		s.logger.With(ctx).Info("Security alert. Invalid password")
	}
	return entity.TokenPair{}, authErr
}

func (s service) Refresh(ctx context.Context, credentials refreshCredentials) (entity.TokenPair, error) {
	session, err := s.repository.GetSessionByRefreshToken(ctx, credentials.RefreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.Info("Security alert. Refresh sessions not found")
			return entity.TokenPair{}, sessionNotFoundErr
		}
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
			IP:           credentials.IP,
			UserAgent:    credentials.UserAgent,
			CreatedAt:    time.Now(),
		}
		if err = s.repository.CreateSession(ctx, session); err != nil {
			return entity.TokenPair{}, createSessionErr
		}
		return entity.TokenPair{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	} else {
		//if refresh token is expired just remove all sessions
		if sessionsCount, err := s.repository.DeleteSessionByUserID(ctx, session.UserID); err == nil {
			s.logger.Info("Security alert. Remove all sessions by user id. Total: %d", sessionsCount)
		}
	}
	return entity.TokenPair{}, expiredSessionErr
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
	if err != nil {
		return "", err
	}
	return token.String(), err
}
