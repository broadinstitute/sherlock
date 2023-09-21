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

// chartsV3Edit godoc
//
//	@summary		Edit an individual Chart
//	@description	Edit an individual Chart.
//	@tags			Charts
//	@produce		json
//	@param			selector				path		string		true	"The selector of the Chart, which can be either a numeric ID or the name."
//	@param			chart					body		ChartV3Edit	true	"The edits to make to the Chart"
//	@success		200						{object}	ChartV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/charts/v3/{selector} [patch]
func chartsV3Edit(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := chartModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var body ChartV3Edit
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	edits := body.toModel()

	var toEdit models.Chart
	if err = db.Preload(clause.Associations).Where(&query).First(&toEdit).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Model(&toEdit).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, chartFromModel(toEdit))
}
