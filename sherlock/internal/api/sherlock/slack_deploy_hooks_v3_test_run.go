package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SlackDeployHookTestRunRequest struct {
	Execute *bool `json:"execute"` // Required, whether to actually send the Slack message
}

type SlackDeployHookTestRunResponse struct {
	OK bool `json:"ok"`
}

// slackDeployHooksV3TestRun godoc
//
//	@summary		Test a SlackDeployHook
//	@description	Send a Slack message to simulate a SlackDeployHook
//	@tags			DeployHooks
//	@produce		json
//	@param			selector				path		string							true	"The ID of the SlackDeployHook to test"
//	@param			request					body		SlackDeployHookTestRunRequest	true	"Whether to fully execute the hook (JSON body helps with CSRF protection)"
//	@success		200						{object}	SlackDeployHookTestRunResponse
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/slack/procedures/v3/test/{selector} [post]
func slackDeployHooksV3TestRun(ctx *gin.Context) {
	user, err := authentication.MustUseUser(ctx)
	if err != nil {
		return
	}
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body SlackDeployHookTestRunRequest
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	} else if body.Execute == nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) 'execute' is required", errors.BadRequest))
		return
	}

	query, err := slackDeployHookModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var hook models.SlackDeployHook
	if err = db.
		Where(&query).
		First(&hook).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if !*body.Execute {
		ctx.JSON(http.StatusOK, SlackDeployHookTestRunResponse{OK: true})
		return
	}
	if err = slack.SendMessageReturnError(ctx, *hook.SlackChannel,
		fmt.Sprintf("This is a deploy hook test message from Beehive, triggered by %s", user.SlackReference(true)), nil); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error between Sherlock and Slack: %w", err))
		return
	}
	ctx.JSON(http.StatusOK, SlackDeployHookTestRunResponse{OK: true})
}
