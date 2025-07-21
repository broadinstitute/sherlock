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

// databaseInstancesV3List godoc
//
//	@summary		List DatabaseInstances matching a filter
//	@description	List DatabaseInstances matching a filter.
//	@tags			DatabaseInstances
//	@produce		json
//	@param			filter					query		DatabaseInstanceV3	false	"Filter the returned DatabaseInstances"
//	@param			limit					query		int					false	"Control how many DatabaseInstances are returned (default 0, meaning all)"
//	@param			offset					query		int					false	"Control the offset for the returned DatabaseInstances (default 0)"
//	@success		200						{array}		DatabaseInstanceV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/database-instances/v3 [get]
func databaseInstancesV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter DatabaseInstanceV3
	if err = ctx.ShouldBindQuery(&filter); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	modelFilter, err := filter.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
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
	var results []models.DatabaseInstance
	chain := db.
		Where(&modelFilter)
	if limit > 0 {
		chain = chain.Limit(limit)
	}
	if err = chain.
		Offset(offset).
		Order("created_at desc").
		Preload(clause.Associations).
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, databaseInstanceFromModel))
}
