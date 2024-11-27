package main

import (
	"context"
	"errors"
	"go-db-migrations/db"
	"go-db-migrations/migration"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	log.Exit(0)
}

func run() error {
	// Get required DB url env var
	dbStr := os.Getenv("DATABASE_URL")
	if dbStr == "" {
		return errors.New("env var DATABASE_URL required, please provide a valid database connection string. Exiting")
	}

	// Initiate database connection
	conn, err := pgx.Connect(context.Background(), dbStr)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	// Attempt to create _migrations table
	if err := db.Init(conn); err != nil {
		return err
	}

	// Fetch applied migrations from DB
	appliedMigs, err := db.GetAppliedMigrations(conn)
	if err != nil {
		return err
	}

	// fetch new migrations from the ./migrations directory
	newMigs, err := migration.GetNew(appliedMigs)
	if err != nil {
		return err
	}

	// Apply migrations one by one
	for _, mig := range newMigs {
		sql, err := os.ReadFile(mig + "/up.sql")
		if err != nil {
			return err
		}
		if err := db.ApplyMigration(conn, strings.TrimPrefix(mig, "migrations/"), string(sql)); err != nil {
			return err
		}
	}

	return nil

}
