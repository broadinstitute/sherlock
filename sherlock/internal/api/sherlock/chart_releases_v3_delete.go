package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// chartReleasesV3Delete godoc
//
//	@summary		Delete an individual ChartRelease
//	@description	Delete an individual ChartRelease by its ID.
//	@tags			ChartReleases
//	@produce		json
//	@param			selector				path		string	true	"The selector of the ChartRelease, which can be either a numeric ID, the name, environment + '/' + chart, or cluster + '/' + namespace + '/' + chart."
//	@success		200						{object}	ChartReleaseV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/chart-releases/v3/{selector} [delete]
func chartReleasesV3Delete(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	query, err := chartReleaseModelFromSelector(db, canonicalizeSelector(ctx.Param("selector")))
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var result models.ChartRelease
	if err = db.Preload(clause.Associations).Where(&query).First(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	if err = db.Omit(clause.Associations).Delete(&result).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, chartReleaseFromModel(result))
}
