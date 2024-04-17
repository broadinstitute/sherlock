package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/bee"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// environmentsProceduresV3UpsertBee godoc
//
//	@summary			Get a Bee, new or existing
//	@description	Creates a Bee Environment, or retreives one that already exists.
//	@tags					Environments
//	@accept				json
//	@produce			json
//	@param				environment				body		EnvironmentV3Create	true	"The Environment to create"
//	@success			201						{object}	EnvironmentV3
//	@failure			400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router				/api/environments/procedures/v3/upsert-bee [put]
func environmentsProceduresV3UpsertBee(ctx *gin.Context) {
	// force auth
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	// Environment Parsing
	// make sure we can map the JSON to the go struct
	//var parsedBody EnvironmentProceduresV3UpsertBee

	var environmentBody EnvironmentV3Create
	// ShouldBindBodyWith used to handle double-reading body
	if err = ctx.ShouldBindBodyWith(&environmentBody, binding.JSON); err != nil {
		//if err = ctx.ShouldBindJSON(&environmentBody); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	//environmentBody := parsedBody.EnvironmentV3Create
	//changesetBody := parsedBody.ChangesetV3PlanRequest

	// set default values
	if err = defaults.Set(&environmentBody); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	// convert the body to the db model
	toCreate, err := environmentBody.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	// changesetParsing
	var changesetBody ChangesetV3PlanRequest
	if err = ctx.ShouldBindBodyWith(&changesetBody, binding.JSON); err != nil {
		//if err = ctx.ShouldBindJSON(&changesetBody); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) JSON error parsing body to %T: %w", errors.BadRequest, changesetBody, err))
		return
	}

	// copypasta from plan, abstract this at some point.
	var chartReleaseChangesets, environmentChangesets, recreateChangesets []models.Changeset

	if chartReleaseChangesets, err = changesetBody.parseChartReleaseEntries(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error handling chart release entries: %w", err))
		return
	}
	if environmentChangesets, err = changesetBody.parseEnvironmentEntries(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error handling environment entries: %w", err))
		return
	}
	if recreateChangesets, err = changesetBody.parseRecreateEntries(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error handling recreate entries: %w", err))
		return

	}
	beeEdits := append(append(chartReleaseChangesets, environmentChangesets...), recreateChangesets...)

	// do the thing
	beeModel, err := bee.BeeUpsert(toCreate, beeEdits, db)

	// populate &result from the db
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	// shove it back into JSON
	beeJSON := environmentFromModel(beeModel)
	ctx.JSON(http.StatusCreated, beeJSON)
}
