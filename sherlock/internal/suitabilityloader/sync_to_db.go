package suitabilityloader

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
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

	for _, suitability := range suitabilities {
		if err = superUserDB.
			Where(&models.Suitability{
				Email: suitability.Email,
			}).
			Assign(&models.Suitability{
				// We want to be explicit about this record being updated:
				// below, we remove results that haven't been updated recently
				UpdatedAt:   time.Now(),
				Suitable:    suitability.Suitable,
				Description: suitability.Description,
			}).
			FirstOrCreate(&suitability).Error; err != nil {
			return fmt.Errorf("failed to update suitability for %s: %w", *suitability.Email, err)
		}
	}

	// We just set the updated_at time above. If there's any records that haven't been updated
	// recently, we should remove them -- they were probably removed from config or firecloud.org
	var removedSuitabilities []models.Suitability
	if err = superUserDB.
		Clauses(clause.Returning{}).
		Where("updated_at < current_timestamp - '1 hour'::interval").
		Delete(&removedSuitabilities).Error; err != nil {
		return fmt.Errorf("failed to find removed suitabilities: %w", err)
	}
	if len(removedSuitabilities) > 0 {
		log.Info().Msgf("removed %d suitabilities: %v",
			len(removedSuitabilities),
			utils.Map(removedSuitabilities, func(s models.Suitability) string {
				if s.Email != nil {
					return *s.Email
				} else {
					return "nil?"
				}
			}))
	}

	return nil
}
