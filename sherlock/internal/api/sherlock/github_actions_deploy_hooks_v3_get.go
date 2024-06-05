package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// githubActionsDeployHooksV3Get godoc
//
//	@summary		Get an individual GithubActionsDeployHook
//	@description	Get an individual GithubActionsDeployHook by its ID.
//	@tags			DeployHooks
//	@produce		json
//	@param			selector				path		string	true	"The ID of the GithubActionsDeployHook"
//	@success		200						{object}	GithubActionsDeployHookV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/github-actions/v3/{selector} [get]
func githubActionsDeployHooksV3Get(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, githubActionsDeployHookFromModel(result))
}

func githubActionsDeployHookModelFromSelector(selector string) (query models.GithubActionsDeployHook, err error) {
	if len(selector) == 0 {
		return models.GithubActionsDeployHook{}, fmt.Errorf("(%s) selector cannot be empty", errors.BadRequest)
	} else if utils.IsNumeric(selector) {
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else {
		return models.GithubActionsDeployHook{}, fmt.Errorf("(%s) invalid githubActions deploy hook selector '%s'", errors.BadRequest, selector)
	}
}
