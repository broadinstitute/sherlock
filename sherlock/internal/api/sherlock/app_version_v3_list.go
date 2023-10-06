package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// appVersionsV3List godoc
//
//	@summary		List AppVersions matching a filter
//	@description	List AppVersions matching a filter.
//	@tags			AppVersions
//	@produce		json
//	@param			filter					query		AppVersionV3	false	"Filter the returned AppVersions"
//	@param			limit					query		int				false	"Control how many AppVersions are returned (default 100)"
//	@param			offset					query		int				false	"Control the offset for the returned AppVersions (default 0)"
//	@success		200						{array}		AppVersionV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/app-versions/v3 [get]
func appVersionsV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter AppVersionV3
	if err = ctx.ShouldBindQuery(&filter); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	modelFilter, err := filter.toModel(db, true)
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}

	limit, err := utils.ParseInt(ctx.DefaultQuery("limit", "100"))
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}
	offset, err := utils.ParseInt(ctx.DefaultQuery("offset", "0"))
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}
	var results []models.AppVersion
	if err = db.
		Where(&modelFilter).
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, appVersionFromModel))
}
