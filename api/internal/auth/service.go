package auth

import (
	"context"

	"github.com/noel-vega/finances/api/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userService *user.Service
}

func NewService(userService *user.Service) *Service {
	return &Service{
		userService,
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
