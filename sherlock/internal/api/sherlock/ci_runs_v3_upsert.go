package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/gha_oidc/gha_oidc_claims"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/deployhooks"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strings"
	"time"
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
		environmentID, err := v2models.InternalEnvironmentStore.ResolveSelector(db, environmentSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		chartReleasesInEnvironment, err := v2models.InternalChartReleaseStore.ListAllMatchingByCreated(db, 0, v2models.ChartRelease{
			EnvironmentID: &environmentID,
		})
		if err != nil {
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
		clusterID, err := v2models.InternalClusterStore.ResolveSelector(db, clusterSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		chartReleasesInCluster, err := v2models.InternalChartReleaseStore.ListAllMatchingByCreated(db, 0, v2models.ChartRelease{
			ClusterID: &clusterID,
		})
		if err != nil {
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
		changeset, err := v2models.InternalChangesetStore.GetBySelector(db, changesetSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		chartRelease, err := v2models.InternalChartReleaseStore.Get(db, v2models.ChartRelease{Model: gorm.Model{ID: changeset.ChartReleaseID}})
		if err != nil {
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
		// If RelateToChangesetNewVersions is "always", or if it is "when-static" and it's targeting a static environment,
		// add relations for any new app/chart versions deployed.
		if body.RelateToChangesetNewVersions == "always" || (body.RelateToChangesetNewVersions == "when-static" && chartRelease.Environment != nil && chartRelease.Environment.Lifecycle == "static") {
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
		bodyAfterSpreading.ChartReleases = append(bodyAfterSpreading.ChartReleases, utils.UintToString(changeset.ChartReleaseID))
	}
	// Finally we handle the spreading of chart releases to their environment and cluster.
	for _, chartReleaseSelector := range utils.Dedupe(body.ChartReleases) {
		chartRelease, err := v2models.InternalChartReleaseStore.GetBySelector(db, chartReleaseSelector)
		if err != nil {
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
		chart, err := v2models.InternalChartStore.GetBySelector(db, chartSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(chart))
	}
	for _, chartVersionSelector := range utils.Dedupe(bodyAfterSpreading.ChartVersions) {
		chartVersion, err := v2models.InternalChartVersionStore.GetBySelector(db, chartVersionSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(chartVersion))
	}
	for _, appVersionSelector := range utils.Dedupe(bodyAfterSpreading.AppVersions) {
		appVersion, err := v2models.InternalAppVersionStore.GetBySelector(db, appVersionSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(appVersion))
	}
	for _, clusterSelector := range utils.Dedupe(bodyAfterSpreading.Clusters) {
		cluster, err := v2models.InternalClusterStore.GetBySelector(db, clusterSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(cluster))
	}
	for _, environmentSelector := range utils.Dedupe(bodyAfterSpreading.Environments) {
		environment, err := v2models.InternalEnvironmentStore.GetBySelector(db, environmentSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(environment))
	}
	for _, chartReleaseSelector := range utils.Dedupe(bodyAfterSpreading.ChartReleases) {
		chartRelease, err := v2models.InternalChartReleaseStore.GetBySelector(db, chartReleaseSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(chartRelease))
	}
	for _, changesetSelector := range utils.Dedupe(bodyAfterSpreading.Changesets) {
		changeset, err := v2models.InternalChangesetStore.GetBySelector(db, changesetSelector)
		if err != nil {
			if body.IgnoreBadSelectors {
				continue
			} else {
				errors.AbortRequest(ctx, err)
				return
			}
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(changeset))
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
		StartedAt:  body.StartedAt,
		TerminalAt: body.TerminalAt,
		Status:     body.Status,
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
			slack.ReportError(ctx, fmt.Errorf("error recording chart release status in ci_runs_for_identifiers table for CiRun %d and ChartRelease %d: %w", result.ID, chartReleaseID, err))
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
			slack.ReportError(ctx, fmt.Errorf("error recording changeset status in ci_runs_for_identifiers table for CiRun %d and ChartRelease %d: %w", result.ID, chartReleaseID, err))
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
				slack.ReportError(ctx, fmt.Errorf("error recording app version status in ci_runs_for_identifiers table for CiRun %d and ChartRelease %d via Changeset %d: %w", result.ID, chartReleaseID, changesetID, err))
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
				slack.ReportError(ctx, fmt.Errorf("error recording chart version status in ci_runs_for_identifiers table for CiRun %d and ChartRelease %d via Changeset %d: %w", result.ID, chartReleaseID, changesetID, err))
			}
		}
	}

	// If the request added any Slack channels for us to notify, record those
	if len(body.NotifySlackChannelsUponSuccess) > 0 || len(body.NotifySlackChannelsUponFailure) > 0 {
		var channelUpdates models.CiRun
		if len(body.NotifySlackChannelsUponSuccess) > 0 {
			channelUpdates.NotifySlackChannelsUponSuccess = utils.Dedupe(append(result.NotifySlackChannelsUponSuccess, body.NotifySlackChannelsUponSuccess...))
		}
		if len(body.NotifySlackChannelsUponFailure) > 0 {
			channelUpdates.NotifySlackChannelsUponFailure = utils.Dedupe(append(result.NotifySlackChannelsUponFailure, body.NotifySlackChannelsUponFailure...))
		}
		if err = db.Model(&result).Updates(&channelUpdates).Error; err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
	}

	// If this workflow finished, claim that we're dispatching any hooks
	var dispatchedAt string
	if result.TerminalAt != nil && result.TerminationHooksDispatchedAt == nil {
		dispatchedAt = time.Now().Format(time.RFC3339Nano)
		if err = db.Model(&result).Update("termination_hooks_dispatched_at", gorm.Expr("COALESCE(termination_hooks_dispatched_at, ?)", dispatchedAt)).Error; err != nil {
			log.Warn().Err(err).Msgf("HOOK | failed to claim dispatch on CiRun %d", result.ID)
			dispatchedAt = ""
		}
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

	// If we said we were going to dispatch this workflow, check that our claim held and then
	// do the dispatch
	if dispatchedAt != "" {
		if result.TerminationHooksDispatchedAt == nil {
			log.Warn().Msgf("HOOK | claimed dispatch on CiRun %d but the field wasn't set when re-queried?", result.ID)
		} else if dispatchedAt != *result.TerminationHooksDispatchedAt {
			log.Info().Msgf("HOOK | parallelism detected; we claimed dispatch on CiRun %d at %s but it the claim in the database ended up being %s", result.ID, dispatchedAt, *result.TerminationHooksDispatchedAt)
		} else if config.Config.String("mode") == "debug" {
			// Locally we do this synchronously so it occurs during the request.
			// We don't do anything to either the db or the result during request
			// exit, so the difference between local dev and remove shouldn't be
			// a cause of issues. If it is, that's what we've got error reporting
			// for.
			dispatchTerminatedCiRun(db, result)
		} else {
			go dispatchTerminatedCiRun(db, result)
		}
	}

	ctx.JSON(http.StatusCreated, ciRunFromModel(result))
}

// dispatchTerminatedCiRun is abstracted from ciRunsV3Upsert so that it can be easily called
// synchronously in debug mode and asynchronously otherwise.
func dispatchTerminatedCiRun(db *gorm.DB, ciRun models.CiRun) {
	if deployhooks.CiRunIsDeploy(ciRun) {
		deployhooks.DispatchCiRun(db, ciRun)
	}
	if ciRun.Status != nil && (len(ciRun.NotifySlackChannelsUponSuccess) > 0 || len(ciRun.NotifySlackChannelsUponFailure) > 0) {
		text := makeSlackMessageText(db, ciRun)
		if ciRun.Succeeded() {
			for _, channel := range ciRun.NotifySlackChannelsUponSuccess {
				slack.SendMessage(db.Statement.Context, channel, "", slack.GreenBlock{Text: text})
			}
		} else {
			for _, channel := range ciRun.NotifySlackChannelsUponFailure {
				slack.SendMessage(db.Statement.Context, channel, "", slack.RedBlock{Text: text})
			}
		}
	}
}

func makeSlackMessageText(db *gorm.DB, ciRun models.CiRun) string {
	var relatedResourceSummaryParts []string
	var chartReleaseIDs, environmentIDs []uint
	for _, identifier := range ciRun.RelatedResources {
		if identifier.ResourceType == "chart-release" {
			chartReleaseIDs = append(chartReleaseIDs, identifier.ResourceID)
		} else if identifier.ResourceType == "environment" {
			environmentIDs = append(environmentIDs, identifier.ResourceID)
		}
	}
	if len(chartReleaseIDs) > 0 {
		var chartReleases []models.ChartRelease
		if err := db.Model(&models.ChartRelease{}).Find(&chartReleases, chartReleaseIDs).Error; err == nil {
			relatedResourceSummaryParts = append(relatedResourceSummaryParts, utils.Map(chartReleases, func(c models.ChartRelease) string {
				return slack.LinkHelper(fmt.Sprintf(config.Config.String("beehive.chartReleaseUrlFormatString"), c.Name), c.Name)
			})...)
		}
	} else if len(environmentIDs) > 0 {
		var environments []models.Environment
		if err := db.Model(&models.Environment{}).Find(&environments, environmentIDs).Error; err == nil {
			relatedResourceSummaryParts = append(relatedResourceSummaryParts, utils.Map(environments, func(e models.Environment) string {
				return slack.LinkHelper(fmt.Sprintf(config.Config.String("beehive.environmentUrlFormatString"), e.Name), e.Name)
			})...)
		}
	}
	var against string
	if len(relatedResourceSummaryParts) > 0 {
		against = fmt.Sprintf(" against %s", strings.Join(relatedResourceSummaryParts, ", "))
	}
	return fmt.Sprintf("%s%s: %s", ciRun.Nickname(), against, slack.LinkHelper(ciRun.WebURL(), *ciRun.Status))
}
