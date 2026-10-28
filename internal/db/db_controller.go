package db

import (
	"context"
	"database/sql"
	"uuid"

	"github.com/Ankumeah/SourceFlame/internal/hash"
	_ "modernc.org/sqlite"
)

type DBController struct{ db *sql.DB }

func GetDBController(url string) (*DBController, error) {
	db, err := sql.Open("sqlite", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DBController{db}, nil
}

// This exposes the raw underlying DB object.
// This is only to be used to execute migrations
func (db *DBController) DB() *sql.DB {
	return db.db
}

// Add a new user
//
// May return the following errors:
//   - `ErrExistUser`
//   - Errors by the underlying DB
func (db *DBController) AddUser(ctx context.Context, user User, hash hash.Hash) error {
	_, err := db.db.ExecContext(
		ctx, `
      INSERT INTO users(uuid, name, password, salt)
      VALUES (?, ?, ?, ?);
    `, uuid.New(), user.Name, hash.Hash, hash.Salt,
	)

	if isUniqueViolation(err) {
		return ErrExistsUser
	}

	return err
}

// Delete a user
//
// May return the following errors:
//   - `ErrInvalidUser`
//   - Errors by the underlying DB
func (db *DBController) DeleteUser(ctx context.Context, name string) error {
	info, err := db.db.ExecContext(
		ctx, `
      DELETE FROM users
      WHERE name = ?;
    `, name,
	)

	if err != nil {
		return err
	}

	if affected, err := info.RowsAffected(); err != nil {
		return err
	} else if affected < 1 {
		return ErrInvalidUser
	}

	return nil
}

// Add a new repo
//
// May return the following errors:
//   - `ErrExistRepo`
//   - Errors by the underlying DB
func (db *DBController) AddRepo(ctx context.Context, repo Repo) error {
	_, err := db.db.ExecContext(
		ctx, `
      INSERT INTO repos(uuid, name, owner_uuid)
      SELECT ?, ?, u.name
      FROM users AS u
      WHERE u.name = ?;
    `, uuid.New(), repo.Name, repo.Owner,
	)

	if isUniqueViolation(err) {
		return ErrExistsRepo
	}

	return err
}

// Delete a repo
//
// May return the following errors:
//   - `ErrInvalidRepo`
//   - Errors by the underlying DB
func (db *DBController) DeleteRepo(ctx context.Context, repo Repo) error {
	info, err := db.db.ExecContext(
		ctx, `
      DELETE FROM repos AS r
      JOIN users AS u ON u.uuid = r.owner_uuid
      WHERE
        u.name = ? AND
        r.name = ?;
    `, repo.Owner, repo.Name,
	)

	if err != nil {
		return err
	}

	if affected, err := info.RowsAffected(); err != nil {
		return err
	} else if affected < 1 {
		return ErrInvalidRepo
	}

	return nil
}

func (db *DBController) GetUserRepos(ctx context.Context, user User) ([]Repo, error) {
	rows, err := db.db.QueryContext(
		ctx, `
      SELECT u.name, r.name
      FROM repos AS r
      JOIN users AS u on u.uuid = r.owner_uuid
      WHERE u.name = ?;
    `, user.Name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	repos := []Repo{}
	for rows.Next() {
		r := Repo{}

		if err := rows.Scan(
			&r.Owner, &r.Name,
		); err != nil {
			return nil, err
		}

		repos = append(repos, r)
	}

	return repos, nil
}
