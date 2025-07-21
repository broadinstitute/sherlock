package sherlock

import (
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
)

// ciRunsProceduresV3GithubInfoList godoc
//
//	@summary		List GitHub info gleaned from CiRuns
//	@description	List info about GitHub repos and their workflow files as determined by CiRuns from the past 90 days.
//	@description	This is a useful proxy for figuring out what repos Sherlock probably has access to: workflows listed
//	@description	here can probably successfully called by a GitHub Actions deploy hook.
//	@tags			CiRuns
//	@produce		json
//	@success		200						{object}	map[string]map[string][]string
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/ci-runs/procedures/v3/github-info [get]
func ciRunsProceduresV3GithubInfoList(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	type result struct {
		GithubActionsOwner        string
		GithubActionsRepo         string
		GithubActionsWorkflowPath string
	}
	var results []result
	if err = db.Model(&models.CiRun{}).
		Where("created_at >= current_timestamp - '90 days'::interval").
		Distinct("github_actions_owner", "github_actions_repo", "github_actions_workflow_path").
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	toReturn := make(map[string]map[string][]string)
	for _, r := range results {
		org, orgExists := toReturn[r.GithubActionsOwner]
		if !orgExists {
			org = make(map[string][]string)
			toReturn[r.GithubActionsOwner] = org
		}
		org[r.GithubActionsRepo] = append(org[r.GithubActionsRepo], r.GithubActionsWorkflowPath)
	}
	ctx.JSON(http.StatusOK, toReturn)
}
