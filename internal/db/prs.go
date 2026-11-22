package db

import (
	"context"
	"database/sql"
	"errors"
	"uuid"

	"github.com/Ankumeah/SourceFlame/internal/errs"
)

// Add a new pr
//
// May return the following errors:
//   - `errs.ErrInvalid`
//   - Errors by the underlying DB
func (db *DBController) AddPR(ctx context.Context, pr PR) error {
	_, err := db.db.ExecContext(
		ctx, `
      INSERT INTO prs(
        uuid,
        title,
        body,
        owner,
        repo,
        repo_owner,
        url,
        status
      )
      VALUES(?, ?, ?, ?, ?, ?, ?, 'pending');
    `,
		uuid.NewV7(),
		pr.Title,
		pr.Body,
		pr.Owner,
		pr.Repo.Name,
		pr.Repo.Owner,
		pr.URL,
	)

	if isForeignKeyViolation(err) {
		return errs.ErrInvalid
	}

	return err
}

// Resolve a PR
//
// May return the following errors:
//   - `errs.ErrInvalidPR`
//   - Errors by the underlying DB
func (db *DBController) ResolvePR(ctx context.Context, uuid uuid.UUID, approved bool) error {
	var status string
	if approved {
		status = "approved"
	} else {
		status = "rejected"
	}

	info, err := db.db.ExecContext(
		ctx, `
      UPDATE prs
      SET status = ?
      WHERE uuid = ?;
    `, status, uuid,
	)

	if err != nil {
		return err
	}

	affected, err := info.RowsAffected()
	if err != nil {
		return err
	} else if affected < 1 {
		return errs.ErrInvalidPR
	}

	return nil
}

// Get a repo's PRs
//
// May return the following errors:
//   - Errors by the underlying DB
func (db *DBController) GetPRs(ctx context.Context, repo Repo, page uint64) ([]PR, error) {
	const limit = 10

	rows, err := db.db.QueryContext(
		ctx, `
      SELECT
        uuid,
        title,
        body,
        owner,
        url,
        status
      FROM prs
      WHERE
        repo = ? AND
        repo_owner = ?
      LIMIT ?
      OFFSET ?;
    `,
		repo.Name,
		repo.Owner,
		limit,
		limit*(page-1),
	)

	if errors.Is(err, sql.ErrNoRows) {
		return []PR{}, nil
	} else if err != nil {
		return nil, err
	}

	prs := []PR{}
	for rows.Next() {
		pr := PR{
			Repo: repo,
		}

		if err := rows.Scan(
			&pr.UUID,
			&pr.Title,
			&pr.Body,
			&pr.Owner,
			&pr.URL,
			&pr.Status,
		); err != nil {
			return nil, err
		}

		prs = append(prs, pr)
	}

	return prs, nil
}
