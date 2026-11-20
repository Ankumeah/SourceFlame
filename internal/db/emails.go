package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ankumeah/SourceFlame/internal/errs"
)

// Add a new email
//
// May return the following errors:
//   - `errs.ErrInvalidRepo`
//   - Errors by the underlying DB
func (db *DBController) AddEmail(ctx context.Context, repo Repo, email Email) error {
	_, err := db.db.ExecContext(
		ctx, `
      INSERT INTO emails(
        user,
        verified,
        email,
        repo,
        repo_owner,
        data
      )
      VALUES (?, ?, ?, ?, ?, ?);
    `,
		email.Sender.Name,
		email.Sender.Verified,
		email.Sender.Email,
		repo.Name,
		repo.Owner,
		email.Data,
	)

	if isForeignKeyViolation(err) {
		return errs.ErrInvalidRepo
	}

	return err
}

func (db *DBController) GetAllEmails(ctx context.Context, repo Repo) ([]Email, error) {
	rows, err := db.db.QueryContext(
		ctx, `
      SELECT
        user,
        verified,
        email,
        data
      FROM emails
      WHERE
        repo = ? AND
        repo_owner = ?;
    `,
		repo.Name,
		repo.Owner,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.ErrInvalidRepo
	} else if err != nil {
		return nil, err
	}

	emails := []Email{}
	for rows.Next() {
		email := Email{}
		if err := rows.Scan(
			&email.Sender.Name,
			&email.Sender.Verified,
			&email.Sender.Email,
			&email.Data,
		); err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}

	return emails, nil
}
