package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
)

// slackDeployHooksV3List godoc
//
//	@summary		List SlackDeployHooks matching a filter
//	@description	List SlackDeployHooks matching a filter.
//	@tags			DeployHooks
//	@produce		json
//	@param			filter					query		SlackDeployHookV3	false	"Filter the returned SlackDeployHooks"
//	@param			limit					query		int					false	"Control how many SlackDeployHooks are returned (default 100)"
//	@param			offset					query		int					false	"Control the offset for the returned SlackDeployHooks (default 0)"
//	@success		200						{array}		SlackDeployHookV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/slack/v3 [get]
func slackDeployHooksV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter SlackDeployHookV3
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
	var results []models.SlackDeployHook
	if err = db.
		InnerJoins("Trigger", db.Where(&modelFilter.Trigger)).
		Where(&modelFilter).
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, slackDeployHookFromModel))
}
