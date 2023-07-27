package deprecated_db

import (
	"database/sql"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strings"
	"sync"
	"testing"
)

var (
	// The model hierarchies must be in dependency order, so that the first model has no dependencies,
	// the second may only depend on the first, and so on.
	v2ModelHierarchy = []any{
		&v2models.CiIdentifier{},
		&v2models.CiRun{},
		&v2models.User{},
		&v2models.PagerdutyIntegration{},
		&v2models.Cluster{},
		&v2models.Environment{},
		&v2models.Chart{},
		&v2models.ChartVersion{},
		&v2models.AppVersion{},
		&v2models.ChartRelease{},
		&v2models.Changeset{},
		&v2models.DatabaseInstance{},
	}
)

var _mutexExistingTestDB = &sync.Mutex{}
var _existingTestDB *sql.DB = nil

// ConnectAndConfigureFromTest is like Connect and Configure but accepts a testing.T in exchange for never returning an
// error--the test will be failed instead if there is one.
func ConnectAndConfigureFromTest(t *testing.T) *gorm.DB {
	if _existingTestDB == nil {
		_mutexExistingTestDB.Lock()
		defer _mutexExistingTestDB.Unlock()
		sqlDB, err := db.Connect()
		if err != nil {
			t.Errorf("failed to connect to database during test: %v", err)
			return nil
		}
		_existingTestDB = sqlDB
	}
	gormDB, err := db.Configure(_existingTestDB)
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
func Truncate(t *testing.T, gormDB *gorm.DB) {
	if t == nil {
		log.Fatal().Msg("refusing to truncate, was not passed a testing.T")
		return
	}
	if mode := config.Config.MustString("mode"); mode != "debug" {
		t.Errorf("refusing to truncate database, mode is '%s' instead of 'debug'", mode)
		return
	}
	if sqlDB, err := gormDB.DB(); err != nil {
		log.Fatal().Msgf("refusing to truncate, failed to get sql.DB from gorm.DB: %v", err)
	} else {
		db.PanicIfLooksLikeCloudSQL(sqlDB)
	}
	var statements []string
	dryRunDB := gormDB.Session(&gorm.Session{
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
	for i := len(v2ModelHierarchy) - 1; i >= 0; i-- {
		// Unscoped disables soft-deletion handling
		statements = append(statements, dryRunDB.Unscoped().Delete(v2ModelHierarchy[i]).Statement.SQL.String())
	}
	statements = append(statements, "COMMIT")
	if err := gormDB.Exec(strings.Join(statements, ";\n")).Error; err != nil {
		t.Errorf("failed to clean/truncate database: %v", err)
	}
}
