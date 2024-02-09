package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// incidentsV3List godoc
//
//	@summary		List Incidents matching a filter
//	@description	List Incidents matching a filter.
//	@tags			Incidents
//	@produce		json
//	@param			filter					query		IncidentV3	false	"Filter the returned Incidents"
//	@param			limit					query		int			false	"Control how many Incidents are returned (default 0, meaning all)"
//	@param			offset					query		int			false	"Control the offset for the returned Incidents (default 0)"
//	@success		200						{array}		IncidentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/incidents/v3 [get]
func incidentsV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter IncidentV3
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
	var results []models.Incident
	chain := db.
		Where(&modelFilter)
	if limit > 0 {
		chain = chain.Limit(limit)
	}
	if err = chain.
		Offset(offset).
		Order("started_at desc").
		Preload(clause.Associations).
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, incidentFromModel))
}
