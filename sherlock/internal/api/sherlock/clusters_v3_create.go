package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// clustersV3Create godoc
//
//	@summary		Create a Cluster
//	@description	Create a Cluster.
//	@tags			Clusters
//	@accept			json
//	@produce		json
//	@param			cluster					body		ClusterV3Create	true	"The Cluster to create"
//	@success		201						{object}	ClusterV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/clusters/v3 [post]
func clustersV3Create(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body ClusterV3Create
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	toCreate := body.toModel()
	if err = db.Create(&toCreate).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	var result models.Cluster
	if err = db.Preload(clause.Associations).First(&result, toCreate.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, clusterFromModel(result))
}
