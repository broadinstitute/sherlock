package sherlock

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// appVersionsV3Get godoc
//
//	@summary		Get an individual AppVersion
//	@description	Get an individual AppVersion.
//	@tags			AppVersions
//	@produce		json
//	@param			selector				path		string	true	"The selector of the AppVersion, which can be either a numeric ID or chart/version."
//	@success		200						{object}	AppVersionV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/app-versions/v3/{selector} [get]
func appVersionsV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := appVersionModelFromSelector(db, canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.AppVersion
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, appVersionFromModel(result))
}

func appVersionModelFromSelector(db *gorm.DB, selector string) (query models.AppVersion, err error) {
	if len(selector) == 0 {
		return models.AppVersion{}, fmt.Errorf("(%s) app version selector cannot be empty", errors.BadRequest)
	}
	if utils.IsNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return models.AppVersion{}, fmt.Errorf("(%s) string to int conversion error of '%s': %w", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.Count(selector, "/") == 1 { // chart + version
		parts := strings.Split(selector, "/")

		// chart
		chartSubQuery, err := chartModelFromSelector(parts[0])
		if err != nil {
			return models.AppVersion{}, fmt.Errorf("error handling chart sub-selector %s: %w", parts[0], err)
		}
		var chartSubResult models.Chart
		if err = db.Where(&chartSubQuery).First(&chartSubResult).Error; err != nil {
			return models.AppVersion{}, fmt.Errorf("error handling chart sub-selector %s: %w", parts[0], err)
		}
		query.ChartID = chartSubResult.ID

		// version
		version := parts[1]
		if len(version) == 0 {
			return models.AppVersion{}, fmt.Errorf("(%s) invalid app version selector %s, version sub-selector was empty", errors.BadRequest, selector)
		}
		query.AppVersion = version

		return query, nil
	}
	return models.AppVersion{}, fmt.Errorf("(%s) invalid app version selector '%s'", errors.BadRequest, selector)

}
