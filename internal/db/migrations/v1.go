package migrations

// Adds the `users` and `repos` table
var v1 = migration{
	version: 1,
	queries: []string{
		`
      CREATE TABLE users (
        id        INTEGER   PRIMARY KEY,
        name      TEXT      NOT NULL      UNIQUE,
        password  TEXT      NOT NULL,
        salt      TEXT      NOT NULL
      );
    `,
		`
      CREATE TABLE repos (
        id    INTEGER   PRIMARY KEY,
        name  TEXT      NOT NULL,
        owner TEXT      NOT NULL      REFERENCES users(name)  ON DELETE CASCADE,

        UNIQUE (owner, name)
      );
    `,
	},
}

func init() {
	migrations = append(migrations, v1)
}
