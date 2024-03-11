package database

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type NotFoundError struct {
	Entity string
}
