package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SlackDeployHookV3Create struct {
	DeployHookTriggerConfigV3
	SlackDeployHookFields
}

// slackDeployHooksV3Create godoc
//
//	@summary		Create a SlackDeployHook
//	@description	Create a SlackDeployHook.
//	@tags			DeployHooks
//	@accept			json
//	@produce		json
//	@param			slackDeployHook			body		SlackDeployHookV3Create	true	"The SlackDeployHook to create"
//	@success		201						{object}	SlackDeployHookV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/slack/v3 [post]
func slackDeployHooksV3Create(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body SlackDeployHookV3Create
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %v", errors.BadRequest, err))
		return
	}

	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %v", err))
		return
	}

	trigger, err := body.DeployHookTriggerConfigV3.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	hook := models.SlackDeployHook{
		Trigger:      trigger,
		SlackChannel: body.SlackChannel,
	}
	if err = db.Create(&hook).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.SlackDeployHook
	if err = db.
		Preload("Trigger").
		Preload("Trigger.OnEnvironment").
		Preload("Trigger.OnChartRelease").
		Where(&hook).
		First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if result.SlackChannel != nil {
		if result.Trigger.OnEnvironment != nil {
			go slack.SendMessage(db.Statement.Context, *result.SlackChannel, fmt.Sprintf("This channel will now receive notifications for Beehive deployments in %s", result.Trigger.OnEnvironment.Name))
		} else if result.Trigger.OnChartRelease != nil {
			go slack.SendMessage(db.Statement.Context, *result.SlackChannel, fmt.Sprintf("This channel will now receive notifications for Beehive deployments to %s", result.Trigger.OnChartRelease.Name))
		}
	}
	ctx.JSON(http.StatusCreated, slackDeployHookFromModel(result))
}
