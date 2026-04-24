package pgerr

import (
	"errors"

	"github.com/jackc/pgerrcode" // optional: named SQLSTATE constants
	"github.com/jackc/pgx/v5/pgconn"
)

func IsUniqueViolation(err error, constraint string) bool {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return false
	}
	if pgErr.Code != pgerrcode.UniqueViolation { // "23505"
		return false
	}
	return constraint == "" || pgErr.ConstraintName == constraint
}
