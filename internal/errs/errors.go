package errs

import (
	"errors"
	"fmt"
)

var ErrExists = errors.New("Exists")
var ErrExistsUser = fmt.Errorf("%w: User exist", ErrExists)
var ErrExistsRepo = fmt.Errorf("%w: Repo exist", ErrExists)
var ErrExistsEmail = fmt.Errorf("%w: Email exist", ErrExists)

var ErrInvalid = errors.New("Invalid")
var ErrInvalidUser = fmt.Errorf("%w: Invalid user", ErrInvalid)
var ErrInvalidRepo = fmt.Errorf("%w: Invalid repo", ErrInvalid)
var ErrInvalidEmail = fmt.Errorf("%w: Invalid email", ErrInvalid)
var ErrInvalidBlob = fmt.Errorf("%w: Invalid blob", ErrInvalid)
var ErrInvalidRef = fmt.Errorf("%w: Invalid ref", ErrInvalid)

var ErrNotExist = errors.New("%w: Not exists")
var ErrNotExistCommits = fmt.Errorf("%w: No commits exist", ErrNotExist)
