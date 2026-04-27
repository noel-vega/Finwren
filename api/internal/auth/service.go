package auth

import (
	"context"
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

func NewService(userService *user.Service, domain, jwtSecret string) *Service {
	return &Service{
		userService,
		domain,
		jwtSecret,
	}
}

func (s *Service) SignIn() {}

type SignUpParams struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

func (s *Service) SignUp(ctx context.Context, params SignUpParams) (user.UserNoPassword, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)
	if err != nil {
		return user.UserNoPassword{}, err
	}

	return s.userService.CreateUser(ctx, user.CreateUserParams{
		Email:     params.Email,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Password:  string(hashedPassword),
	})
}

func (s *Service) createAccessToken(userID int) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userID),
		Issuer:    s.domain,
		Audience:  jwt.ClaimStrings{s.domain},
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
