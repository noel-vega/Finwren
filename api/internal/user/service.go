package user

import (
	"context"
	"log/slog"
	"strings"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository,
	}
}

func (s *Service) CreateUser(ctx context.Context, params CreateUserParams) (UserNoPassword, error) {
	params.Email = strings.ToLower(strings.TrimSpace(params.Email))

	u, err := s.repository.CreateUser(ctx, params)
	if err != nil {
		return UserNoPassword{}, err
	}

	slog.InfoContext(ctx, "user created",
		slog.Int("user_id", int(u.ID)),
	)

	return UserNoPassword{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Avatar:    u.Avatar,
		CreatedAt: u.CreatedAt,
	}, nil
}
