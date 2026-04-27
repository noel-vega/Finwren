package auth

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type EmailTokensRepository struct {
	db *sqlx.DB
}

func NewEmailTokensRepository(db *sqlx.DB) *EmailTokensRepository {
	return &EmailTokensRepository{
		db,
	}
}

type TokenPurpose string

const (
	TokenPurposeVerifyEmail   TokenPurpose = "verify_email"
	TokenPurposePasswordReset TokenPurpose = "password_reset"
)

type CreateTokenParams struct {
	UserID    int32        `db:"user_id"`
	TokenHash []byte       `db:"token_hash"`
	Purpose   TokenPurpose `db:"purpose"`
	ExpiresAt time.Time    `db:"expires_at"`
}

func (r *EmailTokensRepository) Create(ctx context.Context, params CreateTokenParams) error {
	query := `
	  INSERT INTO email_tokens (user_id, token_hash, purpose, expires_at)
	  VALUES(:user_id, :token_hash, :purpose, :expires_at)
	`

	_, err := r.db.NamedExecContext(ctx, query, params)

	return err
}
