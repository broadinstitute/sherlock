package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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

	// make sure we can map the JSON to the go struct
	var body EnvironmentV3Create
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	// set default values
	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	// convert the body to the db model
	toCreate, err := body.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	// do the thing
	if err = db.Create(&toCreate).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	// populate &result from the db
	var result models.Environment
	if err = db.Preload(clause.Associations).First(&result, toCreate.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	// shove it back into JSON
	ctx.JSON(http.StatusCreated, environmentFromModel(result))
}
