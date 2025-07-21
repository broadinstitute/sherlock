package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/jackc/pgx/v5"
	pgxstdlib "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	gormpg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect provides a *gorm.DB and a non-nil cleanup function.
func Connect() (db *gorm.DB, cleanup func() error, err error) {
	// To avoid a gotcha we provide a cleanup function so that's never nil
	cleanup = func() error { return nil }

	logLevel, err := parseGormLogLevel(config.Config.String("db.log.level"))
	if err != nil {
		return nil, cleanup, err
	}

	var connPool gorm.ConnPool
	connPool, cleanup, err = initializeConnPool()
	if err != nil {
		return nil, cleanup, err
	}

	initialAttempts := config.Config.Int("db.retryConnection.times") + 1
	attemptInterval := config.Config.Duration("db.retryConnection.interval")
	for attemptsRemaining := initialAttempts; attemptsRemaining >= 0; attemptsRemaining-- {
		db, err = initializeGorm(connPool, logLevel)
		if err == nil {
			break
		} else if attemptsRemaining > 0 {
			log.Info().Msgf("DB   | will attempt database connection %d more times; waiting %s", attemptsRemaining-1, attemptInterval)
			time.Sleep(attemptInterval)
		}
	}
	if err != nil || db == nil {
		return nil, cleanup, fmt.Errorf("unable to connect to the database after %d attempts: %w", initialAttempts, err)
	}

	if config.Config.MustString("mode") == "debug" {
		panicIfLooksLikeCloudSQL(db)
	}

	return db, cleanup, nil
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

// initializeConnPool sets up a connection to the database but doesn't actually run it,
// so it shouldn't need to be retried. This function will handle using the Cloud SQL
// Go connector if necessary and it's also responsible for setting the prepared statement
// cache setting (since we must manually do so with the Cloud SQL Go connector).
func initializeConnPool() (connPool gorm.ConnPool, cleanup func() error, err error) {
	// To avoid a gotcha we provide a cleanup function so that's never nil
	cleanup = func() error { return nil }

	pgxConfig, err := pgx.ParseConfig(dbConnectionString())
	if err != nil {
		return nil, cleanup, err
	}

	if config.Config.MustString("db.driver") == "cloudsql-postgres" {
		instanceConnectionName := config.Config.String("db.host")
		if instanceConnectionName == "" {
			return nil, cleanup, errors.New("db.driver=cloudsql-postgres requires db.host to be set to the instance connection name")
		}

		opts := make([]cloudsqlconn.Option, 0)
		if config.Config.Bool("db.cloudSql.automaticIamAuthEnabled") {
			opts = append(opts, cloudsqlconn.WithIAMAuthN())
		}
		var dialer *cloudsqlconn.Dialer
		dialer, err = cloudsqlconn.NewDialer(context.Background(), opts...)
		if err != nil {
			return nil, cleanup, err
		}
		cleanup = dialer.Close
		pgxConfig.DialFunc = func(ctx context.Context, _, _ string) (net.Conn, error) {
			return dialer.Dial(ctx, instanceConnectionName)
		}
		// When we use this dialer, the db.host is the instance connection name. This is fairly standard
		// per the Cloud SQL Go connector's documentation -- it's what you do if you use the connector
		// as the entire driver -- but for us, we're only using it as the dialer. The problem is that
		// pgx will still try to resolve it to get an IP address (which we end up ignoring in our
		// pgxConfig.DialFunc). The cleanest way to resolve this is to just set the host as far as
		// pgx is concerned to "localhost". This cleanly resolves to a harmless IP that we can safely
		// ignore.
		pgxConfig.Host = "localhost"
	}

	if config.Config.Bool("db.preparedStatementCache") {
		pgxConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheStatement
	} else {
		pgxConfig.DefaultQueryExecMode = pgx.QueryExecModeExec
	}

	sqlDB := pgxstdlib.OpenDB(*pgxConfig)
	sqlDB.SetMaxOpenConns(config.Config.MustInt("db.maxOpenConnections"))
	sqlDB.SetMaxIdleConns(config.Config.MustInt("db.maxIdleConnections"))
	sqlDB.SetConnMaxIdleTime(config.Config.Duration("db.connectionMaxIdleTime"))
	sqlDB.SetConnMaxLifetime(config.Config.Duration("db.connectionMaxLifetime"))

	return sqlDB, cleanup, nil
}

// initializeGorm does what it says on the tin. It will automatically attempt to ping the
// database, so it will fail if the database is offline.
func initializeGorm(conn gorm.ConnPool, logLevel logger.LogLevel) (*gorm.DB, error) {
	return gorm.Open(gormpg.New(gormpg.Config{
		Conn: conn,
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
func panicIfLooksLikeCloudSQL(db *gorm.DB) {
	var cloudSqlAdminRoleExists bool
	err := db.Raw("SELECT 1 FROM pg_roles WHERE rolname='cloudsqladmin'").Scan(&cloudSqlAdminRoleExists).Error
	if err != nil && !errors.Is(err, sql.ErrNoRows) && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(fmt.Errorf("failed to double-check that the database wasn't running in Cloud SQL: %w", err))
	}
	if cloudSqlAdminRoleExists {
		panic(fmt.Errorf("this database looks like it is running in Cloud SQL, refusing to proceed with test harness"))
	}
}
