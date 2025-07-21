package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
)

type GithubActionsDeployHookV3Create struct {
	DeployHookTriggerConfigV3
	GithubActionsDeployHookFields
}

// githubActionsDeployHooksV3Create godoc
//
//	@summary		Create a GithubActionsDeployHook
//	@description	Create a GithubActionsDeployHook.
//	@tags			DeployHooks
//	@accept			json
//	@produce		json
//	@param			githubActionsDeployHook	body		GithubActionsDeployHookV3Create	true	"The GithubActionsDeployHook to create"
//	@success		201						{object}	GithubActionsDeployHookV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/github-actions/v3 [post]
func githubActionsDeployHooksV3Create(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body GithubActionsDeployHookV3Create
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	trigger, err := body.DeployHookTriggerConfigV3.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	hook := models.GithubActionsDeployHook{
		Trigger:                     trigger,
		GithubActionsOwner:          body.GithubActionsOwner,
		GithubActionsRepo:           body.GithubActionsRepo,
		GithubActionsWorkflowPath:   body.GithubActionsWorkflowPath,
		GithubActionsDefaultRef:     body.GithubActionsDefaultRef,
		GithubActionsRefBehavior:    body.GithubActionsRefBehavior,
		GithubActionsWorkflowInputs: body.GithubActionsWorkflowInputs,
	}
	if err = db.Create(&hook).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.GithubActionsDeployHook
	if err = db.
		Preload("Trigger").
		Preload("Trigger.OnEnvironment").
		Preload("Trigger.OnChartRelease").
		Where(&hook).
		First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, githubActionsDeployHookFromModel(result))
}
