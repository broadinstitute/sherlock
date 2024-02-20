package db

import (
	"cloud.google.com/go/cloudsqlconn"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"gorm.io/gorm/logger"
	"strings"
	"time"

	"cloud.google.com/go/cloudsqlconn/postgres/pgxv5"
	migrationFiles "github.com/broadinstitute/sherlock/sherlock/db"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/rs/zerolog/log"
	gormpg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// RegisterDriver handles setup for the cloudsql-postgres driver, if it's what'll be used.
// That driver can't be used for tests (as a Cloud SQL database can't be used for tests)
// so it is safe for tests to skip calling this and handling the resulting cleanup function.
func RegisterDriver() (cleanup func() error, err error) {
	if config.Config.MustString("db.driver") == "cloudsql-postgres" {
		opts := make([]cloudsqlconn.Option, 0)
		if config.Config.Bool("db.cloudSql.automaticIamAuthEnabled") {
			opts = append(opts, cloudsqlconn.WithIAMAuthN())
		}
		return pgxv5.RegisterDriver("cloudsql-postgres", opts...)
	} else {
		return nil, nil
	}
}

func dbConnectionString() string {
	parts := make([]string, 3, 6)
	parts[0] = fmt.Sprintf("host=%s", config.Config.MustString("db.host"))
	parts[1] = fmt.Sprintf("user=%s", config.Config.MustString("db.user"))
	parts[2] = fmt.Sprintf("dbname=%s", config.Config.MustString("db.name"))
	if config.Config.String("db.password") != "" {
		parts = append(parts, fmt.Sprintf("password=%s", config.Config.MustString("db.password")))
	}
	if config.Config.String("db.port") != "" {
		parts = append(parts, fmt.Sprintf("port=%s", config.Config.MustString("db.port")))
	}
	if config.Config.String("db.ssl") != "" {
		parts = append(parts, fmt.Sprintf("sslmode=%s", config.Config.MustString("db.ssl")))
	}
	return strings.Join(parts, " ")
}

func Connect() (*sql.DB, error) {
	sqlDB, err := sql.Open(config.Config.MustString("db.driver"), dbConnectionString())
	if err != nil {
		return nil, fmt.Errorf("error building SQL connection: %w", err)
	}

	sqlDB.SetMaxOpenConns(config.Config.MustInt("db.maxOpenConnections"))
	sqlDB.SetMaxIdleConns(config.Config.MustInt("db.maxIdleConnections"))
	sqlDB.SetConnMaxIdleTime(config.Config.Duration("db.connectionMaxIdleTime"))
	sqlDB.SetConnMaxLifetime(config.Config.Duration("db.connectionMaxLifetime"))

	initialAttempts := config.Config.Int("db.retryConnection.times") + 1
	for attemptsRemaining := initialAttempts; attemptsRemaining >= 0; attemptsRemaining-- {
		if err = sqlDB.Ping(); err == nil {
			return sqlDB, nil
		} else if attemptsRemaining > 0 {
			interval := config.Config.String("db.retryConnection.interval")
			if duration, durationErr := time.ParseDuration(interval); durationErr == nil {
				log.Info().Msgf("DB   | will attempt database connection %d more times; waiting %s", attemptsRemaining-1, interval)
				time.Sleep(duration)
			} else {
				log.Warn().Msgf("DB   | while retrying database connection, couldn't parse sleep interval duration %s: %v", interval, durationErr)
			}
		}
	}

	if config.Config.MustString("mode") == "debug" {
		panicIfLooksLikeCloudSQL(sqlDB)
	}

	return nil, fmt.Errorf("unable to connect to the database after %d attempts: %w", initialAttempts, err)
}

func applyMigrations(db *sql.DB) error {
	if !config.Config.Bool("db.init") {
		log.Info().Msg("DB   | skipping database migrations")
		return nil
	}

	log.Info().Msg("DB   | executing database migration")
	directory, err := iofs.New(migrationFiles.MigrationFiles, "migrations")
	if err != nil {
		return fmt.Errorf("error accessing embedded migration files: %w", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error building postgres driver instance for migration: %w", err)
	}
	migrationPlan, err := migrate.NewWithInstance(
		"iofs", directory,
		config.Config.MustString("db.name"), driver,
	)

	if err != nil {
		return fmt.Errorf("error building migration plan: %w", err)
	}
	if err = migrationPlan.Up(); errors.Is(err, migrate.ErrNoChange) {
		log.Info().Msg("DB   | no migration to apply, continuing")
	} else if err != nil {
		return fmt.Errorf("error applying migration plan: %w", err)
	}

	log.Info().Msg("DB   | database migration complete")
	return nil
}

func openGorm(conn gorm.ConnPool) (*gorm.DB, error) {
	logLevel, err := parseGormLogLevel(config.Config.String("db.log.level"))
	if err != nil {
		return nil, err
	}
	return gorm.Open(
		gormpg.New(gormpg.Config{
			Conn:                 conn,
			PreferSimpleProtocol: !config.Config.Bool("db.preparedStatementCache"),
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
				SlowThreshold:             config.Config.Duration("db.log.slowThreshold"),
				LogLevel:                  logLevel,
				IgnoreRecordNotFoundError: config.Config.Bool("db.log.ignoreNotFoundWarning"),
				Colorful:                  config.Config.String("mode") == "debug",
			}),
		})
}

func parseGormLogLevel(logLevel string) (logger.LogLevel, error) {
	switch logLevel {
	case "silent":
		return logger.Silent, nil
	case "error":
		return logger.Error, nil
	case "warn":
		return logger.Warn, nil
	case "info":
		return logger.Info, nil
	default:
		return 0, fmt.Errorf("unknown db log level '%s'", logLevel)
	}
}

func Configure(sqlDB *sql.DB) (*gorm.DB, error) {
	if err := applyMigrations(sqlDB); err != nil {
		return nil, fmt.Errorf("error migrating database: %w", err)
	}
	gormDB, err := openGorm(sqlDB)
	if err != nil {
		return nil, fmt.Errorf("error opening gorm: %w", err)
	}
	return gormDB, nil
}

// panicIfLooksLikeCloudSQL does what it says on the tin -- it exits fast and hard if the database has a 'cloudsqladmin'
// role in it. That's not something Sherlock's migration would ever add but it's there by default on Cloud SQL, so
// it's an extra gate to make sure we don't accidentally run tests against a remote database.
func panicIfLooksLikeCloudSQL(db *sql.DB) {
	var cloudSqlAdminRoleExists bool
	err := db.QueryRow("SELECT 1 FROM pg_roles WHERE rolname='cloudsqladmin'").Scan(&cloudSqlAdminRoleExists)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(fmt.Errorf("failed to double-check that the database wasn't running in Cloud SQL: %w", err))
	}
	if cloudSqlAdminRoleExists {
		panic(fmt.Errorf("this database looks like it is running in Cloud SQL, refusing to proceed with test harness"))
	}
}
