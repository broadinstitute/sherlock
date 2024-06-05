package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// chartsV3Create godoc
//
//	@summary		Create a Chart
//	@description	Create a Chart.
//	@tags			Charts
//	@accept			json
//	@produce		json
//	@param			chart					body		ChartV3Create	true	"The Chart to create"
//	@success		201						{object}	ChartV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/charts/v3 [post]
func chartsV3Create(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body ChartV3Create
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	if body.DefaultSubdomain == nil {
		body.DefaultSubdomain = utils.PointerTo(body.Name)
	}

	toCreate := body.toModel()
	if err = db.Create(&toCreate).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var result models.Chart
	if err = db.Preload(clause.Associations).First(&result, toCreate.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, chartFromModel(result))
}
