package sherlock

import (
	"encoding/json"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/github"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GithubActionsDeployHookTestRunRequest struct {
	Execute *bool `json:"execute"` // Required, whether to fully run the GHA
}

type GithubActionsDeployHookTestRunResponse struct {
	OK  bool   `json:"ok"`
	URL string `json:"url,omitempty"`
}

// githubActionsDeployHooksV3TestRun godoc
//
//	@summary		Test a GithubActionsDeployHook
//	@description	Run a GitHub Action to simulate a GithubActionsDeployHook
//	@tags			DeployHooks
//	@produce		json
//	@param			selector				path		string									true	"The ID of the GithubActionsDeployHook"
//	@param			request					body		GithubActionsDeployHookTestRunRequest	true	"Whether to fully execute the hook (JSON body helps with CSRF protection)"
//	@success		200						{object}	GithubActionsDeployHookTestRunResponse
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/deploy-hooks/github-actions/procedures/v3/test/{selector} [post]
func githubActionsDeployHooksV3TestRun(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body GithubActionsDeployHookTestRunRequest
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	} else if body.Execute == nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) 'execute' is required", errors.BadRequest))
		return
	}

	query, err := githubActionsDeployHookModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var hook models.GithubActionsDeployHook
	if err = db.
		Where(&query).
		First(&hook).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var workflowInputs map[string]any
	if hook.GithubActionsWorkflowInputs != nil {
		if bytes, err := hook.GithubActionsWorkflowInputs.MarshalJSON(); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error marshalling inputs: %w", err))
			return
		} else if err = json.Unmarshal(bytes, &workflowInputs); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error unmarshalling stored inputs: %w", err))
			return
		}
	}

	if !*body.Execute {
		ctx.JSON(http.StatusOK, GithubActionsDeployHookTestRunResponse{OK: true})
		return
	}

	if err = github.DispatchWorkflow(ctx,
		*hook.GithubActionsOwner,
		*hook.GithubActionsRepo,
		*hook.GithubActionsWorkflowPath,
		*hook.GithubActionsDefaultRef,
		workflowInputs); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error between Sherlock and GitHub: %w", err))
		return
	}
	ctx.JSON(http.StatusOK, GithubActionsDeployHookTestRunResponse{
		OK:  true,
		URL: fmt.Sprintf("https://github.com/%s/%s/actions", *hook.GithubActionsOwner, *hook.GithubActionsRepo),
	})
}
