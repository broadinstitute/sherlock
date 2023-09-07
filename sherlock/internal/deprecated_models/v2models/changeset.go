package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/pactbroker"
	"github.com/broadinstitute/sherlock/sherlock/internal/pagerduty"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"strings"
	"time"
)

type Changeset struct {
	gorm.Model
	CiIdentifier   *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:changeset"`
	ChartRelease   *ChartRelease
	ChartReleaseID uint

	From             ChartReleaseVersion `gorm:"embedded;embeddedPrefix:from_"`
	To               ChartReleaseVersion `gorm:"embedded;embeddedPrefix:to_"`
	AppliedAt        *time.Time
	SupersededAt     *time.Time
	NewAppVersions   []*AppVersion   `gorm:"many2many:v2_changeset_new_app_versions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	NewChartVersions []*ChartVersion `gorm:"many2many:v2_changeset_new_chart_versions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`

	PlannedBy   *User
	PlannedByID *uint
	AppliedBy   *User
	AppliedByID *uint
}

func (c Changeset) TableName() string {
	return "v2_changesets"
}

func (c Changeset) getID() uint {
	return c.ID
}

func (c Changeset) GetCiIdentifier() *CiIdentifier {
	if c.CiIdentifier != nil {
		return c.CiIdentifier
	} else {
		return &CiIdentifier{ResourceType: "changeset", ResourceID: c.ID}
	}
}

type internalChangesetStore struct {
	*internalModelStore[Changeset]
}

var InternalChangesetStore *internalChangesetStore

func init() {
	InternalChangesetStore = &internalChangesetStore{
		internalModelStore: &internalModelStore[Changeset]{
			selectorToQueryModel: changesetSelectorToQuery,
			modelToSelectors:     changesetToSelectors,
			validateModel:        validateChangeset,
			preCreate:            preCreateChangeset,
			customCreationAssociationsClause: func(db *gorm.DB) *gorm.DB {
				return db.Omit("CiIdentifier", "ChartRelease", "PlannedBy", "AppliedBy")
			},
		},
	}
}

func changesetSelectorToQuery(_ *gorm.DB, selector string) (Changeset, error) {
	if len(selector) == 0 {
		return Changeset{}, fmt.Errorf("(%s) changeset selector cannot be empty", errors.BadRequest)
	}

	var query Changeset
	if utils.IsNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return Changeset{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	}
	return Changeset{}, fmt.Errorf("(%s) invalid changeset selector '%s'", errors.BadRequest, selector)
}

func changesetToSelectors(changeset *Changeset) []string {
	var selectors []string
	if changeset != nil && changeset.ID != 0 {
		selectors = append(selectors, fmt.Sprintf("%d", changeset.ID))
	}
	return selectors
}

func validateChangeset(changeset *Changeset) error {
	if changeset == nil {
		return fmt.Errorf("the model passed was nil")
	}
	// We intentionally don't validate the From because we don't want to re-resolve it, and maybe it became invalid due
	// to manual database changes.
	if err := changeset.To.validate(); err != nil {
		return fmt.Errorf("'to' %T on %T was invalid: %v", changeset.To, changeset, err)
	}
	return nil
}

func preCreateChangeset(db *gorm.DB, toCreate *Changeset, user *models.User) error {
	if toCreate != nil {
		if user != nil {
			if toCreate.PlannedByID != nil {
				log.Warn().Msgf("changeset being created already had a planned by ID (was %d, now setting for %d for %s)", *toCreate.PlannedByID, user.ID, user.Email)
			}
			toCreate.PlannedByID = &user.ID
		}

		// Resolve 'to' versions
		if toCreate.To.ResolvedAt == nil {
			chartRelease, err := InternalChartReleaseStore.Get(db, ChartRelease{Model: gorm.Model{ID: toCreate.ChartReleaseID}})
			if err != nil {
				return err
			}
			if !chartRelease.ChartReleaseVersion.equalTo(toCreate.From) {
				return fmt.Errorf("(%s) the request appears out of date, the recorded 'from' must match the target chart release", errors.BadRequest)
			}
			err = toCreate.To.resolve(db, Chart{Model: gorm.Model{ID: chartRelease.ChartID}})
			if err != nil {
				return err
			}
		}

		// List new app versions
		appVersionPath, _, err := InternalAppVersionStore.getChildrenPathToParent(db, toCreate.To.AppVersion, toCreate.From.AppVersionID)
		if err != nil {
			log.Error().Msgf("swallowing %T path error during changeset creation: %v", toCreate.To.AppVersion, err)
		} else {
			// This is really dumb, but we're going to go through the returned app versions and remove their relations so that Gorm doesn't try to upsert them
			// in a moment. I will try to document what I think is going on later but I need to stop the bleeding.
			for _, appVersion := range appVersionPath {
				appVersion.Chart = nil
				appVersion.CiIdentifier = nil
			}

			toCreate.NewAppVersions = appVersionPath
		}

		// List new chart versions
		chartVersionPath, _, err := InternalChartVersionStore.getChildrenPathToParent(db, toCreate.To.ChartVersion, toCreate.From.ChartVersionID)
		if err != nil {
			log.Error().Msgf("swallowing %T path error during changeset creation: %v", toCreate.From.ChartVersion, err)
		} else {
			// This is really dumb, but we're going to go through the returned app versions and remove their relations so that Gorm doesn't try to upsert them
			// in a moment. I will try to document what I think is going on later but I need to stop the bleeding.
			for _, chartVersion := range chartVersionPath {
				chartVersion.Chart = nil
				chartVersion.CiIdentifier = nil
			}

			toCreate.NewChartVersions = chartVersionPath
		}
	}
	return nil
}

