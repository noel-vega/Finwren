package user

import "time"

type User struct {
	ID        int32
	Email     string
	FirstName string
	LastName  string
	Avatar    *string
	Password  string
	CreatedAt time.Time
}

type UserNoPassword struct {
	ID        int32
	Email     string
	FirstName string
	LastName  string
	Avatar    *string
	CreatedAt time.Time
}
