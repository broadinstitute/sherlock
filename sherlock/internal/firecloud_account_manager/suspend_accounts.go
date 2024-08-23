package firecloud_account_manager

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/bits_data_warehouse"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/advisory_locks"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strings"
	"time"
)

func (m *firecloudAccountManager) suspendAccounts(ctx context.Context) ([]string, []error) {
	results := make([]string, 0)
	errs := make([]error, 0)

	err := m.dbForLocking.Transaction(func(tx *gorm.DB) error {

		// Lock what's essentially the entry *in the config file* for this Firecloud account manager.
		// In other words, prevent anything else from running suspendAccounts on this same entry of
		// the config file at the same time.
		//
		// (Jack): Judgement call that it's more readable to just run the function that does the
		// locking rather than trying to abstract it. Part of my reasoning is that other advisory
		// locks will usually be coupled to a model, and then the model is the abstraction... so an
		// abstraction for this would only be used here.
		if err := tx.Exec(
			// language=SQL
			"SELECT pg_advisory_xact_lock(?, ?)",
			advisory_locks.FIRECLOUD_ACCOUNT_MANAGER,
			m.indexPlusOneForLocking,
		).Error; err != nil {
			return fmt.Errorf("failed to lock %T (%s) for suspension: %w", m, m.Domain, err)
		}

		currentUsers, err := m.workspaceClient.GetCurrentUsers(ctx, m.Domain)
		if err != nil {
			return fmt.Errorf("failed to get current users: %w", err)
		}

		for _, user := range currentUsers {

			// If we have a list of emails to restrict what we affect, make sure this user's email is in that list.
			if len(m.OnlyAffectEmails) > 0 && !utils.Contains(m.OnlyAffectEmails, user.PrimaryEmail) {
				continue
			}

			// Make sure this user's email isn't in the list of emails to never affect.
			if utils.Contains(m.NeverAffectEmails, user.PrimaryEmail) {
				continue
			}

			// If the user is already suspended, skip them.
			if user.Suspended {
				continue
			}

			// suspensionReason, if not empty, is why we should suspend the user. If it's empty that means
			// there's no reason to. It starts empty and we'll fill it in as we go.
			var suspensionReason string

			if user.LastLoginTime == "" {
				// If there's no last login time, this is a new account. Permit a grace period.
				if creationTime, err := time.Parse(time.RFC3339, user.CreationTime); err != nil {
					errs = append(errs, fmt.Errorf("failed to parse creation time %s for new account %s: %w", user.CreationTime, user.PrimaryEmail, err))
				} else if time.Since(creationTime) > m.NewAccountGracePeriod {
					suspensionReason = "due to being new but not setting up their account"
				}
			} else {
				// If the last login time is more than the threshold, suspend the user.
				if lastLoginTime, err := time.Parse(time.RFC3339, user.LastLoginTime); err != nil {
					errs = append(errs, fmt.Errorf("failed to parse last login time %s for %s: %w", user.LastLoginTime, user.PrimaryEmail, err))
				} else if time.Since(lastLoginTime) > m.InactivityThreshold {
					suspensionReason = "due to inactivity"
				}
			}

			// If we don't have a reason to suspend the user yet, check in BITS data. If there's no correlating
			// record, then we'll suspend. This mostly means "no longer employed by the Broad".
			if suspensionReason == "" {
				// If the user doesn't correspond to an @broadinstitute.org user in BITS data, suspend the user.
				// We call it "missing in BITS data" but this mostly means "no longer employed by the Broad".
				biEmail := strings.Split(user.PrimaryEmail, "@")[0] + "@broadinstitute.org"
				if _, found, err := bits_data_warehouse.GetPerson(biEmail); err != nil {
					errs = append(errs, fmt.Errorf("failed to get person %s for %s: %w", biEmail, user.PrimaryEmail, err))
				} else if !found {
					suspensionReason = "due to missing in BITS data"
				}
			}

			// If we have a reason to suspend the user, do it.
			if suspensionReason != "" {
				if m.DryRun {
					results = append(results, fmt.Sprintf("Would've suspended user %s (%s) but dry run enabled", user.PrimaryEmail, suspensionReason))
				} else {
					if err := m.workspaceClient.SuspendUser(ctx, user.PrimaryEmail); err != nil {
						errs = append(errs, fmt.Errorf("failed to suspend user %s (%s): %w", user.PrimaryEmail, suspensionReason, err))
					} else {
						results = append(results, fmt.Sprintf("Suspended user %s (%s)", user.PrimaryEmail, suspensionReason))
					}
				}
			}
		}

		return nil
	})
	if err != nil {
		errs = append(errs, err)
	}

	if len(results) > 0 || len(errs) > 0 {
		slack.SendPermissionChangeNotification(ctx, models.SelfUser.SlackReference(true), slack.PermissionChangeNotificationInputs{
			Summary: fmt.Sprintf("Firecloud account manager for %s ran", m.Domain),
			Results: results,
			Errors:  errs,
		})
	} else {
		log.Info().Msgf("FCAM | Firecloud account manager for %s ran and found no accounts to suspend", m.Domain)
	}

	return results, errs
}
