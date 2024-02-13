package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"slices"
	"strings"
)

type ChangesetV3PlanRequest struct {
	ChartReleases      []ChangesetV3PlanRequestChartReleaseEntry `json:"chartReleases"`
	Environments       []ChangesetV3PlanRequestEnvironmentEntry  `json:"environments"`
	RecreateChangesets []uint                                    `json:"recreateChangesets"`
}

type ChangesetV3PlanRequestChartReleaseEntry struct {
	ChangesetV3Create
	UseExactVersionsFromOtherChartRelease *string `json:"useExactVersionsFromOtherChartRelease"`
	FollowVersionsFromOtherChartRelease   *string `json:"followVersionsFromOtherChartRelease"`
}

type ChangesetV3PlanRequestEnvironmentEntry struct {
	Environment                          string
	UseExactVersionsFromOtherEnvironment *string  `json:"useExactVersionsFromOtherEnvironment"`
	FollowVersionsFromOtherEnvironment   *string  `json:"followVersionsFromOtherEnvironment"`
	IncludeCharts                        []string `json:"includeCharts"` // If omitted, will include all chart releases that haven't opted out of bulk updates
	ExcludeCharts                        []string `json:"excludeCharts"`
}

// changesetsProceduresV3Plan godoc
//
//	@summary		Plan--but do not apply--version changes to Chart Releases
//	@description	Refreshes and calculates version diffs for Chart Releases. If there's a diff, the plan is stored and returned so it can be applied later.
//	@tags			Changesets
//	@accept			json
//	@produce		json
//	@param			changeset-plan-request	body		ChangesetV3PlanRequest	true	"Info on what version changes or refreshes to plan"
//	@param			verbose-output			query		bool					false	"If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned."
//	@success		200,201					{array}		ChangesetV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/changesets/procedures/v3/plan [post]
func changesetsProceduresV3Plan(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body ChangesetV3PlanRequest
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) JSON error parsing body to %T: %w", errors.BadRequest, body, err))
		return
	}

	var verboseOutput bool
	if verboseOutputString := ctx.DefaultQuery("verbose-output", "true"); verboseOutputString == "true" {
		verboseOutput = true
	} else if verboseOutputString == "false" {
		verboseOutput = false
	} else {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) couldn't parse 'verbose-output' to a boolean: %s", errors.BadRequest, verboseOutputString))
		return
	}

	var chartReleaseChangesets, environmentChangesets, recreateChangesets []models.Changeset

	if chartReleaseChangesets, err = body.parseChartReleaseEntries(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error handling chart release entries: %w", err))
		return
	}
	if environmentChangesets, err = body.parseEnvironmentEntries(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error handling environment entries: %w", err))
		return
	}
	if recreateChangesets, err = body.parseRecreateEntries(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error handling recreate entries: %w", err))
		return
	}
	changesets := append(append(chartReleaseChangesets, environmentChangesets...), recreateChangesets...)

	var ret []models.Changeset
	var createdChangesetIDs []uint
	if createdChangesetIDs, err = models.PlanChangesets(db, changesets); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error planning changesets: %w", err))
		return
	} else if len(createdChangesetIDs) == 0 {
		ctx.JSON(http.StatusOK, []ChangesetV3{})
		return
	} else if verboseOutput {
		if err = db.Scopes(models.ReadChangesetScope).Find(&ret, createdChangesetIDs).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error querying planned changesets: %w", err))
			return
		} else {
			slices.SortFunc(ret, func(a, b models.Changeset) int {
				if a.ChartRelease != nil || b.ChartRelease != nil {
					return strings.Compare(a.ChartRelease.Name, b.ChartRelease.Name)
				} else {
					return 0
				}
			})
			ctx.JSON(http.StatusOK, utils.Map(ret, changesetFromModel))
		}
	} else {
		ctx.JSON(http.StatusOK, utils.Map(createdChangesetIDs, func(id uint) ChangesetV3 {
			return ChangesetV3{CommonFields: CommonFields{ID: id}}
		}))
	}
}

