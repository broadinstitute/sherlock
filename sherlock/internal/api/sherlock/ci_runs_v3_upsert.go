package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/ci_hooks"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/gha_oidc/gha_oidc_claims"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"
	"gorm.io/gorm/clause"
	"net/http"
)

type CiRunV3Upsert struct {
	ciRunFields
	Charts        []string `json:"charts"`        // Always appends; will eliminate duplicates.
	ChartVersions []string `json:"chartVersions"` // Always appends; will eliminate duplicates.
	AppVersions   []string `json:"appVersions"`   // Always appends; will eliminate duplicates.
	Clusters      []string `json:"clusters"`      // Always appends; will eliminate duplicates. Spreads to contained chart releases and their environments.
	Environments  []string `json:"environments"`  // Always appends; will eliminate duplicates. Spreads to contained chart releases and their clusters.
	ChartReleases []string `json:"chartReleases"` // Always appends; will eliminate duplicates. Spreads to associated environments and clusters.
	Changesets    []string `json:"changesets"`    // Always appends; will eliminate duplicates. Spreads to associated chart releases, environments, and clusters.

	// Keys treated like chartReleases. Values set resource-specific statuses for chart releases and associated changesets, new app versions, and new chart versions.
	ChartReleaseStatuses map[string]string `json:"chartReleaseStatuses"`

	// Makes entries in the changesets field also spread to new app versions and chart versions deployed by the changeset. 'when-static' is the default and does this spreading only when the chart release is in a static environment.
	RelateToChangesetNewVersions string `json:"relateToChangesetNewVersions" enums:"always,when-static,never" default:"when-static" binding:"oneof=always when-static never ''"`
	// If set to true, errors handling selectors for relations should be ignored. Normally, passing an unknown chart, cluster, etc. will abort the request, but they won't if this is true.
	IgnoreBadSelectors bool `json:"ignoreBadSelectors" default:"false"`
}

