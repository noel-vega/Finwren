package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/noel-vega/finances/api/internal/pgerr"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}

type CreateUserParams struct {
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Password  string `db:"password"`
}

func (r *Repository) CreateUser(ctx context.Context, params CreateUserParams) (User, error) {
	query := `
	INSERT INTO users (email, first_name, last_name, password)
	VALUES (:email, :first_name, :last_name, :password)
	RETURNING *
	`

	query, args, err := sqlx.Named(query, params)
	if err != nil {
		return User{}, err
	}

	query = r.db.Rebind(query)

	u := User{}
	err = r.db.GetContext(ctx, &u, query, args...)
	if err != nil {
		if pgerr.IsUniqueViolation(err, "users_email_key") {
			return u, ErrEmailExists
		}
		return u, err
	}

	return u, nil
}
