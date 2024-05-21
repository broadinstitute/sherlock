package models

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/pact_broker"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/pagerduty"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"runtime/debug"
	"strings"
	"time"
)

func PlanChangesets(db *gorm.DB, changesets []Changeset) (created []uint, err error) {
	var user *User
	if user, err = GetCurrentUserForDB(db); err != nil {
		return nil, fmt.Errorf("unable to get current user for changeset planning: %w", err)
	}
	changesetsToCreate := make([]Changeset, 0, len(changesets))
	for _, changeset := range changesets {
		if changeset.ChartReleaseID == 0 {
			return nil, fmt.Errorf("changeset has no chart release specified")
		}
		var chartRelease ChartRelease
		if err = db.Take(&chartRelease, changeset.ChartReleaseID).Error; err != nil {
			return nil, fmt.Errorf("unable to load chart release %d to plan changeset: %w", changeset.ChartReleaseID, err)
		}

		changeset.From = chartRelease.ChartReleaseVersion

		changeset.fillEmptyToFieldsBasedOnFrom()

		if err = changeset.To.resolve(db, chartRelease.ChartID); err != nil {
			return nil, fmt.Errorf("unable to resolve 'to' versions for changeset against %s: %w", chartRelease.Name, err)
		}

		changeset.PlannedByID = &user.ID

		if changeset.hasDiff() {
			changesetsToCreate = append(changesetsToCreate, changeset)
		}
	}
	created = make([]uint, 0, len(changesetsToCreate))
	if len(changesetsToCreate) > 0 {
		err = db.Transaction(func(tx *gorm.DB) error {
			for _, changeset := range changesetsToCreate {
				if err = tx.Create(&changeset).Error; err != nil {
					return fmt.Errorf("unable to create changeset against chart release %d: %w", changeset.ChartReleaseID, err)
				}
				created = append(created, changeset.ID)
			}
			return nil
		})
	}
	return
}

func ApplyChangesets(db *gorm.DB, ids []uint) (err error) {
	affectedChartReleases := make(map[uint]Changeset, len(ids))
	changesetsToApply := make([]Changeset, 0, len(ids))
	now := time.Now()
	var user *User
	if user, err = GetCurrentUserForDB(db); err != nil {
		return fmt.Errorf("unable to get current user for changeset applying: %w", err)
	}
	err = db.Transaction(func(tx *gorm.DB) error {

		// Load changesets, lock them and their chart releases
		if err = db.
			InnerJoins(
				"ChartRelease",
				db.Clauses(clause.Locking{
					Strength: "SHARE",
					Table:    clause.Table{Name: "chart_releases"},
				})).
			Clauses(clause.Locking{
				Strength: "SHARE",
				Table:    clause.Table{Name: "changesets"},
			}).
			Find(&changesetsToApply, ids).Error; err != nil {
			return fmt.Errorf("unable to load changesets to apply: %w", err)
		}

		// Validation pass, so we error out quickly if there's an obvious problem (don't want to hold lock)
		for _, changeset := range changesetsToApply {
			// Check for two changesets affecting the same chart release
			if existingChangeset, exists := affectedChartReleases[changeset.ChartReleaseID]; exists {
				return fmt.Errorf("(%s) can't apply, changesets %d and %d both affect chart release '%s'", errors.Conflict, existingChangeset.ID, changeset.ID, changeset.ChartRelease.Name)
			} else {
				affectedChartReleases[changeset.ChartReleaseID] = changeset
			}

			// Check for changeset state
			if changeset.AppliedAt != nil {
				return fmt.Errorf("(%s) changeset %d against chart release '%s' already applied", errors.BadRequest, changeset.ID, changeset.ChartRelease.Name)
			} else if changeset.SupersededAt != nil {
				return fmt.Errorf("(%s) changeset %d against chart release '%s' already superseded", errors.BadRequest, changeset.ID, changeset.ChartRelease.Name)
			} else if changeset.From.ResolvedAt == nil {
				return fmt.Errorf("(%s) changeset %d against chart release '%s' has unresolved 'from' version", errors.BadRequest, changeset.ID, changeset.ChartRelease.Name)
			} else if changeset.To.ResolvedAt == nil {
				// SQL validation in theory prevents this, but we'll check anyway
				return fmt.Errorf("(%s) changeset %d against chart release '%s' has unresolved 'to' version", errors.BadRequest, changeset.ID, changeset.ChartRelease.Name)
			} else if changeset.From.hasDiffWith(&changeset.ChartRelease.ChartReleaseVersion) {
				return fmt.Errorf("(%s) changeset %d has 'from' versions that don't match with chart release '%s', so it can't be safely applied (it probably should be marked as superseded but isn't, so let DevOps know); running another plan and applying that instead may resolve the situation", errors.Conflict, changeset.ID, changeset.ChartRelease.Name)
			}
		}

		// Apply pass
		for _, changeset := range changesetsToApply {

			// Edit the chart release
			changeset.ChartRelease.ChartReleaseVersion = changeset.To
			if err = tx.Model(&changeset.ChartRelease).Select("*").Updates(&changeset.ChartRelease).Error; err != nil {
				return fmt.Errorf("unable to apply changeset %d against chart release '%s': %w", changeset.ID, changeset.ChartRelease.Name, err)
			}

			// Edit the changeset
			if err = tx.Model(&changeset).Updates(&Changeset{
				AppliedByID: &user.ID,
				AppliedAt:   &now,
			}).Error; err != nil {
				return fmt.Errorf("unable to record changeset %d apply against chart release '%s': %w", changeset.ID, changeset.ChartRelease.Name, err)
			}

			// Now mark any changesets that aren't either applied or superseded as superseded
			if err = tx.Model(&Changeset{}).
				Where("chart_release_id = ? and applied_at is null and superseded_at is null", changeset.ChartReleaseID).
				Update("superseded_at", &now).Error; err != nil {
				return fmt.Errorf("unable to mark other changesets as superseded for chart release '%s': %w", changeset.ChartRelease.Name, err)
			}
		}

		return nil
	})

	// If the update was successful, spin up some goroutines to let third parties know
	if err == nil {
		changesetPostApplyActions(db, changesetsToApply)
	}

	return
}

