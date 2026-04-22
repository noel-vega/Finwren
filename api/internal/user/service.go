package user

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		repository,
	}
}

func (s *Service) CreateUser(params CreateUserParams) (UserNoPassword, error) {
	user, err := s.repository.CreateUser(params)
	if err != nil {
		return UserNoPassword{}, err
	}

	return UserNoPassword{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
	}, nil
}
