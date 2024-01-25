package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// appVersionsV3Edit godoc
//
//	@summary		Edit an individual AppVersion
//	@description	Edit an individual AppVersion.
//	@tags			AppVersions
//	@produce		json
//	@param			selector				path		string				true	"The selector of the AppVersion, which can be either a numeric ID or chart/version."
//	@param			appVersion				body		AppVersionV3Edit	true	"The edits to make to the AppVersion"
//	@success		200						{object}	AppVersionV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/app-versions/v3/{selector} [patch]
func appVersionsV3Edit(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := appVersionModelFromSelector(db, canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var body AppVersionV3Edit
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	edits, err := body.toModel(db, true)
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}

	var toEdit models.AppVersion
	if err = db.Preload(clause.Associations).Where(&query).First(&toEdit).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Model(&toEdit).Omit(clause.Associations).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, appVersionFromModel(toEdit))
}
