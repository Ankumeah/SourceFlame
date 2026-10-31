package db

import (
	"context"

	"github.com/Ankumeah/SourceFlame/internal/errs"
)

// Add a new repo
//
// May return the following errors:
//   - `errs.ErrInvalidUser`
//   - `errs.ErrExistRepo`
//   - Errors by the underlying DB
func (db *DBController) AddRepo(ctx context.Context, repo Repo) error {
	info, err := db.db.ExecContext(
		ctx, `
      INSERT INTO repos(name, owner)
      SELECT ?, ?;
    `, repo.Name, repo.Owner,
	)

	if isUniqueViolation(err) {
		return errs.ErrExistsRepo
	}

	affected, err := info.RowsAffected()
	if err != nil {
		return nil
	} else if affected < 1 {
		return errs.ErrInvalidUser
	}

	return err
}

// Delete a repo
//
// May return the following errors:
//   - `errs.ErrInvalidRepo`
//   - Errors by the underlying DB
func (db *DBController) DeleteRepo(ctx context.Context, repo Repo) error {
	info, err := db.db.ExecContext(
		ctx, `
      DELETE FROM repos
      WHERE
        owner = ? AND
        name = ?;
    `, repo.Owner, repo.Name,
	)

	if err != nil {
		return err
	}

	if affected, err := info.RowsAffected(); err != nil {
		return err
	} else if affected < 1 {
		return errs.ErrInvalidRepo
	}

	return nil
}

func (db *DBController) GetUserRepos(ctx context.Context, user string) ([]Repo, error) {
	rows, err := db.db.QueryContext(
		ctx, `
      SELECT name
      FROM repos
      WHERE owner = ?;
    `, user,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	repos := []Repo{}
	for rows.Next() {
		r := Repo{Owner: user}

		if err := rows.Scan(&r.Name); err != nil {
			return nil, err
		}

		repos = append(repos, r)
	}

	return repos, nil
}