func changesetPostApplyActions(db *gorm.DB, appliedChangesets []Changeset) {
	// If anything here panics, we don't want it to bubble up to the action apply request, so we forcibly recover
	defer func() {
		if r := recover(); r != nil {
			go slack.ReportError[error](context.Background(), fmt.Sprintf("panic in changesetPostApplyActions: %v", r))
			log.Error().Bytes("stack", debug.Stack()).Msg("panic in changesetPostApplyActions")
		}
	}()

	// We'll use a link to Beehive describing all the changesets in a few places
	beehiveLink := fmt.Sprintf("%s?%s",
		config.Config.String("beehive.reviewChangesetsUrl"),
		strings.Join(
			utils.Map(appliedChangesets, func(c Changeset) string { return fmt.Sprintf("changeset=%d", c.ID) }),
			"&"))

	// When we iterate over changesets, we'll load their environments. There's some actions we only want to do once per
	// environment, so we accumulate the environments initially and then iterate over them specifically later.
	affectedEnvironments := make(map[uint]Environment)

	// For each change:
	for _, appliedChangeset := range appliedChangesets {

		// 1. Report to pagerduty
		if appliedChangeset.ChartRelease.PagerdutyIntegrationID != nil {
			var pagerdutyIntegration PagerdutyIntegration
			if errToSwallow := db.
				Take(&pagerdutyIntegration, *appliedChangeset.ChartRelease.PagerdutyIntegrationID).Error; errToSwallow != nil {
				go slack.ReportError(context.Background(), fmt.Sprintf("unable to load pagerduty integration %d after applying changeset %d", *appliedChangeset.ChartRelease.PagerdutyIntegrationID, appliedChangeset.ID), errToSwallow)
			} else if pagerdutyIntegration.Key != nil {
				go pagerduty.SendChangeSwallowErrors(
					*pagerdutyIntegration.Key,
					fmt.Sprintf("Version changes to %s via Beehive", appliedChangeset.ChartRelease.Name),
					beehiveLink)
			}
		}

		// 2. Load the environment
		if appliedChangeset.ChartRelease.EnvironmentID != nil {
			var environment Environment
			if errToSwallow := db.
				Preload("PagerdutyIntegration").
				Take(&environment, *appliedChangeset.ChartRelease.EnvironmentID).Error; errToSwallow != nil {
				go slack.ReportError(context.Background(), fmt.Sprintf("unable to load environment %d after applying changeset %d", *appliedChangeset.ChartRelease.EnvironmentID, appliedChangeset.ID), errToSwallow)
			} else {
				affectedEnvironments[environment.ID] = environment

				// Disabled pending test coverage
				// TODO: DDO-3385
				//
				// (Comment: I [Jack] rewrote this as part of the v3 migration work and I believe I caught the NPE that
				// was the problem before. What we're still missing is a test that covers this code path, maybe with the
				// mock client pattern used elsewhere, so I'm leaving this commented out for now. If push really comes
				// to shove, I think we could re-enable it and at the very worst the outer function would now handle any
				// NPEs. That said, if push really comes to shove, it also means we actually *need* Sherlock's capability
				// here, so we should really test it.)
				//
				//// 3. Report to Pact
				if environment.PactIdentifier != nil {
					var chart Chart
					if errToSwallow := db.
						Take(&chart, appliedChangeset.ChartRelease.ChartID).Error; errToSwallow != nil {
						go slack.ReportError(context.Background(), fmt.Sprintf("unable to load chart %d after applying changeset %d", appliedChangeset.ChartRelease.ChartID, appliedChangeset.ID), errToSwallow)
					} else if chart.PactParticipant != nil && *chart.PactParticipant && appliedChangeset.To.AppVersionExact != nil {
						go pact_broker.RecordDeployment(
							chart.Name,
							*appliedChangeset.To.AppVersionExact,
							*environment.PactIdentifier)
					}
				}
			}
		}
	}

	for _, affectedEnvironment := range affectedEnvironments {

		// 4. Report to pagerduty (for environments)
		if affectedEnvironment.PagerdutyIntegration != nil && affectedEnvironment.PagerdutyIntegration.Key != nil {
			go pagerduty.SendChangeSwallowErrors(
				*affectedEnvironment.PagerdutyIntegration.Key,
				fmt.Sprintf("Version changes to %s via Beehive", affectedEnvironment.Name),
				beehiveLink)
		}
	}
}
