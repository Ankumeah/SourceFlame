package migrations

// Adds the `email` field to the `users` table
var v2 = migration{
	version: 2,
	queries: []string{
		`
      ALTER TABLE users
      ADD COLUMN email TEXT;
    `,
		`
      CREATE UNIQUE INDEX users_email_unique
      ON users(email);
    `,
	},
}

func init() {
	migrations = append(migrations, v2)
}
