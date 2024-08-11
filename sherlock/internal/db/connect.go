package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	gormpg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*gorm.DB, error) {
	logLevel, err := parseGormLogLevel(config.Config.String("db.log.level"))
	if err != nil {
		return nil, err
	}

	var db *gorm.DB
	initialAttempts := config.Config.Int("db.retryConnection.times") + 1
	attemptInterval := config.Config.Duration("db.retryConnection.interval")
	for attemptsRemaining := initialAttempts; attemptsRemaining >= 0; attemptsRemaining-- {
		db, err = initializeGorm(logLevel)
		if err == nil {
			break
		} else if attemptsRemaining > 0 {
			log.Info().Msgf("DB   | will attempt database connection %d more times; waiting %s", attemptsRemaining-1, attemptInterval)
			time.Sleep(attemptInterval)
		}
	}
	if err != nil || db == nil {
		return nil, fmt.Errorf("unable to connect to the database after %d attempts: %w", initialAttempts, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(config.Config.MustInt("db.maxOpenConnections"))
	sqlDB.SetMaxIdleConns(config.Config.MustInt("db.maxIdleConnections"))
	sqlDB.SetConnMaxIdleTime(config.Config.Duration("db.connectionMaxIdleTime"))
	sqlDB.SetConnMaxLifetime(config.Config.Duration("db.connectionMaxLifetime"))

	if config.Config.MustString("mode") == "debug" {
		panicIfLooksLikeCloudSQL(sqlDB)
	}

	return db, nil
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

// initializeGorm does what it says on the tin. It will automatically attempt to ping the
// database, so it will fail if the database is offline.
//
// Gorm uses pgx under the hood to work with Postgres; RegisterDriver will have wired up driver
// into pgx if necessary. We use Gorm to do the heavy lifting so our prepared statement cache
// setting is respected; since that's a pgx thing if we set up a *sql.DB ourselves and just
// hand it to Gorm the setting won't work.
func initializeGorm(logLevel logger.LogLevel) (*gorm.DB, error) {
	return gorm.Open(gormpg.New(gormpg.Config{
		DriverName:           config.Config.MustString("db.driver"),
		DSN:                  dbConnectionString(),
		PreferSimpleProtocol: !config.Config.Bool("db.preparedStatementCache"),
	}), &gorm.Config{
		// log.Logger is Zerolog's global logger that the rest of Sherlock uses
		Logger: logger.New(&log.Logger, logger.Config{
			SlowThreshold:             config.Config.Duration("db.log.slowThreshold"),
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: config.Config.Bool("db.log.ignoreNotFoundWarning"),
			Colorful:                  config.Config.String("mode") == "debug",
		}),
		// This is to account for the fact that go and postgres have different
		// time stamp precision which causes issues in testing.
		// This is a fix to have gorm round down timestamps to postgres's millisecond
		// precision
		NowFunc: func() time.Time {
			return time.Now().Round(time.Millisecond)
		},
	})
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
