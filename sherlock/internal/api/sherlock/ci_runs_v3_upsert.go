package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm/clause"
	"net/http"
)

type CiRunV3Upsert struct {
	ciRunFields
	Charts                       []string `json:"charts"`                       // Always appends; will eliminate duplicates.
	ChartVersions                []string `json:"chartVersions"`                // Always appends; will eliminate duplicates.
	AppVersions                  []string `json:"appVersions"`                  // Always appends; will eliminate duplicates.
	Clusters                     []string `json:"clusters"`                     // Always appends; will eliminate duplicates. Spreads to contained chart releases and their environments.
	Environments                 []string `json:"environments"`                 // Always appends; will eliminate duplicates. Spreads to contained chart releases and their clusters.
	ChartReleases                []string `json:"chartReleases"`                // Always appends; will eliminate duplicates. Spreads to associated environments and clusters.
	Changesets                   []string `json:"changesets"`                   // Always appends; will eliminate duplicates. Spreads to associated chart releases, environments, and clusters.
	RelateToChangesetNewVersions bool     `json:"relateToChangesetNewVersions"` // Makes entries in the changesets field also spread to new app versions and chart versions deployed by the changeset.
}

// ciRunsV3Upsert godoc
//
//	@summary		Create or update a CiRun
//	@description	Create or update a CiRun with timing, status, and related resource information. This endpoint is idempotent.
//	@description	It's recommended to take note of the description of the individual fields when reporting resource relations.
//	@description	The "spreading" behavior means that clients don't need to be smart about how resources relate -- Sherlock
//	@description	will handle it as long as the client reports what the run directly relates to.
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
	if err = ctx.MustBindWith(&body, binding.JSON); err != nil {
		return
	}
	var result models.CiRun

	// Upsert
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
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return
	}

	// Append related resources

	// We want to handle the "spreading" mechanic that some of the fields have. To do that, we'll literally just re-assemble
	// the body we got into a body where spreading would be a no-op. Then we'll handle that body and de-dupe the resulting
	// CiIdentifiers before adding to the database.

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
	// Environments in the original body should add all their chart releases to the new body, along with the clusters those
	// chart releases belong to.
	for _, environmentSelector := range body.Environments {
		environmentID, err := v2models.InternalEnvironmentStore.ResolveSelector(db, environmentSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		chartReleasesInEnvironment, err := v2models.InternalChartReleaseStore.ListAllMatchingByCreated(db, 0, v2models.ChartRelease{
			EnvironmentID: &environmentID,
		})
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
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
	for _, clusterSelector := range body.Clusters {
		clusterID, err := v2models.InternalClusterStore.ResolveSelector(db, clusterSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		chartReleasesInCluster, err := v2models.InternalChartReleaseStore.ListAllMatchingByCreated(db, 0, v2models.ChartRelease{
			ClusterID: &clusterID,
		})
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		for _, chartReleaseInCluster := range chartReleasesInCluster {
			bodyAfterSpreading.ChartReleases = append(bodyAfterSpreading.ChartReleases, utils.UintToString(chartReleaseInCluster.ID))
			if chartReleaseInCluster.EnvironmentID != nil {
				bodyAfterSpreading.Environments = append(bodyAfterSpreading.Environments, utils.UintToString(*chartReleaseInCluster.ClusterID))
			}
		}
	}
	// Now for changesets in the original body. They can spread to new app/chart versions, but mainly they spread to chart releases and then
	// on to the chart releases' environment/cluster. Rather than duplicating code with handling normal chart release spreading below, we'll
	// just make a list here and use it there in a sec.
	var chartReleasesFromChangesets []string
	for _, changesetSelector := range body.Changesets {
		changeset, err := v2models.InternalChangesetStore.GetBySelector(db, changesetSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		if body.RelateToChangesetNewVersions {
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
		chartReleasesFromChangesets = append(chartReleasesFromChangesets, utils.UintToString(changeset.ChartReleaseID))
	}
	// Finally we handle the spreading of chart releases to their environment and cluster. We care about chart releases in the original
	// body and also ones we just pulled from changesets above, so we concatenate those lists for the loop here.
	for _, chartReleaseSelector := range append(body.ChartReleases, chartReleasesFromChangesets...) {
		chartRelease, err := v2models.InternalChartReleaseStore.GetBySelector(db, chartReleaseSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		if chartRelease.EnvironmentID != nil {
			bodyAfterSpreading.Environments = append(bodyAfterSpreading.Environments, utils.UintToString(*chartRelease.EnvironmentID))
		}
		if chartRelease.ClusterID != nil {
			bodyAfterSpreading.Clusters = append(bodyAfterSpreading.Clusters, utils.UintToString(*chartRelease.ClusterID))
		}
	}

	// With that, we've now handled the spread mechanic. bodyAfterSpreading probably has a ton of duplication, so we go out of our
	// way to de-dupe it. Note the use of utils.Dedupe.
	// We're taking a performance hit because we are potentially re-querying things that we queried above while handling spreading.
	// Frankly, I (Jack) don't care, I prefer an inefficient algorithm that is obviously correct than one where we try to be slick.
	var possiblyDuplicatedRelatedResources []models.CiIdentifier
	for _, chartSelector := range utils.Dedupe(bodyAfterSpreading.Charts) {
		chart, err := v2models.InternalChartStore.GetBySelector(db, chartSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(chart))
	}
	for _, chartVersionSelector := range utils.Dedupe(bodyAfterSpreading.ChartVersions) {
		chartVersion, err := v2models.InternalChartVersionStore.GetBySelector(db, chartVersionSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(chartVersion))
	}
	for _, appVersionSelector := range utils.Dedupe(bodyAfterSpreading.AppVersions) {
		appVersion, err := v2models.InternalAppVersionStore.GetBySelector(db, appVersionSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(appVersion))
	}
	for _, clusterSelector := range utils.Dedupe(bodyAfterSpreading.Clusters) {
		cluster, err := v2models.InternalClusterStore.GetBySelector(db, clusterSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(cluster))
	}
	for _, environmentSelector := range utils.Dedupe(bodyAfterSpreading.Environments) {
		environment, err := v2models.InternalEnvironmentStore.GetBySelector(db, environmentSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(environment))
	}
	for _, chartReleaseSelector := range utils.Dedupe(bodyAfterSpreading.ChartReleases) {
		chartRelease, err := v2models.InternalChartReleaseStore.GetBySelector(db, chartReleaseSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(chartRelease))
	}
	for _, changesetSelector := range utils.Dedupe(bodyAfterSpreading.Changesets) {
		changeset, err := v2models.InternalChangesetStore.GetBySelector(db, changesetSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		possiblyDuplicatedRelatedResources = append(possiblyDuplicatedRelatedResources, ciIdentifierModelFromOldModel(changeset))
	}

	// Now, dedupe one more time (in other words, dedupe by CiIdentifier, not just how the resource was referenced in the body)
	var deduplicatedRelatedResources []models.CiIdentifier
addingToDeduplicatedRelatedResources:
	for _, potentialCiIdentifier := range possiblyDuplicatedRelatedResources {
		for _, existingCiIdentifierInList := range deduplicatedRelatedResources {
			if existingCiIdentifierInList.ID == potentialCiIdentifier.ID &&
				existingCiIdentifierInList.ResourceType == potentialCiIdentifier.ResourceType &&
				existingCiIdentifierInList.ResourceID == potentialCiIdentifier.ResourceID {
				continue addingToDeduplicatedRelatedResources
			}
		}
		deduplicatedRelatedResources = append(deduplicatedRelatedResources, potentialCiIdentifier)
	}

	// Finally, append the association.
	if err = db.Model(&result).Association("RelatedResources").Append(deduplicatedRelatedResources); err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return
	}

	// Re-query
	if err = db.Preload(clause.Associations).First(&result, result.ID).Error; err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, ciRunFromModel(result))
}
