package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ankumeah/SourceFlame/internal/errs"
	"github.com/Ankumeah/SourceFlame/internal/hash"
)

// Add a new user
//
// May return the following errors:
//   - `errs.ErrExistUser`
//   - Errors by the underlying DB
func (db *DBController) AddUser(ctx context.Context, user string, hash hash.Hash) error {
	_, err := db.db.ExecContext(
		ctx, `
      INSERT INTO users(name, password, salt)
      VALUES (?, ?, ?);
    `, user, hash.Hash, hash.Salt,
	)

	if isUniqueViolation(err) {
		return errs.ErrExistsUser
	}

	return err
}

// Delete a user
//
// May return the following errors:
//   - `errs.ErrInvalidUser`
//   - Errors by the underlying DB
func (db *DBController) DeleteUser(ctx context.Context, user string) error {
	info, err := db.db.ExecContext(
		ctx, `
      DELETE FROM users
      WHERE name = ?;
    `, user,
	)

	if err != nil {
		return err
	}

	if affected, err := info.RowsAffected(); err != nil {
		return err
	} else if affected < 1 {
		return errs.ErrInvalidUser
	}

	return nil
}

// Get a user's password hash
//
// May return the following errors:
//   - `errs.ErrInvalidUser`
//   - Errors by the underlying DB
func (db *DBController) GetPassHash(ctx context.Context, user string) (hash.Hash, error) {
	var hash hash.Hash

	err := db.db.QueryRowContext(
		ctx, `
      SELECT password, salt
      FROM users
      WHERE name = ?;
    `, user,
	).Scan(&hash.Hash, &hash.Salt)

	if errors.Is(err, sql.ErrNoRows) {
		err = errs.ErrInvalidUser
	}

	return hash, err
}
