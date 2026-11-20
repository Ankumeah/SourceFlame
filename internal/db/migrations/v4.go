package migrations

// Adds the `emails` table
var v4 = migration{
	version: 4,
	queries: []string{
		`
      CREATE TABLE emails (
        id          INTEGER   PRIMARY KEY,
        data        TEXT      NOT NULL,
        user        TEXT                    REFERENCES users(name),
        verified    BOOLEAN   NOT NULL      DEFAULT 0,
        user_email  TEXT      NOT NULL      REFERENCES users(email),
        repo        TEXT      NOT NULL      REFERENCES repos(name),
        repo_owner  TEXT      NOT NULL      REFERENCES repos(owner)
      );
    `,
	},
}

func init() {
	migrations = append(migrations, v4)
}
