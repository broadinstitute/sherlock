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

// Google's API says that accounts that have never logged in actually have -- just at 1970-01-01T00:00:00.000Z.
// To provide good output, we set a threshold of 30 years and detect accounts with a login time older than that
// as never having logged in.
// We're being permissive here because Google doesn't document this behavior and the time returned isn't a
// Go zero-time.
const thresholdForLastLoginToDetectNeverLoggedIn = 30 * 365 * 24 * time.Hour

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

			// Check the user creation time. If it's so recent it's in the grace period, don't check inactivity for it.
			if parsedCreationTime, creationTimeParseErr := time.Parse(time.RFC3339, user.CreationTime); creationTimeParseErr != nil {
				errs = append(errs, fmt.Errorf("failed to parse creation time %s for %s: %w", user.CreationTime, user.PrimaryEmail, creationTimeParseErr))
			} else if time.Since(parsedCreationTime) > m.NewAccountGracePeriod {

				// If the user's out of the grace period, check log-in activity. We suspend in three cases: if we can't
				// tell if the user's logged in, if their last login is past our threshold for "never logged in", or if
				// they haven't logged in for the configured inactivity threshold.
				//
				// Why suspend if we can't parse the time? Because if we're here, we know that the user is out of the
				// grace period. Google doesn't document what they return if the user has never logged in. As of 2024
				// it seems to be 1970-01-01T00:00:00.000Z, but in case this turns into merely an empty string, we
				// treat any errors as "has not logged in".
				if parsedLastLoginTime, loginTimeParseErr := time.Parse(time.RFC3339, user.LastLoginTime); loginTimeParseErr != nil {
					suspensionReason = fmt.Sprintf("due to being unable to parse the last login time (and the account is out of the grace period): %v", loginTimeParseErr)
				} else if time.Since(parsedLastLoginTime) > thresholdForLastLoginToDetectNeverLoggedIn {
					suspensionReason = "due to being new but not setting up their account"
				} else if time.Since(parsedLastLoginTime) > m.InactivityThreshold {
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
