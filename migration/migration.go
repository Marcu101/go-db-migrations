package migration

import (
	"io/fs"
	"path/filepath"
)

const (
	defaultMigrationsDir = "./migrations"
)

// Reads the ./migrations directory and fetches the list of new migrations to be applied
func GetNew(appliedMigs []string) ([]string, error) {

	migList := make([]string, 0)
	err := filepath.WalkDir(defaultMigrationsDir, func(path string, f fs.DirEntry, err error) error {

		if f.IsDir() && f.Name() != "migrations" && !contains(appliedMigs, f.Name()) {
			migList = append(migList, path)
		}
		return err
	})

	return migList, err
}

// Checks if a string value is present in slice
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
