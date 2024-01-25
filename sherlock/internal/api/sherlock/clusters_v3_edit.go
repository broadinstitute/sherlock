package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// clustersV3Edit godoc
//
//	@summary		Edit an individual Cluster
//	@description	Edit an individual Cluster.
//	@tags			Clusters
//	@produce		json
//	@param			selector				path		string			true	"The selector of the Cluster, which can be either a numeric ID or the name."
//	@param			cluster					body		ClusterV3Edit	true	"The edits to make to the Cluster"
//	@success		200						{object}	ClusterV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/clusters/v3/{selector} [patch]
func clustersV3Edit(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := clusterModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var body ClusterV3Edit
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	edits := body.toModel()

	var toEdit models.Cluster
	if err = db.Preload(clause.Associations).Where(&query).First(&toEdit).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if err = db.Model(&toEdit).Omit(clause.Associations).Updates(&edits).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, clusterFromModel(toEdit))
}
