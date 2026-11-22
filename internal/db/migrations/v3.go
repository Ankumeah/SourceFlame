package migrations

// Removes the `email` field from the `users` table
var v3 = migration{
	version: 3,
	queries: []string{
		`
      ALTER TABLE users
      DROP COLUMN email;
    `,
		`
      DROP INDEX users_email_unique;
    `,
	},
}

func init() {
	migrations = append(migrations, v3)
}
