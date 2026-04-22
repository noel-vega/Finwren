package auth

import "github.com/noel-vega/finances/api/internal/user"

type Service struct {
	userService user.Service
}

func NewService(userService user.Service) *Service {
	return &Service{
		userService,
	}
}

func (s *Service) SignIn() {}

type SignUpParams struct {
	Email           string
	FirstName       string
	LastName        string
	Password        string
	ConfirmPassword string
}

func (p *SignUpParams) Validate() map[string]validationIssue {
	issues := map[string]validationIssue{}
}

type validationIssue struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func (s *Service) SignUp(params SignUpParams) (user.UserNoPassword, error) {
	if params.Password != params.ConfirmPassword {
		return user.UserNoPassword{}, ErrPasswordConfirmMismatch
	}

	return s.userService.CreateUser(user.CreateUserParams{
		Email:     params.Email,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Password:  params.Password,
	})
}
