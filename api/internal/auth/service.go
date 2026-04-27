package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/noel-vega/finances/api/internal/email"
	"github.com/noel-vega/finances/api/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthSender interface {
	SendVerifyEmail(context.Context, string, email.VerifyEmailData) error
}

type Service struct {
	webBaseURL string
	user       *user.Service
	email      AuthSender
	domain     string
	jwtSecret  string
	tokens     *EmailTokensRepository
}

func NewService(db *sqlx.DB, userService *user.Service, email AuthSender, domain, webBaseURL, jwtSecret string) (*Service, error) {
	tokens := NewEmailTokensRepository(db)
	service := &Service{
		user:       userService,
		email:      email,
		domain:     domain,
		webBaseURL: webBaseURL,
		jwtSecret:  jwtSecret,
		tokens:     tokens,
	}

	if _, err := service.createAccessToken(0); err != nil {
		return nil, fmt.Errorf("jwt signing smoke failed: %w", err)
	}
	return service, nil
}

type SignUpParams struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

func (s *Service) SignUp(ctx context.Context, params SignUpParams) (user.UserNoPassword, error) {
	sum := sha256.Sum256([]byte(params.Password))
	prehashed := hex.EncodeToString(sum[:]) // 64 ASCII bytes
	hash, err := bcrypt.GenerateFromPassword([]byte(prehashed), 12)
	if err != nil {
		return user.UserNoPassword{}, err
	}

	u, err := s.user.CreateUser(ctx, user.CreateUserParams{
		Email:     params.Email,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Password:  string(hash),
	})
	if err != nil {
		return u, err
	}

	raw, tokenHash, err := generateRawToken()
	if err != nil {
		return u, err
	}

	err = s.tokens.Create(ctx, CreateTokenParams{
		UserID:    u.ID,
		TokenHash: tokenHash,
		Purpose:   TokenPurposeVerifyEmail,
		ExpiresAt: time.Now().Add(time.Hour * 8),
	})
	if err != nil {
		return u, err
	}

	verifyLink, err := s.buildVerifyEmailLink(raw)
	if err != nil {
		return u, err
	}

	err = s.email.SendVerifyEmail(ctx, u.Email, email.VerifyEmailData{
		Name:           u.FirstName,
		Link:           verifyLink,
		ExpiresInHours: 8,
	})

	return u, err
}

func (s *Service) buildVerifyEmailLink(rawToken string) (string, error) {
	u, err := url.Parse(s.webBaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid BASE_URL: %w", err)
	}
	u.Path = "/auth/verify"
	q := u.Query()
	q.Set("token", rawToken)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

type JWTToken struct {
	Value    string
	Duration time.Duration
}

func (s *Service) createJWTToken(userID int32, duration time.Duration) (JWTToken, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(int(userID)),
		Issuer:    s.domain,
		Audience:  jwt.ClaimStrings{s.domain},
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return JWTToken{}, err
	}
	return JWTToken{
		Value:    tokenStr,
		Duration: duration,
	}, nil
}

func (s *Service) createAccessToken(userID int32) (JWTToken, error) {
	return s.createJWTToken(userID, time.Minute*15)
}

func (s *Service) createRefreshToken(userID int32) (JWTToken, error) {
	return s.createJWTToken(userID, time.Hour*24*7)
}
