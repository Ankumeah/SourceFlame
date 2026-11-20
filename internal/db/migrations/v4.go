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
        email       TEXT      NOT NULL,
        repo        TEXT      NOT NULL,
        repo_owner  TEXT      NOT NULL,

        FOREIGN KEY (repo_owner, repo) REFERENCES repos(owner, name)
      );
    `,
	},
}

func init() {
	migrations = append(migrations, v4)
}
