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

// githubActionsJobsV3List godoc
//
//	@summary		List GithubActionsJobs matching a filter
//	@description	List GithubActionsJobs matching a filter.
//	@description	Results are ordered by start time, starting at most recent.
//	@tags			GithubActionsJobs
//	@produce		json
//	@param			filter					query		GithubActionsJobV3	false	"Filter the returned GithubActionsJobs"
//	@param			limit					query		int					false	"Control how many GithubActionsJobs are returned (default 100)"
//	@param			offset					query		int					false	"Control the offset for the returned GithubActionsJobs (default 0)"
//	@success		200						{array}		GithubActionsJobV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/github-actions-jobs/v3 [get]
func githubActionsJobsV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter GithubActionsJobV3
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
	var results []models.GithubActionsJob
	if err = db.Where(&modelFilter).Limit(limit).Offset(offset).Order("job_created_at desc").Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, githubActionsJobFromModel))
}
