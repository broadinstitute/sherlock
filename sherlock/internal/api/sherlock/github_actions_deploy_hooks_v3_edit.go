package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

type GithubActionsDeployHookV3Edit struct {
	deployHookTriggerConfigV3EditableFields
	GithubActionsDeployHookFields
}

// githubActionsDeployHooksV3Edit godoc
//
//	@summary		Edit an individual GithubActionsDeployHook
//	@description	Edit an individual GithubActionsDeployHook by its ID.
//	@tags			DeployHooks
//	@produce		json
//	@param			selector				path		string							true	"The ID of the GithubActionsDeployHook to edit"
//	@param			githubActionsDeployHook	body		GithubActionsDeployHookV3Edit	true	"The edits to make to the GithubActionsDeployHook"
//	@success		200						{object}	GithubActionsDeployHookV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/github-actions/v3/{selector} [patch]
func githubActionsDeployHooksV3Edit(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	query, err := githubActionsDeployHookModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var body GithubActionsDeployHookV3Edit
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	trigger, err := DeployHookTriggerConfigV3{deployHookTriggerConfigV3EditableFields: body.deployHookTriggerConfigV3EditableFields}.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	edits := models.GithubActionsDeployHook{
		Trigger:                     trigger,
		GithubActionsOwner:          body.GithubActionsOwner,
		GithubActionsRepo:           body.GithubActionsRepo,
		GithubActionsWorkflowPath:   body.GithubActionsWorkflowPath,
		GithubActionsDefaultRef:     body.GithubActionsDefaultRef,
		GithubActionsRefBehavior:    body.GithubActionsRefBehavior,
		GithubActionsWorkflowInputs: body.GithubActionsWorkflowInputs,
	}

	var toEdit models.GithubActionsDeployHook
	if err = db.
		Preload("Trigger").
		Preload("Trigger.OnEnvironment").
		Preload("Trigger.OnChartRelease").
		Where(&query).
		First(&toEdit).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Model(&toEdit).Omit(clause.Associations).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Model(&toEdit.Trigger).Omit(clause.Associations).Updates(&edits.Trigger).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, githubActionsDeployHookFromModel(toEdit))
}