func (s *internalChangesetStore) plan(db *gorm.DB, changesets []Changeset, user *models.User) ([]Changeset, error) {
	var ret []Changeset
	var err error
	for index, changeset := range changesets {
		var chartRelease ChartRelease
		chartRelease, err = InternalChartReleaseStore.Get(db, ChartRelease{Model: gorm.Model{ID: changeset.ChartReleaseID}})
		if err != nil {
			err = fmt.Errorf("plan error on %T %d: failed to get %T: %v", changeset, index+1, chartRelease, err)
			break
		}
		if err = changeset.To.resolve(db, Chart{Model: gorm.Model{ID: chartRelease.ChartID}}); err != nil {
			err = fmt.Errorf("plan error on %T %d: failed to resolve 'to' version: %v", changeset, index+1, err)
			break
		}
		if changeset.To.equalTo(changeset.From) {
			continue
		} else {
			var planned Changeset
			planned, _, err = s.Create(db, changeset, user)
			if err != nil && strings.Contains(err.Error(), "deadlock detected") {
				planned, _, err = s.Create(db, changeset, user)
				if err == nil {
					go slack.SendMessage(db.Statement.Context, "#ap-k8s-monitor", "Sherlock encountered a deadlock during changeset creation but recovered")
				}
			}
			if err != nil {
				err = fmt.Errorf("plan error on %T %d: failed to create %T: %v", changeset, index+1, changeset, err)
				break
			}
			ret = append(ret, planned)
		}
	}
	return ret, err
}

