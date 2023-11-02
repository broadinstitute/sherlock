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

type SlackDeployHookTestRunResponse struct {
	OK bool `json:"ok"`
}

// slackDeployHooksV3TestRun godoc
//
//	@summary		Test a SlackDeployHook
//	@description	Send a Slack message to simulate a SlackDeployHook
//	@tags			DeployHooks
//	@produce		json
//	@param			selector				path		string	true	"The ID of the SlackDeployHook to test"
//	@success		200						{object}	SlackDeployHookTestResponse
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
	if err = slack.SendMessageReturnError(ctx, *hook.SlackChannel,
		fmt.Sprintf("This is a deploy hook test message from Beehive, triggered by %s", user.SlackReference())); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, SlackDeployHookTestRunResponse{OK: true})
}
