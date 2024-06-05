package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// clustersV3Get godoc
//
//	@summary		Get an individual Cluster
//	@description	Get an individual Cluster.
//	@tags			Clusters
//	@produce		json
//	@param			selector				path		string	true	"The selector of the Cluster, which can be either a numeric ID or the name."
//	@success		200						{object}	ClusterV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/clusters/v3/{selector} [get]
func clustersV3Get(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := clusterModelFromSelector(canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.Cluster
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, clusterFromModel(result))
}

func clusterModelFromSelector(selector string) (query models.Cluster, err error) {
	if len(selector) == 0 {
		return models.Cluster{}, fmt.Errorf("(%s) cluster selector cannot be empty", errors.BadRequest)
	}
	if utils.IsNumeric(selector) { // ID
		query.ID, err = utils.ParseUint(selector)
		return query, err
	} else if utils.IsAlphaNumericWithHyphens(selector) &&
		utils.IsStartingWithLetter(selector) &&
		utils.IsEndingWithAlphaNumeric(selector) { // Name
		query.Name = selector
		return query, nil
	}
	return models.Cluster{}, fmt.Errorf("(%s) invalid cluster selector '%s'", errors.BadRequest, selector)
}
