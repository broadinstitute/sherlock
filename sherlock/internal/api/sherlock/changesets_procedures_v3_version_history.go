package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// changesetsProceduresV3VersionHistory godoc
//
//	@summary		List applied Changesets for an App or Chart Version
//	@description	List existing applied Changesets that newly deployed a given App Version or Chart Version, ordered by most recently applied.
//	@tags			Changesets
//	@produce		json
//	@param			version-type			path		string	true	"The type of the version, either 'app' or 'chart'"	Enums(app, chart)
//	@param			chart					path		string	true	"The chart the version belongs to"
//	@param			version					path		string	true	"The version to look for"
//	@success		200						{array}		ChangesetV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/changesets/procedures/v3/version-history/{version-type}/{chart}/{version} [get]
func changesetsProceduresV3VersionHistory(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	// get chart ID
	var chartQuery models.Chart
	if chartQuery, err = chartModelFromSelector(canonicalizeSelector(ctx.Param("chart"))); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var chart models.Chart
	if err = db.Select("id").Where(&chartQuery).First(&chart).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error querying chart '%s': %w", ctx.Param("chart"), err))
		return
	}

	// version string
	version := ctx.Param("version")
	if version == "" {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) version cannot be empty", errors.BadRequest))
		return
	}

	var changesetIDs []uint
	switch ctx.Param("version-type") {
	case "app":
		// check if this app version exists in our database
		var matchingAppVersion []models.AppVersion
		if err = db.Where(&models.AppVersion{ChartID: chart.ID, AppVersion: version}).Limit(1).Find(&matchingAppVersion).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error querying app version '%s': %w", version, err))
			return
		} else if len(matchingAppVersion) > 0 {
			// It was in our database, so scan for entries in the changelog table so we handle intermediary versions
			if err = db.Raw(`
select changesets.id
from changesets
         -- Join through changelog table to the app version we care about
         inner join changeset_new_app_versions
                    on changeset_new_app_versions.changeset_id = changesets.id
                        and changeset_new_app_versions.app_version_id = ?

         -- Join through to chart releases to filter out deleted/irrelevant ones
         inner join chart_releases
                    on chart_releases.id = changesets.chart_release_id
                        and chart_releases.deleted_at is null
                        and chart_releases.chart_id = ?

-- Filter to changes that actually got applied
where changesets.applied_at is not null

order by changesets.applied_at desc`, matchingAppVersion[0].ID, chart.ID).Scan(&changesetIDs).Error; err != nil {
				errors.AbortRequest(ctx, fmt.Errorf("error querying applied changesets for intermediary app version '%s': %w", version, err))
				return
			}
		} else {
			// If we have no record of the version, best we can do is match changesets that specifically deployed that
			// version
			if err = db.Raw(`
select changesets.id
from changesets
         -- Join through to chart releases to filter out deleted/irrelevant ones
         inner join chart_releases
                    on chart_releases.id = changesets.chart_release_id
                        and chart_releases.deleted_at is null
                        and chart_releases.chart_id = ?

-- Filter to changes that actually got applied and target our version
where changesets.applied_at is not null
  and changesets.to_app_version_exact = ?

order by changesets.applied_at desc
`, chart.ID, version).Scan(&changesetIDs).Error; err != nil {
				errors.AbortRequest(ctx, fmt.Errorf("error querying applied changesets for app version '%s': %w", version, err))
			}
		}
	case "chart":
		// check if this app version exists in our database
		var matchingChartVersion []models.ChartVersion
		if err = db.Where(&models.ChartVersion{ChartID: chart.ID, ChartVersion: version}).Limit(1).Find(&matchingChartVersion).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error querying chart version '%s': %w", version, err))
			return
		} else if len(matchingChartVersion) > 0 {
			// It was in our database, so scan for entries in the changelog table so we handle intermediary versions
			if err = db.Raw(`
select changesets.id
from changesets
         -- Join through changelog table to the chart version we care about
         inner join changeset_new_chart_versions
                    on changeset_new_chart_versions.changeset_id = changesets.id
                        and changeset_new_chart_versions.chart_version_id = ?

         -- Join through to chart releases to filter out deleted/irrelevant ones
         inner join chart_releases
                    on chart_releases.id = changesets.chart_release_id
                        and chart_releases.deleted_at is null
                        and chart_releases.chart_id = ?

-- Filter to changes that actually got applied
where changesets.applied_at is not null

order by changesets.applied_at desc`, matchingChartVersion[0].ID, chart.ID).Scan(&changesetIDs).Error; err != nil {
				errors.AbortRequest(ctx, fmt.Errorf("error querying applied changesets for intermediary chart version '%s': %w", version, err))
				return
			}
		} else {
			// If we have no record of the version, best we can do is match changesets that specifically deployed that
			// version
			if err = db.Raw(`
select changesets.id
from changesets
         -- Join through to chart releases to filter out deleted/irrelevant ones
         inner join chart_releases
                    on chart_releases.id = changesets.chart_release_id
                        and chart_releases.deleted_at is null
                        and chart_releases.chart_id = ?

-- Filter to changes that actually got applied and target our version
where changesets.applied_at is not null
  and changesets.to_chart_version_exact = ?

order by changesets.applied_at desc
`, chart.ID, version).Scan(&changesetIDs).Error; err != nil {
				errors.AbortRequest(ctx, fmt.Errorf("error querying applied changesets for chart version '%s': %w", version, err))
			}
		}
	default:
		errors.AbortRequest(ctx, fmt.Errorf("(%s) invalid version type '%s'", errors.BadRequest, ctx.Param("version-type")))
		return
	}

	var ret []models.Changeset
	if len(changesetIDs) > 0 {
		if err = db.Scopes(models.ReadChangesetScope).Order("applied_at desc").Find(&ret, changesetIDs).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error querying changesets to return: %w", err))
			return
		}
	}
	ctx.JSON(http.StatusOK, utils.Map(ret, changesetFromModel))
}
