package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
	"strings"
)

// githubActionsJobsV3Get godoc
//
//	@summary		Get an individual GithubActionsJob
//	@description	Get an individual GithubActionsJob.
//	@tags			GithubActionsJobs
//	@produce		json
//	@param			selector				path		string	true	"The selector of the GithubActionsJob, either Sherlock ID or '{owner}/{repo}/{job ID}'"
//	@success		200						{object}	GithubActionsJobV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/github-actions-jobs/v3/{selector} [get]
func githubActionsJobsV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := githubActionsJobModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.GithubActionsJob
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, githubActionsJobFromModel(result))
}

func githubActionsJobModelFromSelector(selector string) (query models.GithubActionsJob, err error) {
	if len(selector) == 0 {
		return models.GithubActionsJob{}, fmt.Errorf("(%s) incident selector cannot be empty", errors.BadRequest)
	}
	if utils.IsNumeric(selector) { // ID
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else if strings.Count(selector, "/") == 2 { // owner + repo + job ID
		parts := strings.Split(selector, "/")
		query.GithubActionsOwner = parts[0]
		if query.GithubActionsOwner == "" {
			return models.GithubActionsJob{}, fmt.Errorf("(%s) invalid incident selector '%s', owner sub-selector was empty", errors.BadRequest, selector)
		}
		query.GithubActionsRepo = parts[1]
		if query.GithubActionsRepo == "" {
			return models.GithubActionsJob{}, fmt.Errorf("(%s) invalid incident selector '%s', repo sub-selector was empty", errors.BadRequest, selector)
		}
		query.GithubActionsJobID, err = utils.ParseUint(parts[2])
		if err != nil {
			return models.GithubActionsJob{}, fmt.Errorf("(%s) invalid incident selector '%s', job ID sub-selector was invalid", errors.BadRequest, selector)
		}
		return query, nil
	}
	return models.GithubActionsJob{}, fmt.Errorf("(%s) invalid incident selector '%s'", errors.BadRequest, selector)
}
