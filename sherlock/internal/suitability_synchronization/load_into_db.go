package suitability_synchronization

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func KeepLoadingIntoDB(ctx context.Context, db *gorm.DB) {
	if config.Config.Bool("suitabilitySynchronization.enable") && config.Config.Bool("suitabilitySynchronization.behaviors.loadIntoDB.enable") {
		interval := config.Config.MustDuration("suitabilitySynchronization.behaviors.loadIntoDB.interval")
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := LoadIntoDB(ctx, db); err != nil {
					log.Warn().Err(err).Msgf("failed to update suitability table: %v", err)
				}
			}
			time.Sleep(interval)
		}
	}
}

func LoadIntoDB(ctx context.Context, db *gorm.DB) error {
	if !config.Config.Bool("suitabilitySynchronization.enable") || !config.Config.Bool("suitabilitySynchronization.behaviors.loadIntoDB.enable") {
		return nil
	}
	suitabilitiesFromFirecloud, err := fromFirecloud(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "dailyLimitExceeded") {
			log.Warn().Err(err).Msgf("failed to load Firecloud data due to quota limit; absorbing error (see DDO-3765) and relying on existing suitability records")
			return nil
		}
		return err
	}

	// Assume super-user privileges for this operation (required to edit this table)
	superUserDB := models.SetCurrentUserForDB(db, models.SelfUser)

	for _, suitability := range suitabilitiesFromFirecloud {
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
	// recently, we should remove them -- they were probably removed from config or firecloud.org.
	// We calculate this based on the update interval -- if three intervals have passed since the
	// record has been updated, we remove it.
	var removedSuitabilities []models.Suitability
	if err = superUserDB.
		Clauses(clause.Returning{}).
		Where("updated_at < ?", time.Now().Add(config.Config.MustDuration("suitabilitySynchronization.behaviors.loadIntoDB.interval")*-3).Format(time.RFC3339)).
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
