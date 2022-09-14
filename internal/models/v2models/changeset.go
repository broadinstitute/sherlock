package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Changeset struct {
	gorm.Model
	ChartRelease   *ChartRelease
	ChartReleaseID uint

	From             ChartReleaseVersion `gorm:"embedded;embeddedPrefix:from_"`
	To               ChartReleaseVersion `gorm:"embedded;embeddedPrefix:to_"`
	AppliedAt        *time.Time
	SupersededAt     *time.Time
	NewAppVersions   []*AppVersion   `gorm:"many2many:v2_changeset_new_app_versions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	NewChartVersions []*ChartVersion `gorm:"many2many:v2_changeset_new_chart_versions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}

func (c Changeset) TableName() string {
	return "v2_changesets"
}

var changesetStore *internalChangesetEventStore

func init() {
	changesetStore = &internalChangesetEventStore{
		internalModelStore: &internalModelStore[Changeset]{
			selectorToQueryModel: changesetSelectorToQuery,
			modelToSelectors:     changesetToSelectors,
			validateModel:        validateChangeset,
			preCreate:            preCreateChangeset,
		},
	}
}

func changesetSelectorToQuery(_ *gorm.DB, selector string) (Changeset, error) {
	if len(selector) == 0 {
		return Changeset{}, fmt.Errorf("(%s) changeset selector cannot be empty", errors.BadRequest)
	}

	var query Changeset
	if isNumeric(selector) { // ID
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

func preCreateChangeset(db *gorm.DB, toCreate *Changeset, _ *auth.User) error {
	if toCreate != nil {

		// Resolve 'to' versions
		if toCreate.To.ResolvedAt == nil {
			chartRelease, err := chartReleaseStore.get(db, ChartRelease{Model: gorm.Model{ID: toCreate.ChartReleaseID}})
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
		if toCreate.From.AppVersionID != nil && toCreate.To.AppVersionID != nil && *toCreate.From.AppVersionID != *toCreate.To.AppVersionID {
			newVersion := toCreate.To.AppVersion
			for *toCreate.From.AppVersionID != newVersion.ID {
				if newVersion != nil {
					isActuallyNew := true
					for _, listedVersion := range toCreate.NewAppVersions {
						if listedVersion.ID == newVersion.ID {
							isActuallyNew = false
						}
					}
					if isActuallyNew {
						toCreate.NewAppVersions = append(toCreate.NewAppVersions, newVersion)
						if newVersion.ParentAppVersionID != nil {
							potentialNewVersion, err := appVersionStore.get(db, AppVersion{Model: gorm.Model{ID: *newVersion.ParentAppVersionID}})
							if err == nil {
								newVersion = &potentialNewVersion
							} else {
								break
							}
						} else {
							break
						}
					} else {
						break
					}
				}
			}
		}

		// List new chart versions
		if toCreate.From.ChartVersionID != nil && toCreate.To.ChartVersionID != nil && *toCreate.From.ChartVersionID != *toCreate.To.ChartVersionID {
			newVersion := toCreate.To.ChartVersion
			for *toCreate.From.ChartVersionID != newVersion.ID {
				if newVersion != nil {
					isActuallyNew := true
					for _, listedVersion := range toCreate.NewChartVersions {
						if listedVersion.ID == newVersion.ID {
							isActuallyNew = false
						}
					}
					if isActuallyNew {
						toCreate.NewChartVersions = append(toCreate.NewChartVersions, newVersion)
						if newVersion.ParentChartVersionID != nil {
							potentialNewVersion, err := chartVersionStore.get(db, ChartVersion{Model: gorm.Model{ID: *newVersion.ParentChartVersionID}})
							if err == nil {
								newVersion = &potentialNewVersion
							} else {
								break
							}
						} else {
							break
						}
					} else {
						break
					}
				}

			}
		}
	}
	return nil
}
