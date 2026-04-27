package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/noel-vega/finances/api/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userService *user.Service
	domain      string
	jwtSecret   string
}

func NewService(userService *user.Service, domain, jwtSecret string) (*Service, error) {
	service := &Service{
		userService,
		domain,
		jwtSecret,
	}

	if _, err := service.createAccessToken(0); err != nil {
		return nil, fmt.Errorf("jwt signing smoke failed: %w", err)
	}
	return service, nil
}

func (s *Service) SignIn() {}

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

	return s.userService.CreateUser(ctx, user.CreateUserParams{
		Email:     params.Email,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Password:  string(hash),
	})
}

type Token struct {
	Value    string
	Duration time.Duration
}

func (s *Service) createToken(userID int32, duration time.Duration) (Token, error) {
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
		return Token{}, err
	}
	return Token{
		Value:    tokenStr,
		Duration: duration,
	}, nil
}

func (s *Service) createAccessToken(userID int32) (Token, error) {
	return s.createToken(userID, time.Minute*15)
}

func (s *Service) createRefreshToken(userID int32) (Token, error) {
	return s.createToken(userID, time.Hour*24*7)
}
