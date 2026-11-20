package db

import (
	"context"

	"github.com/Ankumeah/SourceFlame/internal/errs"
)

// Add a new email
//
// May return the following errors:
//   - `errs.ErrInvalidRepo`
//   - Errors by the underlying DB
func (db *DBController) AddEmail(ctx context.Context, sender EmailSender, repo Repo, data []byte) error {
	_, err := db.db.ExecContext(
		ctx, `
      INSERT INTO emails(
        user,
        verified,
        user_email,
        repo,
        repo_owner,
        data
      )
      VALUES (?, ?, ?, ?, ?, ?);
    `,
		sender.Name,
		sender.Verified,
		sender.Email,
		repo.Name,
		repo.Owner,
		data,
	)

	if isForeignKeyViolation(err) {
		return errs.ErrInvalidRepo
	}

	return err
}
