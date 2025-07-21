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

// chartReleasesV3Get godoc
//
//	@summary		Get an individual ChartRelease
//	@description	Get an individual ChartRelease.
//	@tags			ChartReleases
//	@produce		json
//	@param			selector				path		string	true	"The selector of the ChartRelease, which can be either a numeric ID, the name, environment + '/' + chart, or cluster + '/' + namespace + '/' + chart."
//	@success		200						{object}	ChartReleaseV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/chart-releases/v3/{selector} [get]
func chartReleasesV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := chartReleaseModelFromSelector(db, canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.ChartRelease
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, chartReleaseFromModel(result))
}

func chartReleaseModelFromSelector(db *gorm.DB, selector string) (query models.ChartRelease, err error) {
	if len(selector) == 0 {
		return models.ChartRelease{}, fmt.Errorf("(%s) chart release selector cannot be empty", errors.BadRequest)
	}
	if utils.IsNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return models.ChartRelease{}, fmt.Errorf("(%s) string to int conversion error of '%s': %w", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.Count(selector, "/") == 1 { // environment + chart
		parts := strings.Split(selector, "/")

		// environment
		environmentSubQuery, err := environmentModelFromSelector(parts[0])
		if err != nil {
			return models.ChartRelease{}, fmt.Errorf("error handling environment sub-selector %s: %w", parts[0], err)
		}
		var environmentSubResult models.Environment
		if err = db.Where(&environmentSubQuery).Select("id").First(&environmentSubResult).Error; err != nil {
			return models.ChartRelease{}, fmt.Errorf("error querying for environment sub-selector %s: %w", parts[0], err)
		}
		query.EnvironmentID = &environmentSubResult.ID

		// chart
		chartSubQuery, err := chartModelFromSelector(parts[1])
		if err != nil {
			return models.ChartRelease{}, fmt.Errorf("error handling chart sub-selector %s: %w", parts[1], err)
		}
		var chartSubResult models.Chart
		if err = db.Where(&chartSubQuery).Select("id").First(&chartSubResult).Error; err != nil {
			return models.ChartRelease{}, fmt.Errorf("error querying for chart sub-selector %s: %w", parts[1], err)
		}
		query.ChartID = chartSubResult.ID

		return query, nil
	} else if strings.Count(selector, "/") == 2 { // cluster + namespace + chart
		parts := strings.Split(selector, "/")

		// cluster
		clusterSubQuery, err := clusterModelFromSelector(parts[0])
		if err != nil {
			return models.ChartRelease{}, fmt.Errorf("error handling cluster sub-selector %s: %w", parts[0], err)
		}
		var clusterSubResult models.Cluster
		if err = db.Where(&clusterSubQuery).Select("id").First(&clusterSubResult).Error; err != nil {
			return models.ChartRelease{}, fmt.Errorf("error querying for cluster sub-selector %s: %w", parts[0], err)
		}
		query.ClusterID = &clusterSubResult.ID

		// namespace
		namespace := parts[1]
		if !(utils.IsAlphaNumericWithHyphens(namespace) &&
			len(namespace) > 0 &&
			utils.IsStartingWithLetter(namespace) &&
			utils.IsEndingWithAlphaNumeric(namespace)) {
			return models.ChartRelease{}, fmt.Errorf("(%s) invalid chart release selector %s, namespace sub-selector %s was invalid", errors.BadRequest, selector, namespace)
		}
		query.Namespace = namespace

		// chart
		chartSubQuery, err := chartModelFromSelector(parts[2])
		if err != nil {
			return models.ChartRelease{}, fmt.Errorf("error handling chart sub-selector %s: %w", parts[2], err)
		}
		var chartSubResult models.Chart
		if err = db.Where(&chartSubQuery).Select("id").First(&chartSubResult).Error; err != nil {
			return models.ChartRelease{}, fmt.Errorf("error querying for chart sub-selector %s: %w", parts[2], err)
		}
		query.ChartID = chartSubResult.ID

		return query, nil
	} else if utils.IsAlphaNumericWithHyphens(selector) &&
		utils.IsStartingWithLetter(selector) &&
		utils.IsEndingWithAlphaNumeric(selector) { // name
		query.Name = selector
		return query, nil
	}
	return models.ChartRelease{}, fmt.Errorf("(%s) invalid chart release selector '%s'", errors.BadRequest, selector)
}
