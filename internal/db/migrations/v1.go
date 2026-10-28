package migrations

import (
	"context"
	"database/sql"
	"time"
)

// Adds the `users` and `repos` table
type v1 struct{}

func init() {
	migrations = append(migrations, v1{})
}

func (v1) Version() uint64 {
	return 1
}

func (v1) Apply(
	ctx context.Context,
	tx *sql.Tx,
) error {
	if _, err := tx.ExecContext(
		ctx, `
      CREATE TABLE users (
        id        INTEGER   PRIMARY KEY,
        uuid      TEXT      NOT NULL      UNIQUE,
        name      TEXT      NOT NULL      UNIQUE,
        password  TEXT      NOT NULL,
        salt      TEXT      NOT NULL
      );
    `,
	); err != nil {
		return err
	}

	if _, err := tx.ExecContext(
		ctx, `
      CREATE TABLE repos (
        id          INTEGER   PRIMARY KEY,
        uuid        TEXT      NOT NULL      UNIQUE,
        name        TEXT      NOT NULL,
        owner_uuid  TEXT      NOT NULL      REFERENCES users(uuid),

        UNIQUE (name, owner_uuid)
      );
    `,
	); err != nil {
		return err
	}

	if _, err := tx.ExecContext(
		ctx, `
      INSERT INTO migrations (version, applied_at)
      VALUES (1, ?);
    `, time.Now().Unix(),
	); err != nil {
		return err
	}

	return tx.Commit()
}
