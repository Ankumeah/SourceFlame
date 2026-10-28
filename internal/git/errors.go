package git

import (
	"errors"
	"fmt"
)

var ErrExists = errors.New("Exists")
var ErrExistsRepo = fmt.Errorf("%w: Repo exists", ErrExists)

var ErrNotFound = errors.New("Not found")
var ErrNotFoundRepo = fmt.Errorf("%w: Repo not found", ErrNotFound)
