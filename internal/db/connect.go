package db

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm/logger"
	"time"

	migrationFiles "github.com/broadinstitute/sherlock/db"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/rs/zerolog/log"
	gormpg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	// The model hierarchies must be in dependency order, so that the first model has no dependencies,
	// the second may only depend on the first, and so on.
	v1ModelHierarchy = []any{
		&v1models.AllocationPool{},
		&v1models.Service{},
		&v1models.Environment{},
		&v1models.Cluster{},
		&v1models.Build{},
		&v1models.ServiceInstance{},
		&v1models.Deploy{},
	}
	v2ModelHierarchy = []any{
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

func dbConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Config.MustString("db.user"),
		config.Config.MustString("db.password"),
		config.Config.MustString("db.host"),
		config.Config.MustString("db.port"),
		config.Config.MustString("db.name"),
		config.Config.MustString("db.ssl"),
	)
}

func Connect() (*sql.DB, error) {
	sqlDB, err := sql.Open("pgx", dbConnectionString())
	if err != nil {
		return nil, fmt.Errorf("error building SQL connection: %v", err)
	}
	initialAttempts := config.Config.Int("db.retryConnection.times") + 1
	for attemptsRemaining := initialAttempts; attemptsRemaining >= 0; attemptsRemaining-- {
		if err = sqlDB.Ping(); err == nil {
			return sqlDB, nil
		} else if attemptsRemaining > 0 {
			interval := config.Config.String("db.retryConnection.interval")
			if duration, durationErr := time.ParseDuration(interval); durationErr == nil {
				log.Debug().Msgf("will attempt database connection %d more times in %s", attemptsRemaining-1, interval)
				time.Sleep(duration)
			} else {
				log.Warn().Msgf("while retrying database connection, couldn't parse sleep interval duration %s: %v", interval, durationErr)
			}
		}
	}
	return nil, fmt.Errorf("unable to connect to the database after %d attempts: %v", initialAttempts, err)
}

func applyMigrations(db *sql.DB) error {
	if !config.Config.Bool("db.init") {
		log.Info().Msg("skipping database migrations")
		return nil
	}

	log.Info().Msg("executing database migration")
	directory, err := iofs.New(migrationFiles.MigrationFiles, "migrations")
	if err != nil {
		return fmt.Errorf("error accessing embedded migration files: %v", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error building postgres driver instance for migration: %v", err)
	}
	migrationPlan, err := migrate.NewWithInstance(
		"iofs", directory,
		config.Config.MustString("db.name"), driver,
	)

	if err != nil {
		return fmt.Errorf("error building migration plan: %v", err)
	}
	if err = migrationPlan.Up(); err == migrate.ErrNoChange {
		log.Info().Msg("no migration to apply, continuing")
	} else if err != nil {
		return fmt.Errorf("error applying migration plan: %v", err)
	}

	log.Info().Msg("database migration complete")
	return nil
}

func openGorm(db *sql.DB) (*gorm.DB, error) {
	logLevel, err := parseGormLogLevel(config.Config.String("db.log.level"))
	if err != nil {
		return nil, err
	}
	return gorm.Open(
		gormpg.New(gormpg.Config{
			Conn: db,
		}),

		&gorm.Config{
			// This is to account for the fact that go and postgres have different
			// time stamp precision which causes issues in testing.
			// This is a fix to have gorm round down timestamps to postgres's millisecond
			// precision
			NowFunc: func() time.Time {
				return time.Now().Round(time.Millisecond)
			},
			// log.Logger is Zerolog's global logger that the rest of Sherlock uses
			Logger: logger.New(&log.Logger, logger.Config{
				SlowThreshold:             config.Config.Duration("db.log.slowThresholdMs"),
				LogLevel:                  logLevel,
				IgnoreRecordNotFoundError: config.Config.Bool("db.log.ignoreNotFoundWarning"),
				Colorful:                  config.Config.String("mode") == "debug",
			}),
		})
}

// applyAutoMigrations is largely a development mechanism. It leverages https://gorm.io/docs/migration.html
// to do a best-effort migration based on Gorm's understanding of a model struct. It works well from a clean
// database but is a bit under-developed for what we need for full production migrations. Before shipping
// database changes, those changes should be represented as normal migration files so that this capability can
// be disabled.
func applyAutoMigrations(db *gorm.DB) error {
	if config.Config.Bool("db.autoMigrateV1") {
		log.Info().Msg("executing database auto-migrations for v1 models")
		if err := db.AutoMigrate(v1ModelHierarchy...); err != nil {
			return fmt.Errorf("error running v1 model auto-migrations: %v", err)
		}
		log.Info().Msg("database auto-migrations for v1 models complete")
	}
	if config.Config.Bool("db.autoMigrateV2") {
		log.Info().Msg("executing database auto-migrations for v2 models")
		if err := db.AutoMigrate(v2ModelHierarchy...); err != nil {
			return fmt.Errorf("error running v2 model auto-migrations: %v", err)
		}
		log.Info().Msg("database auto-migrations for v2 models complete")
	}
	return nil
}

func Configure(sqlDB *sql.DB) (*gorm.DB, error) {
	// defensively set max number of open connections to defend against contention issues
	maxOpenConnections := config.Config.MustInt("db.maxOpenConnections")
	sqlDB.SetMaxOpenConns(maxOpenConnections)

	if err := applyMigrations(sqlDB); err != nil {
		return nil, fmt.Errorf("error migrating database: %v", err)
	}
	gormDB, err := openGorm(sqlDB)
	if err != nil {
		return nil, fmt.Errorf("error opening gorm: %v", err)
	}
	if err = applyAutoMigrations(gormDB); err != nil {
		return nil, fmt.Errorf("error auto-migrating database: %v", err)
	}
	return gormDB, nil
}
