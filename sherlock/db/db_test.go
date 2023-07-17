package db

import (
	"github.com/stretchr/testify/assert"
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
	for _, filename := range filenames {
		if strings.HasSuffix(filename, ".down.sql") {
			assert.Contains(t, filenames, strings.TrimSuffix(filename, ".down.sql")+".up.sql")
		} else if strings.HasSuffix(filename, ".up.sql") {
			assert.Contains(t, filenames, strings.TrimSuffix(filename, ".up.sql")+".down.sql")
		} else {
			assert.Failf(t, "file lacks .up.sql or .down.sql extension", "filename '%s'", filename)
		}
	}
}
