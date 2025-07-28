package sherlock

import (
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
)

// serviceAlertsV3List godoc
//
//	@summary		List ServiceAlerts matching a filter
//	@description	List ServiceAlerts matching a filter.
//	@tags			ServiceAlert
//	@produce		json
//	@param			filter					query		ServiceAlertV3	false	"Filter the returned Service Alerts"
//	@param			limit					query		int				false	"Control how many Service Alerts are returned (default 0, no limit)"
//	@param			offset					query		int				false	"Control the offset for the returned Service Alerts (default 0)"
//	@param			include-deleted			query		bool		false	"Control if only active Service Alerts are returned, set to true to return deleted Alerts (default false)"
//	@success		200						{array}		ServiceAlertV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/service-alerts/v3 [get]
func serviceAlertsV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter ServiceAlertV3
	if err = ctx.ShouldBindQuery(&filter); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	modelFilter, err := filter.toModel(db)

	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}

	limit, err := utils.ParseInt(ctx.DefaultQuery("limit", "0"))
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}
	offset, err := utils.ParseInt(ctx.DefaultQuery("offset", "0"))
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}

	var results []models.ServiceAlert
	chain := db.
		Where(&modelFilter)

	// includes soft deleted items if set to true
	if includeDeleted := ctx.DefaultQuery("include-deleted", "false"); includeDeleted == "true" {
		chain = chain.Unscoped()
	}

	if limit > 0 {
		chain = chain.Limit(limit)
	}
	// change below to order by time or something? Need to see what's available
	if err = chain.
		Offset(offset).
		Order("created_at asc").
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, utils.Map(results, ServiceAlertFromModel))
}
