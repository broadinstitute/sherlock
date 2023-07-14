package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type CiRunV3Upsert struct {
	ciRunFields
	Charts                     []string `json:"charts" form:"-"`                     // Always appends; will eliminate duplicates.
	ChartVersions              []string `json:"chartVersions" form:"-"`              // Always appends; will eliminate duplicates.
	AppVersions                []string `json:"appVersions" form:"-"`                // Always appends; will eliminate duplicates.
	Clusters                   []string `json:"clusters" form:"-"`                   // Always appends; will eliminate duplicates.
	Environments               []string `json:"environments" form:"-"`               // Always appends; will eliminate duplicates.
	ChartReleases              []string `json:"chartReleases" form:"-"`              // Always appends; will eliminate duplicates. Spreads to associated environments and clusters.
	Changesets                 []string `json:"changesets" form:"-"`                 // Always appends; will eliminate duplicates. Spreads to associated chart releases (and environments and clusters).
	ChangesetsSpreadToVersions []string `json:"changesetsSpreadToVersions" form:"-"` // Always appends; will eliminate duplicates. Spreads to associated chart releases (and environments and clusters) but also new app/chart versions.
}

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

	// Set related resources
	var relatedResources []models.CiIdentifier
	for _, changesetSelector := range body.ChangesetsSpreadToVersions {
		changeset, err := v2models.InternalChangesetStore.GetBySelector(db, changesetSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		if changeset.ChartReleaseID != 0 {
			body.ChartReleases = append(body.ChartReleases, utils.UintToString(changeset.ChartReleaseID))
		}
		for _, newAppVersion := range changeset.NewAppVersions {
			if newAppVersion != nil && newAppVersion.ID != 0 {
				body.AppVersions = append(body.AppVersions, utils.UintToString(newAppVersion.ID))
			}
		}
		for _, newChartVersion := range changeset.NewChartVersions {
			if newChartVersion != nil && newChartVersion.ID != 0 {
				body.ChartVersions = append(body.ChartVersions, utils.UintToString(newChartVersion.ID))
			}
		}
		relatedResources = append(relatedResources, ciIdentifierModelFromOldModel(changeset))
	}
	for _, changesetSelector := range body.Changesets {
		changeset, err := v2models.InternalChangesetStore.GetBySelector(db, changesetSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		if changeset.ChartReleaseID != 0 {
			body.ChartReleases = append(body.ChartReleases, utils.UintToString(changeset.ChartReleaseID))
		}
		relatedResources = append(relatedResources, ciIdentifierModelFromOldModel(changeset))
	}
	for _, chartReleaseSelector := range body.ChartReleases {
		chartRelease, err := v2models.InternalChartReleaseStore.GetBySelector(db, chartReleaseSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		if chartRelease.EnvironmentID != nil {
			body.Environments = append(body.Environments, utils.UintToString(*chartRelease.EnvironmentID))
		}
		if chartRelease.ClusterID != nil {
			body.Clusters = append(body.Clusters, utils.UintToString(*chartRelease.ClusterID))
		}
		relatedResources = append(relatedResources, ciIdentifierModelFromOldModel(chartRelease))
	}
	for _, chartSelector := range body.Charts {
		chart, err := v2models.InternalChartStore.GetBySelector(db, chartSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		relatedResources = append(relatedResources, ciIdentifierModelFromOldModel(chart))
	}
	for _, chartVersionSelector := range body.ChartVersions {
		chartVersion, err := v2models.InternalChartVersionStore.GetBySelector(db, chartVersionSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		relatedResources = append(relatedResources, ciIdentifierModelFromOldModel(chartVersion))
	}
	for _, appVersionSelector := range body.AppVersions {
		appVersion, err := v2models.InternalAppVersionStore.GetBySelector(db, appVersionSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		relatedResources = append(relatedResources, ciIdentifierModelFromOldModel(appVersion))
	}
	for _, clusterSelector := range body.Clusters {
		cluster, err := v2models.InternalClusterStore.GetBySelector(db, clusterSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		relatedResources = append(relatedResources, ciIdentifierModelFromOldModel(cluster))
	}
	for _, environmentSelector := range body.Environments {
		environment, err := v2models.InternalEnvironmentStore.GetBySelector(db, environmentSelector)
		if err != nil {
			ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
			return
		}
		relatedResources = append(relatedResources, ciIdentifierModelFromOldModel(environment))
	}
	if err = db.Model(&result).Association("RelatedResources").Append(relatedResources); err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return
	}

	// Re-query
	if err = db.First(&result, result.ID).Error; err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, ciRunFromModel(result))
}
