package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type SlackDeployHookV3Edit struct {
	deployHookTriggerConfigV3EditableFields
	SlackDeployHookFields
}

// slackDeployHooksV3Edit godoc
//
//	@summary		Edit an individual SlackDeployHook
//	@description	Edit an individual SlackDeployHook by its ID.
//	@tags			DeployHooks
//	@produce		json
//	@param			selector				path		string					true	"The ID of the SlackDeployHook to edit"
//	@param			slackDeployHook			body		SlackDeployHookV3Edit	true	"The edits to make to the SlackDeployHook"
//	@success		200						{object}	SlackDeployHookV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/slack/v3/{selector} [patch]
func slackDeployHooksV3Edit(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	query, err := slackDeployHookModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var body SlackDeployHookV3Edit
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	trigger, err := DeployHookTriggerConfigV3{deployHookTriggerConfigV3EditableFields: body.deployHookTriggerConfigV3EditableFields}.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	edits := models.SlackDeployHook{
		Trigger:       trigger,
		SlackChannel:  body.SlackChannel,
		MentionPeople: body.MentionPeople,
	}

	var toEdit models.SlackDeployHook
	if err = db.
		Preload("Trigger").
		Preload("Trigger.OnEnvironment").
		Preload("Trigger.OnChartRelease").
		Where(&query).
		First(&toEdit).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	editedChannel := body.SlackChannel != nil && toEdit.SlackChannel != nil && *body.SlackChannel != *toEdit.SlackChannel

	if err = db.Model(&toEdit).Omit(clause.Associations).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Model(&toEdit.Trigger).Omit(clause.Associations).Updates(&edits.Trigger).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if toEdit.SlackChannel != nil && editedChannel {
		var triggerDescription string
		if toEdit.Trigger.OnEnvironment != nil {
			triggerDescription = toEdit.Trigger.OnEnvironment.Name
		} else if toEdit.Trigger.OnChartRelease != nil {
			triggerDescription = toEdit.Trigger.OnChartRelease.Name
		}
		message := fmt.Sprintf("This channel is set to receive notifications for Beehive deployments to %s", triggerDescription)
		go slack.SendMessage(db.Statement.Context, *toEdit.SlackChannel, message, nil)
	}

	ctx.JSON(http.StatusOK, slackDeployHookFromModel(toEdit))
}
