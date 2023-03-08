package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/pagerduty"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

type ChangesetEventStore struct {
	*ModelStore[Changeset]
	*internalChangesetEventStore
}

func (s *ChangesetEventStore) PlanAndApply(changesets []Changeset, user *auth_models.User) ([]Changeset, error) {
	var ret []Changeset
	err := s.db.Transaction(func(tx *gorm.DB) error {
		planned, err := s.plan(tx, changesets, user)
		if err != nil {
			return err
		}
		ret, err = s.apply(tx, planned, user)
		return err
	})
	return ret, err
}

func (s *ChangesetEventStore) Plan(changesets []Changeset, user *auth_models.User) ([]Changeset, error) {
	return s.plan(s.db, changesets, user)
}

func (s *ChangesetEventStore) Apply(selectors []string, user *auth_models.User) ([]Changeset, error) {
	var queries []Changeset
	for index, selector := range selectors {
		query, err := s.internalChangesetEventStore.selectorToQueryModel(s.db, selector)
		if err != nil {
			return []Changeset{}, fmt.Errorf("pre-apply error parsing selector %d '%s': %v", index+1, selector, err)
		}
		queries = append(queries, query)
	}
	return s.apply(s.db, queries, user)
}

func (s *ChangesetEventStore) QueryApplied(chartReleaseSelector string, offset int, limit int) ([]Changeset, error) {
	chartReleaseQuery, err := chartReleaseStore.selectorToQueryModel(s.db, chartReleaseSelector)
	if err != nil {
		return nil, err
	}
	chartRelease, err := chartReleaseStore.get(s.db, chartReleaseQuery)
	if err != nil {
		return nil, err
	}
	return s.internalChangesetEventStore.queryApplied(s.db, chartRelease.ID, offset, limit)
}

type internalChangesetEventStore struct {
	*internalModelStore[Changeset]
}

func (s *internalChangesetEventStore) plan(db *gorm.DB, changesets []Changeset, user *auth_models.User) ([]Changeset, error) {
	var ret []Changeset
	err := db.Transaction(func(tx *gorm.DB) error {
		for index, changeset := range changesets {
			chartRelease, err := chartReleaseStore.get(tx, ChartRelease{Model: gorm.Model{ID: changeset.ChartReleaseID}})
			if err != nil {
				return fmt.Errorf("plan error on %T %d: failed to get %T: %v", changeset, index+1, chartRelease, err)
			}
			if err = changeset.To.resolve(tx, Chart{Model: gorm.Model{ID: chartRelease.ChartID}}); err != nil {
				return fmt.Errorf("plan error on %T %d: failed to resolve 'to' version: %v", changeset, index+1, err)
			}
			if changeset.To.equalTo(changeset.From) {
				continue
			} else {
				planned, _, err := s.create(tx, changeset, user)
				if err != nil {
					return fmt.Errorf("plan error on %T %d: failed to create %T: %v", changeset, index+1, changeset, err)
				}
				ret = append(ret, planned)
			}
		}
		return nil
	})
	return ret, err
}

