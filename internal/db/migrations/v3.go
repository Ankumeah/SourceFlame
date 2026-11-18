package migrations

// Adds the `pub_key` field to the `users` table
var v3 = migration{
	version: 3,
	queries: []string{
		`
      ALTER TABLE users
      ADD COLUMN pub_key TEXT;
    `,
	},
}

func init() {
	migrations = append(migrations, v3)
}
