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

// changesetsProceduresV3ChartReleaseHistory godoc
//
//	@summary		List applied Changesets for a Chart Release
//	@description	List existing applied Changesets for a particular Chart Release, ordered by most recently applied.
//	@tags			Changesets
//	@produce		json
//	@param			selector				path		string	true	"Selector of the Chart Release to find applied Changesets for"
//	@param			offset					query		int		false	"An optional offset to skip a number of latest Changesets"
//	@param			limit					query		int		false	"An optional limit to the number of entries returned"
//	@success		200						{array}		ChangesetV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/changesets/procedures/v3/chart-release-history/{chart-release} [get]
func changesetsProceduresV3ChartReleaseHistory(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

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

	// get chart release ID
	var chartReleaseQuery models.ChartRelease
	if chartReleaseQuery, err = chartReleaseModelFromSelector(db, canonicalizeSelector(ctx.Param("chart-release"))); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}
	var chartRelease models.ChartRelease
	if err = db.Select("id", "name").Where(&chartReleaseQuery).First(&chartRelease).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error querying chart release '%s': %w", ctx.Param("chart-release"), err))
		return
	}

	// Could we get the changesets all at once... yes. But we get better error output by finding the path first and
	// then loading the actual data to return to the user. With how Gorm does preloading, I wouldn't be surprised if doing
	// that was faster than adding the scope to this chain. We could probably still do it all at once with a sub-query,
	// but this is a seldom-used endpoint and we'd prefer it to just obviously be correct.
	var changesetIDs []uint
	chain := db.
		Model(&models.Changeset{}).
		Unscoped().
		Where(&models.Changeset{ChartReleaseID: chartRelease.ID}).
		Where("applied_at is not null").
		Order("applied_at desc").
		Offset(offset)
	if limit > 0 {
		chain = chain.Limit(limit)
	}
	if err = chain.Pluck("id", &changesetIDs).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error querying changesets for chart release '%s': %w", chartRelease.Name, err))
		return
	}

	var ret []models.Changeset
	if err = db.Scopes(models.ReadChangesetScope).Order("applied_at desc").Find(&ret, changesetIDs).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error querying changesets to return: %w", err))
		return
	}
	ctx.JSON(http.StatusOK, utils.Map(ret, changesetFromModel))
}
