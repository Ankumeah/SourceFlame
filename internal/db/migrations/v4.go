package migrations

// Adds the `prs` table
var v4 = migration{
	version: 4,
	queries: []string{
		`
      CREATE TABLE prs (
        uuid        TEXT  PRIMARY KEY,
        title       TEXT  NOT NULL,
        body        TEXT  NOT NULL,
        owner       TEXT  NOT NULL      REFERENCES users(name),
        repo        TEXT  NOT NULL,
        repo_owner  TEXT  NOT NULL,
        url         TEXT  NOT NULL,
        status      TEXT  NOT NULL      CHECK (status IN ('approved', 'rejected', 'pending'))

        FOREIGN KEY (repo, repo_owner) REFERENCES parent(name, owner)
      );
    `,
	},
}

func init() {
	migrations = append(migrations, v4)
}
