package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// ciRunsV3List godoc
//
//	@summary		List CiRuns matching a filter
//	@description	List CiRuns matching a filter. The CiRuns would have to re-queried directly to load any related resources.
//	@description	Results are ordered by start time, starting at most recent.
//	@tags			CiRuns
//	@produce		json
//	@param			filter					query		CiRunV3	false	"Filter the returned CiRuns"
//	@param			limit					query		int		false	"Control how many CiRuns are returned (default 100)"
//	@param			offset					query		int		false	"Control the offset for the returned CiRuns (default 0)"
//	@success		200						{array}		CiRunV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/ci-runs/v3 [get]
func ciRunsV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter CiRunV3
	if err = ctx.MustBindWith(&filter, binding.Query); err != nil {
		return
	}
	modelFilter := filter.toModel()
	limit, err := utils.ParseInt(ctx.DefaultQuery("limit", "100"))
	if err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) %v", errors.BadRequest, err)))
		return
	}
	offset, err := utils.ParseInt(ctx.DefaultQuery("offset", "0"))
	if err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(fmt.Errorf("(%s) %v", errors.BadRequest, err)))
		return
	}
	var results []models.CiRun
	if err = db.Where(&modelFilter).Limit(limit).Offset(offset).Order("started_at desc").Find(&results).Error; err != nil {
		ctx.AbortWithStatusJSON(errors.ErrorToApiResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, ciRunFromModel))
}
