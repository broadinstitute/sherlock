package suitabilityloader

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func KeepSuitabilitiesInDBUpdated(ctx context.Context, db *gorm.DB) {
	interval := time.Duration(config.Config.MustInt("auth.updateIntervalMinutes")) * time.Minute
	for {
		time.Sleep(interval)
		if err := SyncSuitabilitiesToDB(ctx, db); err != nil {
			log.Warn().Err(err).Msgf("failed to update suitability table")
		}
	}
}

func SyncSuitabilitiesToDB(ctx context.Context, db *gorm.DB) error {
	suitabilitiesFromConfig, err := fromConfig()
	if err != nil {
		return err
	}
	suitabilitiesFromFirecloud, err := fromFirecloud(ctx)
	if err != nil {
		return err
	}
	suitabilities := append(suitabilitiesFromConfig, suitabilitiesFromFirecloud...)

	// Assume super-user privileges for this operation (required to edit this table)
	superUserDB := models.SetCurrentUserForDB(db, models.SelfUser)

	if err = superUserDB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&suitabilities).Error; err != nil {
		return err
	}

	// TODO: once we know that the updatedAt field is being set properly, we'll want to add a step here
	// to drop anything from the database that hasn't been updated recently. That'll match the current
	// functionality, where if someone completely disappears from firecloud.org (or more likely, from
	// Sherlock's config file), they'll cached suitability will eventually expire. With the old
	// in-memory solution we'd just replace the whole cache to do that, but since the database is
	// persistent we'll have to manually delete those rows.

	return nil
}
