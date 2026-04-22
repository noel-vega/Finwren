package user

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{db}
}

type CreateUserParams struct {
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Avatar    string `db:"avatar"`
	Password  string `db:"password"`
}

func (r *Repository) CreateUser(params CreateUserParams) (User, error) {
	query := `
	INSERT INTO users (email, first_name, last_name, avatar, password)
	VALUES (:email, :first_name, :last_name, :avatar, :password)
	RETURNING *
	`

	query, args, err := sqlx.Named(query, params)
	if err != nil {
		return User{}, err
	}

	query = r.db.Rebind(query)

	user := User{}
	err = r.db.Get(&user, query, args...)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
