package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// appVersionsV3Upsert godoc
//
//	@summary		Upsert a AppVersion
//	@description	Upsert a AppVersion.
//	@tags			AppVersions
//	@accept			json
//	@produce		json
//	@param			appVersion				body		AppVersionV3Create	true	"The AppVersion to upsert"
//	@success		201						{object}	AppVersionV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/appVersions/v3 [post]
func appVersionsV3Upsert(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body AppVersionV3Create
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

	if err = db.Where(&models.AppVersion{
		AppVersion: toUpsert.AppVersion,
		ChartID:    toUpsert.ChartID,
	}).Assign(&models.AppVersion{
		Description: toUpsert.Description,
	}).Attrs(&models.AppVersion{
		GitCommit:          toUpsert.GitCommit,
		GitBranch:          toUpsert.GitBranch,
		ParentAppVersionID: toUpsert.ParentAppVersionID,
	}).FirstOrCreate(&toUpsert).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var result models.AppVersion
	if err = db.Preload(clause.Associations).First(&result, toUpsert.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, appVersionFromModel(result))
}