// ciRunsV3Upsert godoc
//
//	@summary		Create or update a CiRun
//	@description	Create or update a CiRun with timing, status, and related resource information. This endpoint is idempotent.
//	@description	The fields for clusters, charts, chart releases, environments, etc. all accept selectors, and they will
//	@description	be smart about "spreading" to indirect relations. More info is available on the CiRunV3Upsert data type,
//	@description	but the gist is that specifying a changeset implies its chart release (and optionally app/chart versions),
//	@description	specifying or implying a chart release implies its environment/cluster, and specifying an environment/cluster
//	@description	implies all chart releases they contain.
//	@tags			CiRuns
//	@accept			json
//	@produce		json
//	@param			ciRun					body		CiRunV3Upsert	true	"The CiRun to upsert"
//	@success		201						{object}	CiRunV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/ci-runs/v3 [put]
func ciRunsV3Upsert(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var body CiRunV3Upsert
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	// The notifySlackCustomIcon has some special handling noted in the API docs -- if it's empty, we ignore it here,
	// so that GitHub Actions etc. can just always send the field and don't have to worry about doing conditional
	// stuff to omit the field entirely.
	if body.NotifySlackCustomIcon != nil && *body.NotifySlackCustomIcon == "" {
		body.NotifySlackCustomIcon = nil
	}

	// Opportunistically fill empty fields with information passed in the GHA OIDC JWT
	if body.Platform == "" || body.Platform == "github-actions" {
		var claims *gha_oidc_claims.Claims
		if claims, err = authentication.ShouldUseGithubClaims(ctx); err == nil {
			body.Platform = "github-actions"
			if body.GithubActionsOwner == "" {
				body.GithubActionsOwner = claims.RepositoryOwner
			}
			if body.GithubActionsRepo == "" {
				body.GithubActionsRepo = claims.TrimmedRepositoryName()
			}
			if body.GithubActionsRunID == 0 {
				body.GithubActionsRunID, _ = utils.ParseUint(claims.RunID)
			}
			if body.GithubActionsAttemptNumber == 0 {
				body.GithubActionsAttemptNumber, _ = utils.ParseUint(claims.RunAttempt)
			}
			if body.GithubActionsWorkflowPath == "" {
				body.GithubActionsWorkflowPath = claims.TrimmedWorkflowPath()
			}
		}
	}

	// If this is a retry, we can smooth over a UX snag by grabbing previous runs' "notify upon retry" channels.
	// The snag is that some folks use "rerun failed jobs" in GitHub Actions, which *wouldn't* rerun the part that
	// tells Sherlock which channels to notify for that run. It's possible that different parts would've run on
	// different previous runs, so we just grab data on all of them.
	if body.Platform == "github-actions" &&
		body.GithubActionsOwner != "" &&
		body.GithubActionsRepo != "" &&
		body.GithubActionsRunID != 0 &&
		body.GithubActionsAttemptNumber > 1 {

		var previousRuns []models.CiRun
		if recoverableErr := db.
			Model(&models.CiRun{}).
			// Filter for the same run ID, which are unique per owner and repo
			Where(&models.CiRun{
				Platform:           "github-actions",
				GithubActionsOwner: body.GithubActionsOwner,
				GithubActionsRepo:  body.GithubActionsRepo,
				GithubActionsRunID: body.GithubActionsRunID,
			}).
			// Filter for an attempt number less than ours and non-empty notify upon retry channels
			Where("github_actions_attempt_number < ? AND notify_slack_channels_upon_retry IS NOT NULL", body.GithubActionsAttemptNumber).
			// Limit to the number of previous attempts there could've been, might as well get the speed boost
			Limit(max(1, int(body.GithubActionsAttemptNumber)-1)).
			Find(&previousRuns).Error; recoverableErr != nil {

			// If there was an error, just Slack it, no need to blow up
			slack.ReportError(ctx, "error fetching previous runs for notify upon retry", recoverableErr)

		} else if len(previousRuns) > 0 {

			// If we got previous runs matching our criteria, add all their retry channels to our current body.
			// We dedupe later (for all fields of the body) so we don't worry about that here.
			for _, previousRun := range previousRuns {
				body.NotifySlackChannelsUponRetry = append(body.NotifySlackChannelsUponRetry, previousRun.NotifySlackChannelsUponRetry...)
			}

		}

	}

	// We want to handle the "spreading" mechanic that some of the fields have. To do that, we'll literally just re-assemble
	// the body we got into one post-spread. Then we'll handle that body and de-dupe the resulting CiIdentifiers before
	// adding to the database (the SQL gets messed up if there's duplicates in what we give to Gorm).

	// As a bit of pre-processing, treat the keys of body.ChartReleaseStatuses like they were also passed in
	// body.ChartReleases. We already dedupe later so this is harmless.
	body.ChartReleases = append(body.ChartReleases, maps.Keys(body.ChartReleaseStatuses)...)

	// First, a new body, starting from the old one.
	bodyAfterSpreading := CiRunV3Upsert{
		Charts:        body.Charts,
		ChartVersions: body.ChartVersions,
		AppVersions:   body.AppVersions,
		Clusters:      body.Clusters,
		Environments:  body.Environments,
		ChartReleases: body.ChartReleases,
		Changesets:    body.Changesets,
	}

	// We know we'll want to handle the ChartReleaseStatuses later by chart release ID,
	// and we're about to happen to query them all, so we'll go ahead and build that mapping.
	chartReleaseIDStatuses := make(map[uint]string, len(body.ChartReleaseStatuses))

	// Environments in the original body should add all their chart releases to the new body, along with the clusters those
	// chart releases belong to.
	for _, environmentSelector := range utils.Dedupe(body.Environments) {
		environmentQuery, err := environmentModelFromSelector(environmentSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var environment models.Environment
		if err = db.Where(&environmentQuery).First(&environment).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var chartReleasesInEnvironment []models.ChartRelease
		if err = db.Where("environment_id = ?", environment.ID).Find(&chartReleasesInEnvironment).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		for _, chartReleaseInEnvironment := range chartReleasesInEnvironment {
			bodyAfterSpreading.ChartReleases = append(bodyAfterSpreading.ChartReleases, utils.UintToString(chartReleaseInEnvironment.ID))
			if chartReleaseInEnvironment.ClusterID != nil {
				bodyAfterSpreading.Clusters = append(bodyAfterSpreading.Clusters, utils.UintToString(*chartReleaseInEnvironment.ClusterID))
			}
		}
	}
	// Same goes for clusters in the original body; we add their chart releases and any environments those chart releases
	// belong to.
	for _, clusterSelector := range utils.Dedupe(body.Clusters) {
		clusterQuery, err := clusterModelFromSelector(clusterSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var cluster models.Cluster
		if err = db.Where(&clusterQuery).First(&cluster).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var chartReleasesInCluster []models.ChartRelease
		if err = db.Where("cluster_id = ?", cluster.ID).Find(&chartReleasesInCluster).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		for _, chartReleaseInCluster := range chartReleasesInCluster {
			bodyAfterSpreading.ChartReleases = append(bodyAfterSpreading.ChartReleases, utils.UintToString(chartReleaseInCluster.ID))
			if chartReleaseInCluster.EnvironmentID != nil {
				bodyAfterSpreading.Environments = append(bodyAfterSpreading.Environments, utils.UintToString(*chartReleaseInCluster.ClusterID))
			}
		}
	}
	// Now for changesets in the original body. They spread to chart releases (and to environments/clusters from there) but they can also
	// spread to new versions they deploy based on the RelateToChangesetNewVersions field.
	for _, changesetSelector := range utils.Dedupe(body.Changesets) {
		changesetID, err := utils.ParseUint(changesetSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var changeset models.Changeset
		if err = db.Scopes(models.ReadChangesetScope).First(&changeset, changesetID).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		if changeset.ChartRelease != nil {
			if changeset.ChartRelease.EnvironmentID != nil {
				bodyAfterSpreading.Environments = append(bodyAfterSpreading.Environments, utils.UintToString(*changeset.ChartRelease.EnvironmentID))
			}
			if changeset.ChartRelease.ClusterID != nil {
				bodyAfterSpreading.Clusters = append(bodyAfterSpreading.Clusters, utils.UintToString(*changeset.ChartRelease.ClusterID))
			}
			// If RelateToChangesetNewVersions is "always", or if it is "when-static" and it's targeting a static environment,
			// add relations for any new app/chart versions deployed.
			if body.RelateToChangesetNewVersions == "always" || (body.RelateToChangesetNewVersions == "when-static" && changeset.ChartRelease.Environment != nil && changeset.ChartRelease.Environment.Lifecycle == "static") {
				for _, newAppVersion := range changeset.NewAppVersions {
					if newAppVersion != nil && newAppVersion.ID != 0 {
						bodyAfterSpreading.AppVersions = append(bodyAfterSpreading.AppVersions, utils.UintToString(newAppVersion.ID))
					}
				}
				for _, newChartVersion := range changeset.NewChartVersions {
					if newChartVersion != nil && newChartVersion.ID != 0 {
						bodyAfterSpreading.ChartVersions = append(bodyAfterSpreading.ChartVersions, utils.UintToString(newChartVersion.ID))
					}
				}
			}
		}
		bodyAfterSpreading.ChartReleases = append(bodyAfterSpreading.ChartReleases, utils.UintToString(changeset.ChartReleaseID))
	}
	// Finally we handle the spreading of chart releases to their environment and cluster.
	for _, chartReleaseSelector := range utils.Dedupe(body.ChartReleases) {
		chartReleaseQuery, err := chartReleaseModelFromSelector(db, chartReleaseSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var chartRelease models.ChartRelease
		if err = db.Where(&chartReleaseQuery).First(&chartRelease).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		if chartRelease.EnvironmentID != nil {
			bodyAfterSpreading.Environments = append(bodyAfterSpreading.Environments, utils.UintToString(*chartRelease.EnvironmentID))
		}
		if chartRelease.ClusterID != nil {
			bodyAfterSpreading.Clusters = append(bodyAfterSpreading.Clusters, utils.UintToString(*chartRelease.ClusterID))
		}
		if status, present := body.ChartReleaseStatuses[chartReleaseSelector]; present {
			chartReleaseIDStatuses[chartRelease.ID] = status
		}
	}

	// With that, we've now handled the spread mechanic. bodyAfterSpreading probably has a ton of duplication, so we go out of our
	// way to de-dupe it. Note the use of utils.Dedupe.
	// We're taking a performance hit because we are potentially re-querying things that we queried above while handling spreading.
	// Frankly, I (Jack) don't care, I prefer an inefficient algorithm that is obviously correct than one where we try to be slick.
	// Requests to this endpoint are going to be from automation anyway; the endpoint could have a latency of several seconds and
	// we wouldn't care.
	var possiblyDuplicatedRelatedResources []models.CiIdentifier
	for _, chartSelector := range utils.Dedupe(bodyAfterSpreading.Charts) {
		chartQuery, err := chartModelFromSelector(chartSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var chart models.Chart
		if err = db.Preload("CiIdentifier").Where(&chartQuery).First(&chart).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, chart.GetCiIdentifier())
	}
	for _, chartVersionSelector := range utils.Dedupe(bodyAfterSpreading.ChartVersions) {
		chartVersionQuery, err := chartVersionModelFromSelector(db, chartVersionSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var chartVersion models.ChartVersion
		if err = db.Preload("CiIdentifier").Where(&chartVersionQuery).First(&chartVersion).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, chartVersion.GetCiIdentifier())
	}
	for _, appVersionSelector := range utils.Dedupe(bodyAfterSpreading.AppVersions) {
		appVersionQuery, err := appVersionModelFromSelector(db, appVersionSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var appVersion models.AppVersion
		if err = db.Preload("CiIdentifier").Where(&appVersionQuery).First(&appVersion).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, appVersion.GetCiIdentifier())
	}
	for _, clusterSelector := range utils.Dedupe(bodyAfterSpreading.Clusters) {
		clusterQuery, err := clusterModelFromSelector(clusterSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var cluster models.Cluster
		if err = db.Preload("CiIdentifier").Where(&clusterQuery).First(&cluster).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, cluster.GetCiIdentifier())
	}
	for _, environmentSelector := range utils.Dedupe(bodyAfterSpreading.Environments) {
		environmentQuery, err := environmentModelFromSelector(environmentSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var environment models.Environment
		if err = db.Preload("CiIdentifier").Where(&environmentQuery).First(&environment).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, environment.GetCiIdentifier())
	}
	for _, chartReleaseSelector := range utils.Dedupe(bodyAfterSpreading.ChartReleases) {
		chartReleaseQuery, err := chartReleaseModelFromSelector(db, chartReleaseSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var chartRelease models.ChartRelease
		if err = db.Preload("CiIdentifier").Where(&chartReleaseQuery).First(&chartRelease).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, chartRelease.GetCiIdentifier())
	}
	for _, changesetSelector := range utils.Dedupe(bodyAfterSpreading.Changesets) {
		changesetID, err := utils.ParseUint(changesetSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		var changeset models.Changeset
		if err = db.Preload("CiIdentifier").First(&changeset, changesetID).Error; err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, changeset.GetCiIdentifier())
	}

	// Now, dedupe one more time (in other words, dedupe by CiIdentifier, not just how the resource was referenced in the body)
	var deduplicatedRelatedResources []models.CiIdentifier
addingToDeduplicatedRelatedResources:
	for _, potentialCiIdentifier := range possiblyDuplicatedRelatedResources {
		for _, existingCiIdentifierInList := range deduplicatedRelatedResources {
			// If we somehow hit a case where we're about to add an empty/uninitialized CiIdentifier...
			// don't add it and log that this happened. I (Jack) think it's impossible to hit this
			// case, but I'd rather skip it and log than knowingly set Gorm up to write bad SQL.
			if potentialCiIdentifier.ID == 0 &&
				(potentialCiIdentifier.ResourceType == "" || potentialCiIdentifier.ResourceID == 0) {
				log.Warn().Msg("encountered an empty CiIdentifier that was considered for upsert via CiRun, skipping")
				continue addingToDeduplicatedRelatedResources
			}

			// If the ID is filled, skip duplicates based on it
			if existingCiIdentifierInList.ID != 0 &&
				potentialCiIdentifier.ID != 0 &&
				existingCiIdentifierInList.ID == potentialCiIdentifier.ID {
				continue addingToDeduplicatedRelatedResources
			}

			// If the resource ID/type is filled, skip duplicates based on it
			if existingCiIdentifierInList.ResourceType != "" &&
				potentialCiIdentifier.ResourceType != "" &&
				existingCiIdentifierInList.ResourceID != 0 &&
				potentialCiIdentifier.ResourceID != 0 &&
				existingCiIdentifierInList.ResourceType == potentialCiIdentifier.ResourceType &&
				existingCiIdentifierInList.ResourceID == potentialCiIdentifier.ResourceID {
				continue addingToDeduplicatedRelatedResources
			}
		}
		// If we didn't continue back to the loop, then add the potential candidate to the final list
		deduplicatedRelatedResources = append(deduplicatedRelatedResources, potentialCiIdentifier)
	}

	// Now to actually mutate the database
	var result models.CiRun

	// Upsert the basic data
	if err = db.Where(&models.CiRun{
		Platform:                   body.Platform,
		GithubActionsOwner:         body.GithubActionsOwner,
		GithubActionsRepo:          body.GithubActionsRepo,
		GithubActionsRunID:         body.GithubActionsRunID,
		GithubActionsAttemptNumber: body.GithubActionsAttemptNumber,
		GithubActionsWorkflowPath:  body.GithubActionsWorkflowPath,
		ArgoWorkflowsNamespace:     body.ArgoWorkflowsNamespace,
		ArgoWorkflowsName:          body.ArgoWorkflowsName,
		ArgoWorkflowsTemplate:      body.ArgoWorkflowsTemplate,
	}).Assign(&models.CiRun{
		StartedAt:             body.StartedAt,
		TerminalAt:            body.TerminalAt,
		Status:                body.Status,
		NotifySlackCustomIcon: body.NotifySlackCustomIcon,
	}).FirstOrCreate(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	// Append the related resources
	if len(deduplicatedRelatedResources) > 0 {
		if err = db.Model(&result).Association("RelatedResources").Append(deduplicatedRelatedResources); err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
	}

	// If we have any resource-specific statuses, add those
	for chartReleaseID, status := range chartReleaseIDStatuses {
		// We want to update the join table ci_runs_for_identifiers's resource_status for ci_identifiers where:
		// 1. The identifier is for a "chart-release" matching our chart release ID
		if err = db.Exec(
			//language=SQL
			`
UPDATE ci_runs_for_identifiers
SET resource_status = ?
FROM ci_identifiers
WHERE
    ci_identifiers.id = ci_runs_for_identifiers.ci_identifier_id
    AND ci_runs_for_identifiers.ci_run_id = ?
    AND ci_identifiers.resource_type = 'chart-release' 
    AND ci_identifiers.resource_id = ?
`, status, result.ID, chartReleaseID).Error; err != nil {
			slack.ReportError(ctx, fmt.Sprintf("error recording chart release status in ci_runs_for_identifiers table for CiRun %d and ChartRelease %d", result.ID, chartReleaseID), err)
		}
		// 2. The identifier is for a "changeset" where the changeset's chart release ID matches our chart release ID
		//		(We have to do this one separately because this one needs a join, and the operation above shouldn't
		//      be limited to identifiers that can join with changesets)
		var changesetIDs []uint
		if err = db.Raw(
			//language=SQL
			`
UPDATE ci_runs_for_identifiers
SET resource_status = ?
FROM ci_identifiers
    JOIN changesets ON changesets.id = ci_identifiers.resource_id
WHERE
    ci_identifiers.id = ci_runs_for_identifiers.ci_identifier_id
    AND ci_runs_for_identifiers.ci_run_id = ?
    AND ci_identifiers.resource_type = 'changeset' 
    AND changesets.chart_release_id = ?
RETURNING changesets.id
`, status, result.ID, chartReleaseID).Scan(&changesetIDs).Error; err != nil {
			slack.ReportError(ctx, fmt.Sprintf("error recording changeset status in ci_runs_for_identifiers table for CiRun %d and ChartRelease %d", result.ID, chartReleaseID), err)
		}
		for _, changesetID := range changesetIDs {
			// If there were changesets from step 2:
			// 3. The identifier is for a new app version on that changeset
			if err = db.Exec(
				//language=SQL
				`
UPDATE ci_runs_for_identifiers
SET resource_status = ?
FROM ci_identifiers
    JOIN changeset_new_app_versions ON ci_identifiers.resource_id = changeset_new_app_versions.app_version_id
WHERE
    ci_runs_for_identifiers.ci_run_id = ?
    AND ci_runs_for_identifiers.ci_identifier_id = ci_identifiers.id
    AND ci_identifiers.resource_type = 'app-version'
    AND changeset_new_app_versions.changeset_id = ?
`, status, result.ID, changesetID).Error; err != nil {
				slack.ReportError(ctx, fmt.Sprintf("error recording app version status in ci_runs_for_identifiers table for CiRun %d and ChartRelease %d via Changeset %d", result.ID, chartReleaseID, changesetID), err)
			}
			// 4. The identifier is for a new chart version on that changeset
			if err = db.Exec(
				//language=SQL
				`
UPDATE ci_runs_for_identifiers
SET resource_status = ?
FROM ci_identifiers
    JOIN changeset_new_chart_versions ON ci_identifiers.resource_id = changeset_new_chart_versions.chart_version_id
WHERE
    ci_runs_for_identifiers.ci_run_id = ?
    AND ci_runs_for_identifiers.ci_identifier_id = ci_identifiers.id
    AND ci_identifiers.resource_type = 'chart-version'
    AND changeset_new_chart_versions.changeset_id = ?
`, status, result.ID, changesetID).Error; err != nil {
				slack.ReportError(ctx, fmt.Sprintf("error recording chart version status in ci_runs_for_identifiers table for CiRun %d and ChartRelease %d via Changeset %d", result.ID, chartReleaseID, changesetID), err)
			}
		}
	}

	// If the request added any Slack channels for us to notify, record those
	if len(body.NotifySlackChannelsUponSuccess) > 0 || len(body.NotifySlackChannelsUponFailure) > 0 || len(body.NotifySlackChannelsUponRetry) > 0 {
		var channelUpdates models.CiRun
		if len(body.NotifySlackChannelsUponSuccess) > 0 {
			channelUpdates.NotifySlackChannelsUponSuccess = utils.Dedupe(append(result.NotifySlackChannelsUponSuccess, body.NotifySlackChannelsUponSuccess...))
		}
		if len(body.NotifySlackChannelsUponFailure) > 0 {
			channelUpdates.NotifySlackChannelsUponFailure = utils.Dedupe(append(result.NotifySlackChannelsUponFailure, body.NotifySlackChannelsUponFailure...))
		}
		if len(body.NotifySlackChannelsUponRetry) > 0 {
			channelUpdates.NotifySlackChannelsUponRetry = utils.Dedupe(append(result.NotifySlackChannelsUponRetry, body.NotifySlackChannelsUponRetry...))
		}
		if err = db.Model(&result).Omit(clause.Associations).Updates(&channelUpdates).Error; err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
	}

	// If this workflow finished, claim that we're doing the final dispatch of any hooks
	var claim string
	if result.TerminalAt != nil {
		claim = result.AttemptToClaimTerminationDispatch(db)
	}

	// Re-query so we load all the CiIdentifiers, including any added by previous requests
	// This also gets TerminationHooksDispatchedAt back out of the database
	if err = db.Preload(clause.Associations).First(&result, result.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if err = result.FillRelatedResourceStatuses(db); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	// Dispatch either when the workflow hasn't finished or when it has but we claimed
	// the final dispatch
	if result.TerminalAt == nil || result.EvaluateIfTerminationClaimHeld(claim) {
		ci_hooks.Dispatch(db, result)
	}

	ctx.JSON(http.StatusCreated, ciRunFromModel(result))
}