func (s *internalChangesetStore) apply(db *gorm.DB, changesets []Changeset, user *models.User) ([]Changeset, error) {
	var ret []Changeset
	affectedChartReleases := make(map[uint]ChartRelease)
	err := db.Transaction(func(tx *gorm.DB) error {
		for index, changeset := range changesets {
			toApply, err := s.Get(tx, changeset)
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
			chartRelease, err := InternalChartReleaseStore.Get(tx, ChartRelease{Model: gorm.Model{ID: toApply.ChartReleaseID}})
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
				_, err = s.Edit(tx, Changeset{Model: gorm.Model{ID: toApply.ID}}, Changeset{SupersededAt: &now}, user, false)
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
			chartRelease, err = InternalChartReleaseStore.Edit(tx, ChartRelease{Model: gorm.Model{ID: toApply.ChartReleaseID}}, chartRelease, user, true)
			if err != nil {
				return fmt.Errorf("apply error on %T %d (ID: %d): failed to modify %T (ID: %d): %v", changeset, index+1, toApply.ID, chartRelease, toApply.ChartReleaseID, err)
			}
			// Forcibly include AppliedAt and SupersededAt in the match criteria so we only find things where both of those
			// fields are empty.
			consumedChangesets, err := s.ListAllMatchingByUpdated(tx, 0, Changeset{ChartReleaseID: toApply.ChartReleaseID}, "ChartReleaseID", "AppliedAt", "SupersededAt")
			if err != nil {
				return fmt.Errorf("post-apply error on %T %d (ID: %d): couldn't query consumed %T: %v", changeset, index+1, toApply.ID, changeset, err)
			}
			for _, consumedChangeset := range consumedChangesets {
				if consumedChangeset.ID == toApply.ID {
					applied, err := s.Edit(tx, Changeset{Model: gorm.Model{ID: consumedChangeset.ID}}, Changeset{AppliedAt: &chartRelease.UpdatedAt, AppliedByID: &user.ID}, user, false)
					if err != nil {
						return fmt.Errorf("post-apply error on %T %d (ID: %d): couldn't mark it as applied: %v", changeset, index+1, toApply.ID, err)
					}
					ret = append(ret, applied)
				} else {
					_, err := s.Edit(tx, Changeset{Model: gorm.Model{ID: consumedChangeset.ID}}, Changeset{SupersededAt: &chartRelease.UpdatedAt}, user, false)
					if err != nil {
						return fmt.Errorf("post-apply error on %T %d (ID: %d): couldn't mark superseded %T (ID: %d) as superseded: %v", changeset, index+1, toApply.ID, changeset, consumedChangeset.ID, err)
					}
				}
			}
		}
		return nil
	})

	// If the update happened successfully, report these changes to any relevant Pagerduty and Pact integrations
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
			// Record app version to Pact broker
			if chartRelease.Environment != nil && chartRelease.Chart.PactParticipant != nil && *chartRelease.Chart.PactParticipant && chartRelease.Environment.PactIdentifier != nil {
				go pactbroker.RecordDeployment(
					chartRelease.Chart.Name,
					chartRelease.AppVersion.AppVersion,
					*chartRelease.Environment.PactIdentifier,
				)
			}
		}

		for environmentID, chartReleaseNames := range environmentReleases {
			environment, err := InternalEnvironmentStore.Get(db, Environment{Model: gorm.Model{ID: environmentID}})
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

func (s *internalChangesetStore) queryAppliedForChartRelease(db *gorm.DB, chartReleaseID uint, offset int, limit int) ([]Changeset, error) {
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

func (s *internalChangesetStore) queryAppliedForVersion(db *gorm.DB, chartID uint, version string, versionType string) ([]Changeset, error) {
	var ret []Changeset
	var changesetIDs []uint
	switch versionType {

	case "app":
		if appVersionIfExists, err := InternalAppVersionStore.GetIfExists(db, AppVersion{AppVersion: version, ChartID: chartID}); err != nil {
			return ret, err
		} else if appVersionIfExists != nil {
			// If we have record of the version, it would show up in the "new versions" introduced by a changeset.
			// Use that to match so we have a shot at catching deployments where this version was an intermediary.
			if err = db.Raw(`
select v2_changesets.id
from v2_changesets
         -- Join through changelog table to the app version we care about
         inner join v2_changeset_new_app_versions
                    on v2_changeset_new_app_versions.changeset_id = v2_changesets.id
                        and v2_changeset_new_app_versions.app_version_id = ?

         -- Join through to chart releases to filter out deleted/irrelevant ones
         inner join v2_chart_releases
                    on v2_chart_releases.id = v2_changesets.chart_release_id
                        and v2_chart_releases.deleted_at is null
                        and v2_chart_releases.chart_id = ?

-- Filter to changes that actually got applied
where v2_changesets.applied_at is not null

order by v2_changesets.applied_at desc`, appVersionIfExists.ID, chartID).Scan(&changesetIDs).Error; err != nil {
				return ret, err
			}
		} else {
			// If we have no record of the version at all, it might just be a custom string that was entered
			// into the API. All we can do is match based on changesets that specifically deployed that version.
			if err = db.Raw(`
select v2_changesets.id
from v2_changesets
         -- Join through to chart releases to filter out deleted/irrelevant ones
         inner join v2_chart_releases
                    on v2_chart_releases.id = v2_changesets.chart_release_id
                        and v2_chart_releases.deleted_at is null
                        and v2_chart_releases.chart_id = ?

-- Filter to changes that actually got applied and target our version
where v2_changesets.applied_at is not null
  and v2_changesets.to_app_version_exact = ?

order by v2_changesets.applied_at desc
`, chartID, version).Error; err != nil {
				return ret, err
			}
		}

	case "chart":
		if chartVersionIfExists, err := InternalChartVersionStore.GetIfExists(db, ChartVersion{ChartVersion: version, ChartID: chartID}); err != nil {
			return ret, err
		} else if chartVersionIfExists != nil {
			// If we have record of the version, it would show up in the "new versions" introduced by a changeset.
			// Use that to match so we have a shot at catching deployments where this version was an intermediary.
			if err = db.Raw(`
select v2_changesets.id
from v2_changesets
         -- Join through changelog table to the chart version we care about
         inner join v2_changeset_new_chart_versions
                    on v2_changeset_new_chart_versions.changeset_id = v2_changesets.id
                        and v2_changeset_new_chart_versions.chart_version_id = ?

         -- Join through to chart releases to filter out deleted/irrelevant ones
         inner join v2_chart_releases
                    on v2_chart_releases.id = v2_changesets.chart_release_id
                        and v2_chart_releases.deleted_at is null
                        and v2_chart_releases.chart_id = ?

-- Filter to changes that actually got applied
where v2_changesets.applied_at is not null

order by v2_changesets.applied_at desc`, chartVersionIfExists.ID, chartID).Scan(&changesetIDs).Error; err != nil {
				return ret, err
			}
		} else {
			// If we have no record of the version at all, it might just be a custom string that was entered
			// into the API. All we can do is match based on changesets that specifically deployed that version.
			if err = db.Raw(`
select v2_changesets.id
from v2_changesets
         -- Join through to chart releases to filter out deleted/irrelevant ones
         inner join v2_chart_releases
                    on v2_chart_releases.id = v2_changesets.chart_release_id
                        and v2_chart_releases.deleted_at is null
                        and v2_chart_releases.chart_id = ?

-- Filter to changes that actually got applied and target our version
where v2_changesets.applied_at is not null
  and v2_changesets.to_chart_version_exact = ?

order by v2_changesets.applied_at desc
`, chartID, version).Error; err != nil {
				return ret, err
			}
		}

	default:
		return ret, fmt.Errorf("(%s) need to specify 'app' or 'chart' for version type, instead got '%s'", errors.BadRequest, versionType)
	}

	// Now we do a normal query to load more than just the ID out of the database.
	ret = make([]Changeset, len(changesetIDs))
	var err error
	for index, id := range changesetIDs {
		if ret[index], err = InternalChangesetStore.Get(db, Changeset{Model: gorm.Model{ID: id}}); err != nil {
			return ret, err
		}
	}
	return ret, nil
}
