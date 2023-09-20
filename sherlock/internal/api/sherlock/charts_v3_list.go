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

// chartsV3List godoc
//
//	@summary		List Charts matching a filter
//	@description	List Charts matching a filter.
//	@tags			Charts
//	@produce		json
//	@param			filter					query		ChartV3	false	"Filter the returned Charts"
//	@param			limit					query		int		false	"Control how many Charts are returned (default 100)"
//	@param			offset					query		int		false	"Control the offset for the returned Charts (default 0)"
//	@success		200						{array}		ChartV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/charts/v3 [get]
func chartsV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter ChartV3
	if err = ctx.ShouldBindQuery(&filter); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	modelFilter := filter.toModel()

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
	var results []models.Chart
	if err = db.
		Where(&modelFilter).
		Limit(limit).
		Offset(offset).
		Order("name asc").
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, chartFromModel))
}
