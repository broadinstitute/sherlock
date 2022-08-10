package db

import (
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strings"
	"testing"
)

// ConnectAndConfigureFromTest is like Connect and Configure but accepts a testing.T in exchange for never returning an
// error--the test will be failed instead if there is one.
func ConnectAndConfigureFromTest(t *testing.T) *gorm.DB {
	sqlDB, err := Connect()
	if err != nil {
		t.Errorf("failed to connect to database during test: %v", err)
		return nil
	}
	gormDB, err := Configure(sqlDB)
	if err != nil {
		t.Errorf("failed to configure database during test: %v", err)
	}
	return gormDB
}

// Truncate cleans up tables, intended for usage with functional tests. It will refuse to run if
// not given a testing.T or if Sherlock's current mode is anything other than 'debug'.
//
// As an implementation note, the SQL this function runs is fairly consistent, but as Jack discovered,
// it is surprisingly annoying to debug, because it requires specific ordering, must be run manually,
// and (with the addition of v2 models) has effectively variable table names. This function handles
// and documents that complexity so we don't need to rediscover it each time we make a database change.
func Truncate(t *testing.T, db *gorm.DB) {
	if t == nil {
		log.Fatal().Msg("refusing to truncate, was not passed a testing.T")
		return
	}
	if mode := config.Config.MustString("mode"); mode != "debug" {
		t.Errorf("refusing to truncate database, mode is '%s' instead of 'debug'", mode)
		return
	}
	var statements []string
	dryRunDB := db.Session(&gorm.Session{
		// Performance boost, don't transact each delete individually
		SkipDefaultTransaction: true,
		// Disable Gorm's guardrails against doing unqualified batch deletions
		AllowGlobalUpdate: true,
		// We can't let Gorm actually run this for us, because it tries to be helpful and fail-fast
		// when there's been an exception in an outer transaction. We don't want that because we
		// want to use this function to clean up--and we often cause database errors intentionally
		// in tests. Instead, we set the database to dry-run, grab the SQL, and manually execute it.
		DryRun: true,
	})
	statements = append(statements, "BEGIN")
	// We must iterate backwards through the hierarchies so we don't violate foreign key restraints while we delete
	for i := len(v1ModelHierarchy) - 1; i >= 0; i-- {
		statements = append(statements, dryRunDB.Delete(v1ModelHierarchy[i]).Statement.SQL.String())
	}
	for i := len(v2ModelHierarchy) - 1; i >= 0; i-- {
		// Unscoped disables soft-deletion handling
		statements = append(statements, dryRunDB.Unscoped().Delete(v2ModelHierarchy[i]).Statement.SQL.String())
	}
	statements = append(statements, "COMMIT")
	if err := db.Exec(strings.Join(statements, ";\n")).Error; err != nil {
		t.Errorf("failed to clean/truncate database: %v", err)
	}
}
