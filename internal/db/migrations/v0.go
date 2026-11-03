package migrations

// This is the core migration which tracks
// all other migrations
var v0 = migration{
	version: 0,
	queries: []string{
		`
      CREATE TABLE IF NOT EXISTS migrations (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        version BIGINT UNIQUE NOT NULL,
        applied_at BIGINT NOT NULL
      );
    `,
	},
}