func (r *ChangesetV3PlanRequest) parseChartReleaseEntries(db *gorm.DB) (changesets []models.Changeset, err error) {
	changesets = make([]models.Changeset, 0, len(r.ChartReleases))
	for index, chartReleaseRequestEntry := range r.ChartReleases {
		var changeset models.Changeset
		if changeset, err = chartReleaseRequestEntry.toModel(db); err != nil {
			return nil, fmt.Errorf("error converting chart release request entry %d to model: %w", index, err)

		}
		if chartReleaseRequestEntry.UseExactVersionsFromOtherChartRelease != nil || chartReleaseRequestEntry.FollowVersionsFromOtherChartRelease != nil {
			if chartReleaseRequestEntry.UseExactVersionsFromOtherChartRelease != nil && chartReleaseRequestEntry.FollowVersionsFromOtherChartRelease != nil {
				return nil, fmt.Errorf("(%s) can't specify both useExactVersionsFromOtherChartRelease and followVersionsFromOtherChartRelease", errors.BadRequest)
			}
			var otherChartRelease models.ChartRelease
			if chartReleaseRequestEntry.UseExactVersionsFromOtherChartRelease != nil {
				if otherChartRelease, err = chartReleaseModelFromSelector(db, *chartReleaseRequestEntry.UseExactVersionsFromOtherChartRelease); err != nil {
					return nil, fmt.Errorf("error parsing chart release selector '%s' for useExactVersionsFromOtherChartRelease: %w", *chartReleaseRequestEntry.UseExactVersionsFromOtherChartRelease, err)
				} else if err = db.Where(&otherChartRelease).Take(&otherChartRelease).Error; err != nil {
					return nil, fmt.Errorf("error querying chart release '%s' for useExactVersionsFromOtherChartRelease: %w", *chartReleaseRequestEntry.UseExactVersionsFromOtherChartRelease, err)
				}
			} else if chartReleaseRequestEntry.FollowVersionsFromOtherChartRelease != nil {
				if otherChartRelease, err = chartReleaseModelFromSelector(db, *chartReleaseRequestEntry.FollowVersionsFromOtherChartRelease); err != nil {
					return nil, fmt.Errorf("error parsing chart release selector '%s' for followVersionsFromOtherChartRelease: %w", *chartReleaseRequestEntry.FollowVersionsFromOtherChartRelease, err)
				} else if err = db.Where(&otherChartRelease).Take(&otherChartRelease).Error; err != nil {
					return nil, fmt.Errorf("error querying chart release '%s' for followVersionsFromOtherChartRelease: %w", *chartReleaseRequestEntry.FollowVersionsFromOtherChartRelease, err)
				}
			}

			changeset.To = otherChartRelease.ChartReleaseVersion
			if chartReleaseRequestEntry.UseExactVersionsFromOtherChartRelease != nil {
				changeset.To.ChartVersionResolver = utils.PointerTo("exact")
				if otherChartRelease.AppVersionResolver != nil && *otherChartRelease.AppVersionResolver == "none" {
					changeset.To.AppVersionResolver = utils.PointerTo("none")
				} else {
					changeset.To.AppVersionResolver = utils.PointerTo("exact")
				}
			} else if chartReleaseRequestEntry.FollowVersionsFromOtherChartRelease != nil {
				changeset.To.AppVersionResolver = utils.PointerTo("follow")
				changeset.To.AppVersionFollowChartReleaseID = &otherChartRelease.ID
				changeset.To.ChartVersionResolver = utils.PointerTo("follow")
				changeset.To.ChartVersionFollowChartReleaseID = &otherChartRelease.ID
			}
		}
		changesets = append(changesets, changeset)
	}
	return changesets, nil
}

