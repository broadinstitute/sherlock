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

// ciIdentifiersV3Get godoc
//
//	@summary		List CiIdentifiers matching a filter
//	@description	List CiIdentifiers matching a filter. The CiRuns would have to re-queried directly to load the CiRuns.
//	@description	This is mainly helpful for debugging and directly querying the existence of a CiIdentifier. Results are
//	@description	ordered by creation date, starting at most recent.
//	@tags			CiIdentifiers
//	@produce		json
//	@param			filter					query		CiIdentifierV3	false	"Filter the returned CiIdentifiers"
//	@param			limit					query		int				false	"Control how many CiIdentifiers are returned (default 100)"
//	@param			offset					query		int				false	"Control the offset for the returned CiIdentifiers (default 0)"
//	@success		200						{array}		CiIdentifierV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/ci-identifiers/v3 [get]
func ciIdentifiersV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter CiIdentifierV3
	if err = ctx.ShouldBindQuery(&filter); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	modelFilter := filter.toModel()
	limit, err := utils.ParseInt(ctx.DefaultQuery("limit", "100"))
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}
	offset, err := utils.ParseInt(ctx.DefaultQuery("offset", "0"))
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) %v", errors.BadRequest, err))
		return
	}
	var results []models.CiIdentifier
	if err = db.Where(&modelFilter).Limit(limit).Offset(offset).Order("created_at desc").Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, ciIdentifierFromModel))
}
