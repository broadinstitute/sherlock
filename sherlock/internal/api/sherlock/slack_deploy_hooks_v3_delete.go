package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// slackDeployHooksV3Delete godoc
//
//	@summary		Delete an individual SlackDeployHook
//	@description	Delete an individual SlackDeployHook by its ID.
//	@tags			DeployHooks
//	@produce		json
//	@param			selector				path		string	true	"The ID of the SlackDeployHook"
//	@success		200						{object}	SlackDeployHookV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/slack/v3/{selector} [delete]
func slackDeployHooksV3Delete(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := slackDeployHookModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if err = db.Select("Trigger").Delete(&query).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, slackDeployHookFromModel(query))
}
