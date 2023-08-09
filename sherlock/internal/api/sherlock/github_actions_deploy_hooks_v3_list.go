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

// githubActionsDeployHooksV3List godoc
//
//	@summary		List GithubActionsDeployHooks matching a filter
//	@description	List GithubActionsDeployHooks matching a filter.
//	@tags			DeployHooks
//	@produce		json
//	@param			filter					query		GithubActionsDeployHookV3	false	"Filter the returned GithubActionsDeployHooks"
//	@param			limit					query		int							false	"Control how many GithubActionsDeployHooks are returned (default 100)"
//	@param			offset					query		int							false	"Control the offset for the returned GithubActionsDeployHooks (default 0)"
//	@success		200						{array}		GithubActionsDeployHookV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/github-actions/v3 [get]
func githubActionsDeployHooksV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter GithubActionsDeployHookV3
	if err = ctx.ShouldBindQuery(&filter); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	modelFilter, err := filter.toModel(db)
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
	var results []models.GithubActionsDeployHook
	if err = db.Preload(clause.Associations).Where(&modelFilter).Limit(limit).Offset(offset).Order("created_at desc").Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, githubActionsDeployHookFromModel))
}
