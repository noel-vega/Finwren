package user

import "time"

type User struct {
	ID        int32     `db:"id"`
	Email     string    `db:"email"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Avatar    *string   `db:"avatar"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserNoPassword struct {
	ID        int32     `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Avatar    *string   `json:"avatar" db:"avatar"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
