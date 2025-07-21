package sherlock

import (
	"fmt"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// databaseInstancesV3Upsert godoc
//
//	@summary		Create or edit a DatabaseInstance
//	@description	Create or edit a DatabaseInstance, depending on whether one already exists for the chart release
//	@tags			DatabaseInstances
//	@accept			json
//	@produce		json
//	@param			databaseInstance		body		DatabaseInstanceV3Create	true	"The DatabaseInstance to create or edit. Defaults will only be set if creating."
//	@success		200,201					{object}	DatabaseInstanceV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/database-instances/v3 [put]
func databaseInstancesV3Upsert(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body DatabaseInstanceV3Create
	// ShouldBindBodyWith used to handle double-reading body when redirected from databaseInstancesV3Upsert
	if err = ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	toUpsert, err := body.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if toUpsert.ChartReleaseID == 0 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) chartRelease is required", errors.BadRequest))
		return
	}

	var existing []models.DatabaseInstance
	if err = db.Where(&models.DatabaseInstance{ChartReleaseID: toUpsert.ChartReleaseID}).Select("id").Find(&existing).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	} else if len(existing) >= 1 {
		// If it exists, redirect to the edit method and set the selector param as if it had been passed
		ctx.AddParam("selector", utils.UintToString(existing[0].ID))
		databaseInstancesV3Edit(ctx)
	} else {
		// Otherwise, redirect to the create method
		databaseInstancesV3Create(ctx)
	}
}
