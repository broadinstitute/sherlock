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

// environmentsV3Edit godoc
//
//	@summary		Edit an individual Environment
//	@description	Edit an individual Environment.
//	@tags			Environments
//	@produce		json
//	@param			selector				path		string			true	"The selector of the Environment, which can be either a numeric ID, the name, or 'resource-prefix' + / + the unique resource prefix."
//	@param			environment					body		EnvironmentV3Edit	true	"The edits to make to the Environment"
//	@success		200						{object}	EnvironmentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/environments/v3/{selector} [patch]
func environmentsV3Edit(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := environmentModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var body EnvironmentV3Edit
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	edits, err := body.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var toEdit models.Environment
	if err = db.Preload(clause.Associations).Where(&query).First(&toEdit).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Model(&toEdit).Omit(clause.Associations).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, environmentFromModel(toEdit))
}
