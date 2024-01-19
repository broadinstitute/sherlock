package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
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
	var result models.SlackDeployHook
	if err = db.
		Preload("Trigger").
		Preload("Trigger.OnEnvironment").
		Preload("Trigger.OnChartRelease").
		Where(&query).
		First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if err = db.Select("Trigger").Delete(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if result.SlackChannel != nil {
		if result.Trigger.OnEnvironment != nil {
			go slack.SendMessage(db.Statement.Context, *result.SlackChannel, fmt.Sprintf("This channel will no longer receive notifications for Beehive deployments in %s", result.Trigger.OnEnvironment.Name), nil)
		} else if result.Trigger.OnChartRelease != nil {
			go slack.SendMessage(db.Statement.Context, *result.SlackChannel, fmt.Sprintf("This channel will no longer receive notifications for Beehive deployments to %s", result.Trigger.OnChartRelease.Name), nil)
		}
	}
	ctx.JSON(http.StatusOK, slackDeployHookFromModel(result))
}
