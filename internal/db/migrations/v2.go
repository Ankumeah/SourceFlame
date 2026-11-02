package migrations

import (
	"context"
	"database/sql"
	"time"
)

// Adds the `email` field to the `users` table
type v2 struct{}

func init() {
	migrations = append(migrations, v2{})
}

func (v2) Version() uint64 {
	return 2
}

func (v2) Apply(
	ctx context.Context,
	tx *sql.Tx,
) error {
	if _, err := tx.ExecContext(
		ctx, `
      ALTER TABLE users
      ADD COLUMN email TEXT;
    `,
	); err != nil {
		return err
	}

	if _, err := tx.ExecContext(
		ctx, `
      CREATE UNIQUE INDEX users_email_unique
      ON users(email);
    `,
	); err != nil {
		return err
	}

	if _, err := tx.ExecContext(
		ctx, `
      INSERT INTO migrations (version, applied_at)
      VALUES (2, ?);
    `, time.Now().Unix(),
	); err != nil {
		return err
	}

	return tx.Commit()
}
