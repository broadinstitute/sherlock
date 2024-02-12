package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

// ciIdentifiersV3Get godoc
//
//	@summary		Get CiRuns for a resource by its CiIdentifier
//	@description	Get CiRuns for a resource by its CiIdentifier, which can be referenced by '{type}/{selector...}'.
//	@tags			CiIdentifiers
//	@produce		json
//	@param			selector				path		string	true	"The selector of CiIdentifier, which can be referenced either by numeric ID or indirectly by '{type}/{selector...}'"
//	@param			limitCiRuns				query		int		false	"Control how many CiRuns are returned (default 10)"
//	@param			offsetCiRuns			query		int		false	"Control the offset for the returned CiRuns (default 0)"
//	@param			allowStubCiRuns			query		bool	false	"Allow stub CiRuns potentially lacking fields like status or startedAt to be returned (default false)"
//	@success		200						{object}	CiIdentifierV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/ci-identifiers/v3/{selector} [get]
func ciIdentifiersV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := ciIdentifierModelFromSelector(db, canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	limitCiRuns, err := utils.ParseInt(ctx.DefaultQuery("limitCiRuns", "10"))
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}
	offsetCiRuns, err := utils.ParseInt(ctx.DefaultQuery("offsetCiRuns", "0"))
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}
	allowStubCiRuns := ctx.DefaultQuery("allowStubCiRuns", "false") == "true"
	var result models.CiIdentifier
	if err = db.Preload("CiRuns", func(tx *gorm.DB) *gorm.DB {
		if allowStubCiRuns {
			return tx.Limit(limitCiRuns).Offset(offsetCiRuns).Order("created_at desc")
		} else {
			return tx.Where("status != '' AND status IS NOT NULL AND started_at IS NOT NULL").Limit(limitCiRuns).Offset(offsetCiRuns).Order("started_at desc")
		}
	}).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if err = result.FillCiRunResourceStatuses(db); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, ciIdentifierFromModel(result))
}

func ciIdentifierModelFromSelector(db *gorm.DB, selector string) (query models.CiIdentifier, err error) {
	if len(selector) == 0 {
		return models.CiIdentifier{}, fmt.Errorf("(%s) selector cannot be empty", errors.BadRequest)
	} else if utils.IsNumeric(selector) {
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else if strings.Count(selector, "/") > 0 {
		// resource type + type's selector ...
		parts := strings.Split(selector, "/")
		query.ResourceType = parts[0]
		resourceSelector := strings.Join(parts[1:], "/")
		switch query.ResourceType {
		case "chart":
			chartQuery, err := chartModelFromSelector(resourceSelector)
			if err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error parsing chart selector '%s': %w", resourceSelector, err)
			}
			var result models.Chart
			if err = db.Where(&chartQuery).Select("id").First(&result).Error; err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error fetching chart '%s': %w", resourceSelector, err)
			}
			query.ResourceID = result.ID
		case "chart-version":
			chartVersionQuery, err := chartVersionModelFromSelector(db, resourceSelector)
			if err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error parsing chart version selector '%s': %w", resourceSelector, err)
			}
			var result models.ChartVersion
			if err = db.Where(&chartVersionQuery).Select("id").First(&result).Error; err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error fetching chart version '%s': %w", resourceSelector, err)
			}
			query.ResourceID = result.ID
		case "app-version":
			appVersionQuery, err := appVersionModelFromSelector(db, resourceSelector)
			if err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error parsing app version selector '%s': %w", resourceSelector, err)
			}
			var result models.AppVersion
			if err = db.Where(&appVersionQuery).Select("id").First(&result).Error; err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error fetching app version '%s': %w", resourceSelector, err)
			}
			query.ResourceID = result.ID
		case "cluster":
			clusterQuery, err := clusterModelFromSelector(resourceSelector)
			if err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error parsing cluster selector '%s': %w", resourceSelector, err)
			}
			var result models.Cluster
			if err = db.Where(&clusterQuery).Select("id").First(&result).Error; err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error fetching cluster '%s': %w", resourceSelector, err)
			}
			query.ResourceID = result.ID
		case "environment":
			environmentQuery, err := environmentModelFromSelector(resourceSelector)
			if err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error parsing environment selector '%s': %w", resourceSelector, err)
			}
			var result models.Environment
			if err = db.Where(&environmentQuery).Select("id").First(&result).Error; err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error fetching environment '%s': %w", resourceSelector, err)
			}
			query.ResourceID = result.ID
		case "chart-release":
			chartReleaseQuery, err := chartReleaseModelFromSelector(db, resourceSelector)
			if err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error parsing chart release selector '%s': %w", resourceSelector, err)
			}
			var result models.ChartRelease
			if err = db.Where(&chartReleaseQuery).Select("id").First(&result).Error; err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error fetching chart release '%s': %w", resourceSelector, err)
			}
			query.ResourceID = result.ID
		case "changeset":
			// To match the behavior of the other resource types, we actually do check that the changeset ID exists
			queryID, err := utils.ParseUint(resourceSelector)
			if err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error parsing changeset ID '%s': %w", resourceSelector, err)
			}
			var result models.Changeset
			if err = db.Select("id").First(&result, queryID).Error; err != nil {
				return models.CiIdentifier{}, fmt.Errorf("error fetching changeset ID '%s': %w", resourceSelector, err)
			}
			query.ResourceID = result.ID
		default:
			err = fmt.Errorf("(%s) invalid CI identifier selector '%s', resource type sub-selector '%s' wasn't recognized", errors.BadRequest, selector, query.ResourceType)
		}
		return query, err
	} else {
		return models.CiIdentifier{}, fmt.Errorf("(%s) invalid CI identifier selector '%s'", errors.BadRequest, selector)
	}
}
