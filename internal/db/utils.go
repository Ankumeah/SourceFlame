package db

import (
	"errors"

	"modernc.org/sqlite"
	lib "modernc.org/sqlite/lib"
)

func isUniqueViolation(err error) bool {
	if sqliteErr, ok := errors.AsType[*sqlite.Error](err); ok {
		return sqliteErr.Code() == lib.SQLITE_CONSTRAINT_UNIQUE
	}

	return false
}
