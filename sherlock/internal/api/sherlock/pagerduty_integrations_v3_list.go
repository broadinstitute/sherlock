package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// pagerdutyIntegrationsV3List godoc
//
//	@summary		List PagerdutyIntegrations matching a filter
//	@description	List PagerdutyIntegrations matching a filter.
//	@tags			PagerdutyIntegrations
//	@produce		json
//	@param			filter					query		PagerdutyIntegrationV3	false	"Filter the returned PagerdutyIntegrations"
//	@param			limit					query		int						false	"Control how many PagerdutyIntegrations are returned (default 0, meaning all)"
//	@param			offset					query		int						false	"Control the offset for the returned PagerdutyIntegrations (default 0)"
//	@success		200						{array}		PagerdutyIntegrationV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/pagerduty-integrations/v3 [get]
func pagerdutyIntegrationsV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter PagerdutyIntegrationV3
	if err = ctx.ShouldBindQuery(&filter); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	modelFilter := filter.toModel()

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
	var results []models.PagerdutyIntegration
	chain := db.
		Where(&modelFilter)
	if limit > 0 {
		chain = chain.Limit(limit)
	}
	if err = chain.
		Offset(offset).
		Order("pagerduty_id asc").
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, pagerdutyIntegrationFromModel))
}
