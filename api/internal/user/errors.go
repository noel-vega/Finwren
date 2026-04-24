package user

import "errors"

var ErrEmailExists = errors.New("User with this email already exists")
