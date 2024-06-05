package sherlock

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"math"
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

	ctx.JSON(http.StatusCreated, githubActionsJobFromModel(toUpsert))

	// We're done with the actual request and would be good to return, but let's fire
	// off some metrics if the job completed
	if toUpsert.JobTerminalAt != nil && !toUpsert.JobTerminalAt.IsZero() &&
		toUpsert.Status != nil && *toUpsert.Status != "skipped" {

		// If the job has a completed timestamp and wasn't skipped, calculate how long
		// it was queued and how long it was running
		var queuedSeconds, runningSeconds int64
		if toUpsert.JobCreatedAt != nil && !toUpsert.JobCreatedAt.IsZero() &&
			toUpsert.JobStartedAt != nil && !toUpsert.JobStartedAt.IsZero() {
			queuedSeconds = int64(math.Ceil(toUpsert.JobStartedAt.Sub(*toUpsert.JobCreatedAt).Seconds()))
		}
		if toUpsert.JobStartedAt != nil && !toUpsert.JobStartedAt.IsZero() &&
			toUpsert.JobTerminalAt != nil && !toUpsert.JobTerminalAt.IsZero() {
			runningSeconds = int64(math.Ceil(toUpsert.JobTerminalAt.Sub(*toUpsert.JobStartedAt).Seconds()))
		}

		// If either of those times are non-zero, we'll report both as metrics
		// (GHA internals are complex, and a good way to understand if a job
		// actually happened is to check if time elapsed during it)
		if queuedSeconds > 0 || runningSeconds > 0 {

			// We'll make a new context for the tags and add the base ones
			var tagCtx context.Context
			if tagCtx, err = tag.New(ctx,
				tag.Upsert(metrics.GithubActionsOrganizationKey, toUpsert.GithubActionsOwner),
				tag.Upsert(metrics.GithubActionsRepoKey, fmt.Sprintf("%s/%s", toUpsert.GithubActionsOwner, toUpsert.GithubActionsRepo))); err != nil {
				log.Error().Err(err).Msg("unable to create tag context")
				return
			}

			// It's potentially helpful to include the file of the workflow that
			// defined the job. We don't have that info here, but the CiRun table
			// might, if it already got a webhook payload for the overall workflow.
			if toUpsert.GithubActionsRunID != 0 && toUpsert.GithubActionsAttemptNumber != 0 {
				var matchingWorkflowFiles []string
				if err = db.
					Model(&models.CiRun{}).
					Where(&models.CiRun{
						GithubActionsRunID:         toUpsert.GithubActionsRunID,
						GithubActionsAttemptNumber: toUpsert.GithubActionsAttemptNumber,
					}).
					Limit(1).
					Pluck("github_actions_workflow_path", &matchingWorkflowFiles).Error; err != nil {
					log.Error().Err(err).Msg("errored finding matching workflow runs for job")
					return
				} else if len(matchingWorkflowFiles) > 0 {
					// We have a match for the workflow file, so we'll mutate the tag
					// context to include it
					if tagCtx, err = tag.New(tagCtx,
						tag.Upsert(metrics.GithubActionsWorkflowFileKey, matchingWorkflowFiles[0])); err != nil {
						log.Error().Err(err).Msg("unable to create tag context")
						return
					}
				}
			}

			// Send off the queued time
			if tagCtx, err = tag.New(tagCtx,
				tag.Upsert(metrics.GithubActionsJobStageKey, "queued")); err != nil {
				log.Error().Err(err).Msg("unable to create tag context")
				return
			} else {
				stats.Record(tagCtx, metrics.GithubActionsJobLatencyMeasure.M(queuedSeconds))
			}

			// Send off the running time
			if tagCtx, err = tag.New(tagCtx,
				tag.Upsert(metrics.GithubActionsJobStageKey, "running")); err != nil {
				log.Error().Err(err).Msg("unable to create tag context")
				return
			} else {
				stats.Record(tagCtx, metrics.GithubActionsJobLatencyMeasure.M(runningSeconds))
			}
		}
	}
}
