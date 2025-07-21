package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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
