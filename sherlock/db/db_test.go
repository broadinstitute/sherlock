package db

import (
	"github.com/stretchr/testify/assert"
	"path"
	"strings"
	"testing"
)

// TestMigrationFiles doesn't worry about running migrations -- that'll happen
// during connected database tests -- just that every file has an up/down pair.
func TestMigrationFiles(t *testing.T) {
	entries, err := MigrationFiles.ReadDir("migrations")
	assert.NoError(t, err)
	filenames := make([]string, len(entries))
	for index, entry := range entries {
		assert.Falsef(t, entry.IsDir(), "%s is directory", entry.Name())
		filenames[index] = entry.Name()
	}
	t.Run("migration file naming", func(t *testing.T) {
		for _, filename := range filenames {
			if strings.HasSuffix(filename, ".down.sql") {
				assert.Contains(t, filenames, strings.TrimSuffix(filename, ".down.sql")+".up.sql")
			} else if strings.HasSuffix(filename, ".up.sql") {
				assert.Contains(t, filenames, strings.TrimSuffix(filename, ".up.sql")+".down.sql")
			} else {
				assert.Failf(t, "file lacks .up.sql or .down.sql extension", "filename '%s'", filename)
			}
		}
	})

	// This is kinda a linting thing but it's easier to enforce via test than anything else.
	t.Run("migration files end with newlines", func(t *testing.T) {
		for _, filename := range filenames {
			data, err := MigrationFiles.ReadFile(path.Join("migrations", filename))
			assert.NoError(t, err)
			assert.Truef(t, strings.HasSuffix(string(data), "\n"), "migrations/%s should end with a newline", filename)
		}
	})
}
