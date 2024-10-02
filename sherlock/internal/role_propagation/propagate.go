package role_propagation

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"time"
)

func DoOnDemandPropagation(ctx context.Context, db *gorm.DB, roleID uint) {
	if config.Config.Bool("rolePropagation.enable") {
		if config.Config.Bool("rolePropagation.asynchronous") {
			go waitToPropagate(ctx, db, roleID)
		} else {
			waitToPropagate(ctx, db, roleID)
		}
	}
}

// waitToPropagate is a blocking function that will forcibly run propagation for
// the given role. It will wait until it can acquire a propagation lock on the
// role.
func waitToPropagate(ctx context.Context, db *gorm.DB, roleID uint) {

	// Load the role so we can lock it
	var role models.Role
	if err := db.Take(&role, roleID).Error; err != nil {
		log.Error().Err(err).Msgf("failed to load role %d", roleID)
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		// Acquire lock, which will be released when the transaction ends
		if err := role.WaitPropagationLock(tx); err != nil {
			return err
		}

		// Reload the role with all associations now that we've locked it
		if err := db.Scopes(models.ReadRoleScope).Take(&role, roleID).Error; err != nil {
			return err
		}

		// Do the propagation
		doNonConcurrentPropagation(ctx, role)

		// Update the role's PropagatedAt field
		return role.UpdatePropagatedAt(tx)
	})
	if err != nil {
		log.Error().Err(err).Msgf("failed to propagate role %s (%d)", *role.Name, roleID)
	}
}

// KeepPropagatingStale runs tryToPropagateStale every 30 seconds.
func KeepPropagatingStale(ctx context.Context, db *gorm.DB) {
	if config.Config.Bool("rolePropagation.enable") {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				tryToPropagateStale(ctx, db)
			}
			time.Sleep(30 * time.Second)
		}
	}
}

// tryToPropagateStale will attempt to Propagate all roles that are "stale" -- meaning they
// were last propagated past the threshold defined in the config. It will skip roles that
// have been propagated more recently or ones that it can't immediately obtain a lock on
// (as this implies some other process is already looking at it).
func tryToPropagateStale(ctx context.Context, db *gorm.DB) {

	// Get the list of roles to Propagate
	var roleIDs []uint
	if err := db.Model(&models.Role{}).Pluck("id", &roleIDs).Error; err != nil {
		log.Error().Err(err).Msg("failed to get list of roles to propagate")
		return
	}

	for _, roleID := range roleIDs {

		// Load the role so we can lock it
		var role models.Role
		if err := db.Take(&role, roleID).Error; err != nil {
			log.Error().Err(err).Msgf("failed to load role %d", roleID)
			continue
		}

		err := db.Transaction(func(tx *gorm.DB) error {
			// Try to acquire lock, which will be released when the transaction ends
			if obtainedLock, err := role.TryPropagationLock(tx); err != nil {
				return err
			} else if !obtainedLock {
				// If it was already locked, another process is already looking at this role, so we skip
				return nil
			}

			// Reload the role with all associations now that we've locked it
			if err := db.Scopes(models.ReadRoleScope).Take(&role, roleID).Error; err != nil {
				return err
			}

			// If the role was propagated recently, we skip
			threshold := time.Now().Add(-1 * config.Config.Duration("rolePropagation.driftAlignmentStaleThreshold"))
			if role.PropagatedAt.Valid && role.PropagatedAt.Time.After(threshold) {
				return nil
			}

			// Do the propagation
			doNonConcurrentPropagation(ctx, role)

			// Update the role's PropagatedAt field
			return role.UpdatePropagatedAt(tx)
		})
		if err != nil {
			log.Error().Err(err).Msgf("failed to propagate role %s (%d)", *role.Name, roleID)
		}
	}
}

// doNonConcurrentPropagation runs each propagator in sequence for the given role.
// It assumes it won't run concurrently for the same role, even across replicas.
// This means that the propagators don't need to be idempotent (though the
// propagators could be used concurrently for different roles, so they can't be
// naively stateful).
func doNonConcurrentPropagation(ctx context.Context, role models.Role) {
	results := make([]string, 0)
	errors := make([]error, 0)
	for _, p := range propagators {
		additionalResults, additionalErrors := p.Propagate(ctx, role)
		results = append(results, utils.Map(additionalResults, func(result string) string {
			return p.LogPrefix() + result
		})...)
		errors = append(errors, utils.Map(additionalErrors, func(err error) error {
			return fmt.Errorf("%s%w", p.LogPrefix(), err)
		})...)
	}
	if len(results) > 0 || len(errors) > 0 {
		slack.SendPermissionChangeNotification(ctx, models.SelfUser.SlackReference(true), slack.PermissionChangeNotificationInputs{
			Summary: fmt.Sprintf("propagation for Role \"%s\" reports:", *role.Name),
			Results: results,
			Errors:  errors,
		})
	}
}
