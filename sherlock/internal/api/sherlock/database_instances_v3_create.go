package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

// databaseInstancesV3Create godoc
//
//	@summary		Create a DatabaseInstance
//	@description	Create a DatabaseInstance.
//	@tags			DatabaseInstances
//	@accept			json
//	@produce		json
//	@param			databaseInstance		body		DatabaseInstanceV3Create	true	"The DatabaseInstance to create"
//	@success		201						{object}	DatabaseInstanceV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/database-instances/v3 [post]
func databaseInstancesV3Create(ctx *gin.Context) {
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

	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	toCreate, err := body.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	// We'll error regardless but we can be explicit with a good error message by just checking here
	if toCreate.ChartReleaseID == 0 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) chartRelease is required", errors.BadRequest))
		return
	}

	if toCreate.DefaultDatabase == nil {
		var chart models.Chart
		if err = db.
			Where("id = (?)", db.
				Model(&models.ChartRelease{Model: gorm.Model{ID: toCreate.ChartReleaseID}}).
				Select("chart_id").
				Limit(1)).
			Select("name").
			Take(&chart).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("unable to get chart name to fill default database: %w", err))
			return
		}
		toCreate.DefaultDatabase = &chart.Name
	}

	if err = db.Create(&toCreate).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var result models.DatabaseInstance
	if err = db.Preload(clause.Associations).First(&result, toCreate.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, databaseInstanceFromModel(result))
}
