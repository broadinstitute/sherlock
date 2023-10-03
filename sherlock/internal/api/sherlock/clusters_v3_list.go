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

// clustersV3List godoc
//
//	@summary		List Clusters matching a filter
//	@description	List Clusters matching a filter.
//	@tags			Clusters
//	@produce		json
//	@param			filter					query		ClusterV3	false	"Filter the returned Clusters"
//	@param			limit					query		int			false	"Control how many Clusters are returned (default 0, meaning all)"
//	@param			offset					query		int			false	"Control the offset for the returned Clusters (default 0)"
//	@success		200						{array}		ClusterV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/clusters/v3 [get]
func clustersV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter ClusterV3
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
	var results []models.Cluster
	chain := db.
		Where(&modelFilter)
	if limit > 0 {
		chain = chain.Limit(limit)
	}
	if err = chain.
		Offset(offset).
		Order("name asc").
		Find(&results).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(results, clusterFromModel))
}