func (s *internalChangesetEventStore) apply(db *gorm.DB, changesets []Changeset, user *auth_models.User) ([]Changeset, error) {
	var ret []Changeset
	affectedChartReleases := make(map[uint]ChartRelease)
	err := db.Transaction(func(tx *gorm.DB) error {
		for index, changeset := range changesets {
			toApply, err := s.get(tx, changeset)
			if err != nil {
				return fmt.Errorf("apply error on %T %d: failed to query referenced %T: %v", changeset, index+1, changeset, err)
			}
			if toApply.AppliedAt != nil {
				return fmt.Errorf("(%s) apply validation error on %T %d (ID %d): this has already been applied", errors.BadRequest, changeset, index+1, toApply.ID)
			}
			if toApply.SupersededAt != nil {
				return fmt.Errorf("(%s) apply validation error on %T %d (ID: %d): this has been superseded by some other already-applied %T", errors.BadRequest, changeset, index+1, toApply.ID, changeset)
			}
			if toApply.To.ResolvedAt == nil {
				return fmt.Errorf("(%s) apply validation error on %T %d (ID: %d): the 'to' version appears not to have been internally resolved", errors.InternalServerError, changeset, index+1, toApply.ID)
			}
			chartRelease, err := chartReleaseStore.get(tx, ChartRelease{Model: gorm.Model{ID: toApply.ChartReleaseID}})
			if err != nil {
				return fmt.Errorf("apply error on %T %d (ID: %d): failed to get %T: %v", changeset, index+1, toApply.ID, chartRelease, err)
			}
			if _, alreadyAffected := affectedChartReleases[chartRelease.ID]; alreadyAffected {
				return fmt.Errorf("(%s) apply validation error: multiple changesets were against %T '%s'", errors.BadRequest, chartRelease, chartRelease.Name)
			} else {
				affectedChartReleases[chartRelease.ID] = chartRelease
			}
			if !chartRelease.ChartReleaseVersion.equalTo(toApply.From) {
				// We really shouldn't ever hit this case--when a Changeset is applied we mark all other as superseded,
				// which would've already been caught. In any case, we won't be applying this Changeset, so first mark it as
				// superseded:
				now := time.Now()
				_, err = s.edit(tx, Changeset{Model: gorm.Model{ID: toApply.ID}}, Changeset{SupersededAt: &now}, user, false)
				if err != nil {
					log.Error().Err(err).Msgf("couldn't retroactively mark Changeset %d as superseded", toApply.ID)
				}
				// Now try some hail-mary attempts to figure out how we got into this state:
				if chartRelease.ChartReleaseVersion.equalTo(toApply.To) {
					log.Error().Msgf("it appears that Changeset %d, or a copy of it, was already applied, but %d wasn't marked as applied or superseded", toApply.ID, toApply.ID)
				}
				if chartRelease.UpdatedAt.After(toApply.UpdatedAt) {
					log.Error().Msgf("it appears that ChartRelease %d was updated more recently than Changeset %d was, maybe a direct edit happened to the ChartRelease?", chartRelease.ID, toApply.ID)
				}
				// Return an error with errors.InternalServerError so it'll get logged
				return fmt.Errorf("(%s) apply validation error on %T %d (ID: %d): the %T was detected as being out-of-date before it could be applied--it has now been properly marked as superseded; please plan again or contact DevOps if the problem persists", errors.InternalServerError, changeset, index+1, toApply.ID, changeset.To)
			}
			// Update the struct fields of what came from the database
			chartRelease.ChartReleaseVersion = toApply.To
			// Now save what we have--*all* of it, including zero fields--back into the database
			chartRelease, err = chartReleaseStore.edit(tx, ChartRelease{Model: gorm.Model{ID: toApply.ChartReleaseID}}, chartRelease, user, true)
			if err != nil {
				return fmt.Errorf("apply error on %T %d (ID: %d): failed to modify %T (ID: %d): %v", changeset, index+1, toApply.ID, chartRelease, toApply.ChartReleaseID, err)
			}
			// Forcibly include AppliedAt and SupersededAt in the match criteria so we only find things where both of those
			// fields are empty.
			consumedChangesets, err := s.listAllMatchingByUpdated(tx, 0, Changeset{ChartReleaseID: toApply.ChartReleaseID}, "ChartReleaseID", "AppliedAt", "SupersededAt")
			if err != nil {
				return fmt.Errorf("post-apply error on %T %d (ID: %d): couldn't query consumed %T: %v", changeset, index+1, toApply.ID, changeset, err)
			}
			for _, consumedChangeset := range consumedChangesets {
				if consumedChangeset.ID == toApply.ID {
					applied, err := s.edit(tx, Changeset{Model: gorm.Model{ID: consumedChangeset.ID}}, Changeset{AppliedAt: &chartRelease.UpdatedAt}, user, false)
					if err != nil {
						return fmt.Errorf("post-apply error on %T %d (ID: %d): couldn't mark it as applied: %v", changeset, index+1, toApply.ID, err)
					}
					ret = append(ret, applied)
				} else {
					_, err := s.edit(tx, Changeset{Model: gorm.Model{ID: consumedChangeset.ID}}, Changeset{SupersededAt: &chartRelease.UpdatedAt}, user, false)
					if err != nil {
						return fmt.Errorf("post-apply error on %T %d (ID: %d): couldn't mark superseded %T (ID: %d) as superseded: %v", changeset, index+1, toApply.ID, changeset, consumedChangeset.ID, err)
					}
				}
			}
		}
		return nil
	})

	// If the update happened successfully, report these changes to any relevant Pagerduty integrations
	if err == nil {
		environmentReleases := make(map[uint][]string)

		for _, chartRelease := range affectedChartReleases {
			if chartRelease.EnvironmentID != nil {
				environmentReleases[*chartRelease.EnvironmentID] = append(environmentReleases[*chartRelease.EnvironmentID], chartRelease.Name)
			}
			if chartRelease.PagerdutyIntegration != nil && chartRelease.PagerdutyIntegration.Key != nil {
				go pagerduty.SendChangeSwallowErrors(
					*chartRelease.PagerdutyIntegration.Key,
					fmt.Sprintf("Version changes to %s via Sherlock/Beehive", chartRelease.Name),
					fmt.Sprintf(config.Config.MustString("beehive.chartReleaseUrlFormatString"), chartRelease.Name),
				)
			}
		}

		for environmentID, chartReleaseNames := range environmentReleases {
			environment, err := environmentStore.get(db, Environment{Model: gorm.Model{ID: environmentID}})
			if err == nil && environment.PagerdutyIntegration != nil && environment.PagerdutyIntegration.Key != nil {
				go pagerduty.SendChangeSwallowErrors(
					*environment.PagerdutyIntegration.Key,
					fmt.Sprintf("Version changes to %s via Sherlock/Beehive", strings.Join(chartReleaseNames, ", ")),
					fmt.Sprintf(config.Config.MustString("beehive.environmentUrlFormatString"), environment.Name),
				)
			}
		}
	}

	return ret, err
}

func (s *internalChangesetEventStore) queryApplied(db *gorm.DB, chartReleaseID uint, offset int, limit int) ([]Changeset, error) {
	ret := make([]Changeset, 0)
	chain := db.
		Unscoped().
		Where(&Changeset{ChartReleaseID: chartReleaseID}).
		Where("applied_at is not null").
		Order("applied_at desc").
		Preload(clause.Associations).
		Offset(offset)
	if limit > 0 {
		chain = chain.Limit(limit)
	}
	err := chain.Find(&ret).Error
	return ret, err
}
