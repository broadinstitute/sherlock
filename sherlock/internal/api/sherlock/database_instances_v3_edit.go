package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm/clause"
	"net/http"
)

// databaseInstancesV3Edit godoc
//
//	@summary		Edit an individual DatabaseInstance
//	@description	Edit an individual DatabaseInstance by its selector.
//	@tags			DatabaseInstances
//	@produce		json
//	@param			selector				path		string				true	"The selector of the DatabaseInstance, which can be either a numeric ID or 'chart-release/' followed by a chart release selector."
//	@param			databaseInstance			body		DatabaseInstanceV3Edit	true	"The edits to make to the DatabaseInstance"
//	@success		200						{object}	DatabaseInstanceV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/database-instances/v3/{selector} [patch]
func databaseInstancesV3Edit(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := databaseInstanceModelFromSelector(db, canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var body DatabaseInstanceV3Edit
	// ShouldBindBodyWith used to handle double-reading body when redirected from databaseInstancesV3Upsert
	if err = ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	edits, err := body.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var toEdit models.DatabaseInstance
	if err = db.Preload(clause.Associations).Where(&query).First(&toEdit).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Model(&toEdit).Omit(clause.Associations).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, databaseInstanceFromModel(toEdit))
}
