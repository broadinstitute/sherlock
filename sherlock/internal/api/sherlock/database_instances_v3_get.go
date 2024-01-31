package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

// databaseInstancesV3Get godoc
//
//	@summary		Get an individual DatabaseInstance
//	@description	Get an individual DatabaseInstance by its selector.
//	@tags			DatabaseInstances
//	@produce		json
//	@param			selector				path		string	true	"The selector of the DatabaseInstance, which can be either a numeric ID or 'chart-release/' followed by a chart release selector."
//	@success		200						{object}	DatabaseInstanceV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/database-instances/v3/{selector} [get]
func databaseInstancesV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := databaseInstanceModelFromSelector(db, canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.DatabaseInstance
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, databaseInstanceFromModel(result))
}

func databaseInstanceModelFromSelector(db *gorm.DB, selector string) (query models.DatabaseInstance, err error) {
	if len(selector) == 0 {
		return models.DatabaseInstance{}, fmt.Errorf("(%s) database instance selector cannot be empty", errors.BadRequest)
	}
	if utils.IsNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return models.DatabaseInstance{}, fmt.Errorf("(%s) string to int conversion error of '%s': %w", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.HasPrefix(selector, "chart-release/") { // 'chart-release' + chart release selector
		chartReleaseSubSelector := strings.TrimPrefix(selector, "chart-release/")
		chartReleaseSubQuery, err := chartReleaseModelFromSelector(db, chartReleaseSubSelector)
		if err != nil {
			return models.DatabaseInstance{}, fmt.Errorf("error handling chart release sub-selector %s: %w", chartReleaseSubSelector, err)
		}
		var chartReleaseSubResult models.ChartRelease
		if err = db.Where(&chartReleaseSubQuery).Select("id").First(&chartReleaseSubResult).Error; err != nil {
			return models.DatabaseInstance{}, fmt.Errorf("error querying for chart release sub-selector %s: %w", chartReleaseSubSelector, err)
		}
		query.ChartReleaseID = chartReleaseSubResult.ID
		return query, nil
	}
	return models.DatabaseInstance{}, fmt.Errorf("(%s) invalid database instance selector '%s'", errors.BadRequest, selector)
}
