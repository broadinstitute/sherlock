package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// chartsV3Get godoc
//
//	@summary		Get an individual Chart
//	@description	Get an individual Chart.
//	@tags			Charts
//	@produce		json
//	@param			selector				path		string	true	"The selector of the Chart, which can be either a numeric ID or the name."
//	@success		200						{object}	ChartV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/charts/v3/{selector} [get]
func chartsV3Get(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, chartFromModel(result))
}

func chartModelFromSelector(selector string) (query models.Chart, err error) {
	if len(selector) == 0 {
		return models.Chart{}, fmt.Errorf("(%s) chart selector cannot be empty", errors.BadRequest)
	}
	if utils.IsNumeric(selector) { // ID
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else if utils.IsAlphaNumericWithHyphens(selector) &&
		utils.IsStartingWithLetter(selector) &&
		utils.IsEndingWithAlphaNumeric(selector) { // Name
		query.Name = selector
		return query, nil
	}
	return models.Chart{}, fmt.Errorf("(%s) invalid chart selector '%s'", errors.BadRequest, selector)
}
