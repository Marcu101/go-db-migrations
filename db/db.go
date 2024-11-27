package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Creates a table for tracking DB migrations, if it does not exist
func Init(c *pgx.Conn) error {
	sql := `
	CREATE TABLE IF NOT EXISTS _migrations (
    	version TEXT PRIMARY KEY,
    	applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);`
	_, err := c.Exec(context.Background(), sql)

	return err
}

// Fetches the list of already applied migrations
func GetAppliedMigrations(c *pgx.Conn) ([]string, error) {
	var appliedMigrations []string
	rows, err := c.Query(context.Background(), "SELECT version FROM _migrations ORDER BY applied_at;")
	if err != nil {
		return appliedMigrations, err
	}
	defer rows.Close()

	for rows.Next() {
		var mig string
		err = rows.Scan(&mig)
		if err != nil {
			return appliedMigrations, err
		}
		appliedMigrations = append(appliedMigrations, mig)
	}

	return appliedMigrations, err
}

// Applies a migration
func ApplyMigration(c *pgx.Conn, migration, sql string) error {
	_, err := c.Exec(context.Background(), sql)
	if err != nil {
		return err
	}

	_, err = c.Exec(context.Background(), "INSERT INTO _migrations (version) VALUES ($1);", migration)
	return err
}
