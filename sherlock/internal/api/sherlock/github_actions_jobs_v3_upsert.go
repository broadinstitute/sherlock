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

// githubActionsJobsV3Upsert godoc
//
//	@summary		Upsert GithubActionsJob
//	@description	Upsert GithubActionsJob.
//	@tags			GithubActionsJobs
//	@produce		json
//	@param			githubActionsJob		body		GithubActionsJobV3Create	true	"The GithubActionsJob to upsert"
//	@success		201						{object}	GithubActionsJobV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/github-actions-jobs/v3 [put]
func githubActionsJobsV3Upsert(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var body GithubActionsJobV3Create
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	toUpsert := body.toModel()
	if err = db.Where(&models.GithubActionsJob{
		GithubActionsOwner: toUpsert.GithubActionsOwner,
		GithubActionsRepo:  toUpsert.GithubActionsRepo,
		GithubActionsJobID: toUpsert.GithubActionsJobID,
	}).Assign(&models.GithubActionsJob{
		// GitHub technically just identifies by owner, repo, and job ID, so we assign this data rather than filtering on it
		GithubActionsRunID:         toUpsert.GithubActionsRunID,
		GithubActionsAttemptNumber: toUpsert.GithubActionsAttemptNumber,
		JobCreatedAt:               toUpsert.JobCreatedAt,
		JobStartedAt:               toUpsert.JobStartedAt,
		JobTerminalAt:              toUpsert.JobTerminalAt,
		Status:                     toUpsert.Status,
	}).FirstOrCreate(&toUpsert).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var result models.GithubActionsJob
	if err = db.Preload(clause.Associations).First(&result, toUpsert.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, githubActionsJobFromModel(result))
}
