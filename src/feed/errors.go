package feed

import (
	"errors"
)

var (
	ErrUserDoesNotExist   = errors.New("user does not exist")
	ErrSomethingWentWrong = errors.New("something went wrong")
)