func (r *ChangesetV3PlanRequest) parseEnvironmentEntries(db *gorm.DB) (changesets []models.Changeset, err error) {
	changesets = make([]models.Changeset, 0)
	for _, environmentRequestEntry := range r.Environments {
		var environment models.Environment
		if environment, err = environmentModelFromSelector(environmentRequestEntry.Environment); err != nil {
			return nil, fmt.Errorf("error parsing environment selector '%s' for environment: %w", environmentRequestEntry.Environment, err)
		} else if err = db.Where(&environment).Select("id").Take(&environment).Error; err != nil {
			return nil, fmt.Errorf("error querying environment '%s': %w", environmentRequestEntry.Environment, err)
		}
		chartsToExplicitlyInclude := make(map[uint]struct{})
		for _, chartName := range environmentRequestEntry.IncludeCharts {
			var chart models.Chart
			if chart, err = chartModelFromSelector(chartName); err != nil {
				return nil, fmt.Errorf("error parsing chart selector '%s' for includeCharts: %w", chartName, err)
			} else if err = db.Where(&chart).Select("id").Take(&chart).Error; err != nil {
				return nil, fmt.Errorf("error querying chart '%s' for includeCharts: %w", chartName, err)
			}
			chartsToExplicitlyInclude[chart.ID] = struct{}{}
		}
		chartsToExplicitlyExclude := make(map[uint]struct{})
		for _, chartName := range environmentRequestEntry.ExcludeCharts {
			var chart models.Chart
			if chart, err = chartModelFromSelector(chartName); err != nil {
				return nil, fmt.Errorf("error parsing chart selector '%s' for excludeCharts: %w", chartName, err)
			} else if err = db.Where(&chart).Select("id").Take(&chart).Error; err != nil {
				return nil, fmt.Errorf("error querying chart '%s' for excludeCharts: %w", chartName, err)
			}
			chartsToExplicitlyExclude[chart.ID] = struct{}{}
		}
		var chartReleasesInEnvironment []models.ChartRelease
		if err = db.Where(&models.ChartRelease{EnvironmentID: &environment.ID}).Find(&chartReleasesInEnvironment).Error; err != nil {
			return nil, fmt.Errorf("error querying chart releases in environment '%s': %w", environmentRequestEntry.Environment, err)
		}
		var targetedChartReleases []models.ChartRelease
		for _, potentialTargetChartRelease := range chartReleasesInEnvironment {
			_, explicitlyIncluded := chartsToExplicitlyInclude[potentialTargetChartRelease.ChartID]
			_, explicitlyExcluded := chartsToExplicitlyExclude[potentialTargetChartRelease.ChartID]
			defaultIncluded := potentialTargetChartRelease.IncludeInBulkChangesets == nil || *potentialTargetChartRelease.IncludeInBulkChangesets
			// Explicitly included is always included.
			// Otherwise, it's included if it is included by default and isn't explicitly excluded.
			if explicitlyIncluded || (!explicitlyExcluded && defaultIncluded) {
				targetedChartReleases = append(targetedChartReleases, potentialTargetChartRelease)
			}
		}
		chartReleasesFromOtherEnvironment := make(map[uint]models.ChartRelease)
		if environmentRequestEntry.UseExactVersionsFromOtherEnvironment != nil || environmentRequestEntry.FollowVersionsFromOtherEnvironment != nil {
			if environmentRequestEntry.UseExactVersionsFromOtherEnvironment != nil && environmentRequestEntry.FollowVersionsFromOtherEnvironment != nil {
				return nil, fmt.Errorf("(%s) can't specify both useExactVersionsFromOtherEnvironment and followVersionsFromOtherEnvironment", errors.BadRequest)
			}
			var otherEnvironment models.Environment
			if environmentRequestEntry.UseExactVersionsFromOtherEnvironment != nil {
				if otherEnvironment, err = environmentModelFromSelector(*environmentRequestEntry.UseExactVersionsFromOtherEnvironment); err != nil {
					return nil, fmt.Errorf("error parsing environment selector '%s' for useExactVersionsFromOtherEnvironment: %w", *environmentRequestEntry.UseExactVersionsFromOtherEnvironment, err)
				} else if err = db.Where(&otherEnvironment).Select("id").Take(&otherEnvironment).Error; err != nil {
					return nil, fmt.Errorf("error querying environment '%s': %w", *environmentRequestEntry.UseExactVersionsFromOtherEnvironment, err)
				}
			} else {
				if otherEnvironment, err = environmentModelFromSelector(*environmentRequestEntry.FollowVersionsFromOtherEnvironment); err != nil {
					return nil, fmt.Errorf("error parsing environment selector '%s' for followVersionsFromOtherEnvironment: %w", *environmentRequestEntry.FollowVersionsFromOtherEnvironment, err)
				} else if err = db.Where(&otherEnvironment).Select("id").Take(&otherEnvironment).Error; err != nil {
					return nil, fmt.Errorf("error querying environment '%s': %w", *environmentRequestEntry.FollowVersionsFromOtherEnvironment, err)
				}
			}
			var chartReleasesFromOtherEnvironmentList []models.ChartRelease
			if err = db.Where(&models.ChartRelease{EnvironmentID: &otherEnvironment.ID}).Find(&chartReleasesFromOtherEnvironmentList).Error; err != nil {
				return nil, fmt.Errorf("error querying chart releases in environment '%s': %w", otherEnvironment.Name, err)
			}
			for _, chartRelease := range chartReleasesFromOtherEnvironmentList {
				chartReleasesFromOtherEnvironment[chartRelease.ChartID] = chartRelease
			}
		}
		for _, targetedChartRelease := range targetedChartReleases {
			var changeset models.Changeset
			changeset.ChartReleaseID = targetedChartRelease.ID
			if otherChartRelease, present := chartReleasesFromOtherEnvironment[targetedChartRelease.ChartID]; present {
				changeset.To = otherChartRelease.ChartReleaseVersion
				if environmentRequestEntry.UseExactVersionsFromOtherEnvironment != nil {
					changeset.To.ChartVersionResolver = utils.PointerTo("exact")
					if otherChartRelease.AppVersionResolver != nil && *otherChartRelease.AppVersionResolver == "none" {
						changeset.To.AppVersionResolver = utils.PointerTo("none")
					} else {
						changeset.To.AppVersionResolver = utils.PointerTo("exact")
					}
				} else if environmentRequestEntry.FollowVersionsFromOtherEnvironment != nil {
					changeset.To.AppVersionResolver = utils.PointerTo("follow")
					changeset.To.AppVersionFollowChartReleaseID = &otherChartRelease.ID
					changeset.To.ChartVersionResolver = utils.PointerTo("follow")
					changeset.To.ChartVersionFollowChartReleaseID = &otherChartRelease.ID
				}
			}
			changesets = append(changesets, changeset)
		}
	}
	return changesets, nil
}

func (r *ChangesetV3PlanRequest) parseRecreateEntries(db *gorm.DB) (changesets []models.Changeset, err error) {
	changesets = make([]models.Changeset, 0, len(r.RecreateChangesets))
	for _, existingChangesetID := range r.RecreateChangesets {
		var existingChangeset, newChangeset models.Changeset
		if existingChangesetID <= 0 {
			return nil, fmt.Errorf("(%s) invalid existing changeset ID: %d", errors.BadRequest, existingChangesetID)
		}
		if err = db.Take(&existingChangeset, existingChangesetID).Error; err != nil {
			return nil, fmt.Errorf("error querying changeset %d: %w", existingChangesetID, err)
		}
		newChangeset.To = existingChangeset.To
		newChangeset.ChartReleaseID = existingChangeset.ChartReleaseID
		// We force the resolvers to "exact" so that we resolve to the exact versions from this prior changeset,
		// even if a branch or something has updated since then.
		newChangeset.To.AppVersionResolver = utils.PointerTo("exact")
		newChangeset.To.ChartVersionResolver = utils.PointerTo("exact")
		changesets = append(changesets, newChangeset)
	}
	return changesets, nil
}
