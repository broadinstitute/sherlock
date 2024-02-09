package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// chartVersionsV3List godoc
//
//	@summary		List ChartVersions matching a filter
//	@description	List ChartVersions matching a filter.
//	@tags			ChartVersions
//	@produce		json
//	@param			filter					query		ChartVersionV3	false	"Filter the returned ChartVersions"
//	@param			limit					query		int				false	"Control how many ChartVersions are returned (default 100)"
//	@param			offset					query		int				false	"Control the offset for the returned ChartVersions (default 0)"
//	@success		200						{array}		ChartVersionV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/chart-versions/v3 [get]
func chartVersionsV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter ChartVersionV3
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
	var results []models.ChartVersion
	if err = db.
		Where(&modelFilter).
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Preload(clause.Associations).
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, chartVersionFromModel))
}
