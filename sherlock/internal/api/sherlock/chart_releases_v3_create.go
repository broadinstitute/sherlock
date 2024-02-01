package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

// chartReleasesV3Create godoc
//
//	@summary		Create a ChartRelease
//	@description	Create a ChartRelease.
//	@tags			ChartReleases
//	@accept			json
//	@produce		json
//	@param			chartRelease			body		ChartReleaseV3Create	true	"The ChartRelease to create"
//	@success		201						{object}	ChartReleaseV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/chart-releases/v3 [post]
func chartReleasesV3Create(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body ChartReleaseV3Create
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	toCreate, err := body.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if toCreate.ChartID == 0 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) chart is required", errors.BadRequest))
		return
	}
	if (toCreate.EnvironmentID == nil || *toCreate.EnvironmentID == 0) && (toCreate.ClusterID == nil || *toCreate.ClusterID == 0) {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) either environment or cluster is required", errors.BadRequest))
		return
	}

	if err = fillChartReleaseDynamicDefaultsForCreation(db, &toCreate); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Create(&toCreate).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var result models.ChartRelease
	if err = db.Preload(clause.Associations).First(&result, toCreate.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, chartReleaseFromModel(result))
}

// fillChartReleaseDynamicDefaultsForCreation is a helper specifically for chartReleasesV3Create.
// It is inefficient -- it potentially loads things from the database it doesn't need -- but that's intentional.
// chartReleasesV3Create will rarely be called (and almost always from the UI) so it's better for this function
// to be easy to understand than to worry about an errant load of a row or two by primary key.
func fillChartReleaseDynamicDefaultsForCreation(db *gorm.DB, toCreate *models.ChartRelease) error {
	var chart models.Chart
	if err := db.Model(&models.Chart{Model: gorm.Model{ID: toCreate.ChartID}}).Take(&chart).Error; err != nil {
		return fmt.Errorf("failed to read chart to evaluate defaults: %w", err)
	}

	// If the chart has a branch and toCreate doesn't, use the chart's branch.
	// We may end up ignoring this if the branch resolver doesn't end up being chosen.
	if chart.AppImageGitMainBranch != nil && *chart.AppImageGitMainBranch != "" && toCreate.AppVersionBranch == nil {
		toCreate.AppVersionBranch = chart.AppImageGitMainBranch
	}

	// If the chart is marked as exposing an endpoint, copy fields from it.
	if chart.ChartExposesEndpoint != nil && *chart.ChartExposesEndpoint {
		if toCreate.Subdomain == nil {
			toCreate.Subdomain = chart.DefaultSubdomain
		}
		if toCreate.Protocol == nil {
			toCreate.Protocol = chart.DefaultProtocol
		}
		if toCreate.Port == nil {
			toCreate.Port = chart.DefaultPort
		}
	}

	// If toCreate doesn't have a set app version resolver, set it based on what's available.
	// Branch takes last priority because we always try to fill that from the chart.
	if toCreate.AppVersionResolver == nil {
		resolver := "none"
		if toCreate.AppVersionExact != nil {
			resolver = "exact"
		} else if toCreate.AppVersionCommit != nil {
			resolver = "commit"
		} else if toCreate.AppVersionFollowChartReleaseID != nil {
			resolver = "follow"
		} else if toCreate.AppVersionBranch != nil {
			resolver = "branch"
		}
		toCreate.AppVersionResolver = &resolver
	}

	// If toCreate doesn't have a set chart version resolver, set it based on what's available.
	if toCreate.ChartVersionResolver == nil {
		resolver := "latest"
		if toCreate.ChartVersionExact != nil {
			resolver = "exact"
		} else if toCreate.ChartVersionFollowChartReleaseID != nil {
			resolver = "follow"
		}
		toCreate.ChartVersionResolver = &resolver
	}

	// If we have an environment, use it to fill in defaults.
	if toCreate.EnvironmentID != nil {
		var environment models.Environment
		if err := db.Model(&models.Environment{Model: gorm.Model{ID: *toCreate.EnvironmentID}}).Take(&environment).Error; err != nil {
			return fmt.Errorf("failed to read environment to evaluate defaults: %w", err)
		}

		// Name like "leonardo-prod"
		if toCreate.Name == "" {
			toCreate.Name = fmt.Sprintf("%s-%s", chart.Name, environment.Name)
		}

		// If there's no cluster, add it
		if toCreate.ClusterID == nil && environment.DefaultClusterID != nil {
			toCreate.ClusterID = environment.DefaultClusterID
		}

		// If there's no namespace, add it
		if toCreate.Namespace == "" && environment.DefaultNamespace != "" {
			toCreate.Namespace = environment.DefaultNamespace
		}

		// If there's no firecloud develop ref, add it
		// (We'll remove this shortly because fc-dev is no more, but keeping behavioral parity makes sense for the moment)
		if toCreate.FirecloudDevelopRef == nil && environment.DefaultFirecloudDevelopRef != nil &&
			chart.LegacyConfigsEnabled != nil && *chart.LegacyConfigsEnabled {
			toCreate.FirecloudDevelopRef = environment.DefaultFirecloudDevelopRef
		}
	}

	// If we have a cluster, use it to fill in defaults.
	// The only one we care about is name, so we dodge the whole block if the name is already filled.
	// This also means that when the cluster got filled from the environment, we won't run this, because the name
	// would've been filled from the environment too.
	if toCreate.ClusterID != nil && toCreate.Name == "" {
		var cluster models.Cluster
		if err := db.Model(&models.Cluster{Model: gorm.Model{ID: *toCreate.ClusterID}}).Take(&cluster).Error; err != nil {
			return fmt.Errorf("failed to read cluster to evaluate defaults: %w", err)
		}
		if toCreate.Namespace == "" || toCreate.Namespace == cluster.Name {
			toCreate.Name = fmt.Sprintf("%s-%s", chart.Name, cluster.Name)
		} else {
			toCreate.Name = fmt.Sprintf("%s-%s-%s", chart.Name, toCreate.Namespace, cluster.Name)
		}
	}

	return nil
}
