package db

import (
	"errors"
	"fmt"
)

var ErrExists = errors.New("Exists")
var ErrExistsUser = fmt.Errorf("%w: User exist", ErrExists)
var ErrExistsRepo = fmt.Errorf("%w: Repo exist", ErrExists)

var ErrInvalid = errors.New("Invalid")
var ErrInvalidUser = fmt.Errorf("%w: Invalid user", ErrInvalid)
var ErrInvalidRepo = fmt.Errorf("%w: Invalid repo", ErrInvalid)
