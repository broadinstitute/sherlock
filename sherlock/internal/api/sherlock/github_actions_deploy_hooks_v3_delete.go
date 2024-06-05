package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// githubActionsDeployHooksV3Delete godoc
//
//	@summary		Delete an individual GithubActionsDeployHook
//	@description	Delete an individual GithubActionsDeployHook by its ID.
//	@tags			DeployHooks
//	@produce		json
//	@param			selector				path		string	true	"The ID of the GithubActionsDeployHook"
//	@success		200						{object}	GithubActionsDeployHookV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/github-actions/v3/{selector} [delete]
func githubActionsDeployHooksV3Delete(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := githubActionsDeployHookModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.GithubActionsDeployHook
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
	ctx.JSON(http.StatusOK, githubActionsDeployHookFromModel(result))
}
