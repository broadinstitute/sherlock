package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// chartVersionsV3Upsert godoc
//
//	@summary		Upsert a ChartVersion
//	@description	Upsert a ChartVersion.
//	@tags			ChartVersions
//	@accept			json
//	@produce		json
//	@param			chartVersion			body		ChartVersionV3Create	true	"The ChartVersion to upsert"
//	@success		201						{object}	ChartVersionV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/chart-versions/v3 [put]
func chartVersionsV3Upsert(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body ChartVersionV3Create
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	toUpsert, err := body.toModel(db, false)
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error parsing body: %w", err))
		return
	}

	if err = db.Where(&models.ChartVersion{
		ChartVersion: toUpsert.ChartVersion,
		ChartID:      toUpsert.ChartID,
	}).Assign(&models.ChartVersion{
		Description: toUpsert.Description,
	}).Attrs(&models.ChartVersion{
		ParentChartVersionID: toUpsert.ParentChartVersionID,
	}).FirstOrCreate(&toUpsert).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var result models.ChartVersion
	if err = db.Preload(clause.Associations).First(&result, toUpsert.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, chartVersionFromModel(result))
}
