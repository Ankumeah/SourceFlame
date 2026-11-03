package migrations

import (
	"context"
	"database/sql"
	"strings"
)

type migration struct {
	version uint64
	queries []string
}

func (m migration) apply(
	ctx context.Context,
	tx *sql.Tx,
) error {
	for _, query := range m.queries {
		if strings.TrimSpace(query) == "" {
			continue
		}

		if _, err := tx.ExecContext(ctx, query); err != nil {
			return err
		}
	}

	var err error

	if m.version == 0 {
		_, err = tx.ExecContext(
			ctx, `
        INSERT INTO migrations (version, applied_at)
        VALUES (?, unixepoch())
        ON CONFLICT(version) DO NOTHING;
      `, m.version,
		)
	} else {
		_, err = tx.ExecContext(
			ctx, `
        INSERT INTO migrations (version, applied_at)
        VALUES (?, unixepoch());
    `, m.version,
		)
	}

	return err
}
