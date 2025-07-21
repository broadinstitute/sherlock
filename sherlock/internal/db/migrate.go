package db

import (
	"errors"
	"fmt"

	migrationFiles "github.com/broadinstitute/sherlock/sherlock/db"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if !config.Config.Bool("db.init") {
		log.Info().Msg("DB   | skipping database migrations")
		return nil
	}

	log.Info().Msg("DB   | executing database migration")
	directory, err := iofs.New(migrationFiles.MigrationFiles, "migrations")
	if err != nil {
		return fmt.Errorf("error accessing embedded migration files: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	migrationDriver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		return fmt.Errorf("error building postgres driver instance for migration: %w", err)
	}

	migrationPlan, err := migrate.NewWithInstance(
		"iofs", directory,
		config.Config.MustString("db.name"), migrationDriver,
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
