package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// chartsV3Delete godoc
//	@summary		Delete an individual Chart
//	@description	Delete an individual Chart by its ID.
//	@tags			Charts
//	@produce		json
//	@param			selector				path		string	true	"The selector of the Chart, which can be either a numeric ID or the name."
//	@success		200						{object}	ChartV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/charts/v3/{selector} [delete]

func chartsV3Delete(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := chartModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.Chart
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if err = db.Delete(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, chartFromModel(result))
}
