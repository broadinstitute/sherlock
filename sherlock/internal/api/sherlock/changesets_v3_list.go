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

// changesetsV3List godoc
//
//	@summary		List Changesets matching a filter
//	@description	List Changesets matching a filter.
//	@tags			Changesets
//	@produce		json
//	@param			filter					query		ChangesetV3	false	"Filter the returned Changesets"
//	@param			id						query		[]int		false	"Get specific changesets by their IDs, can be passed multiple times"
//	@param			limit					query		int			false	"Control how many Changesets are returned (default 100), ignored if specific IDs are passed"
//	@param			offset					query		int			false	"Control the offset for the returned Changesets (default 0), ignored if specific IDs are passed"
//	@success		200						{array}		ChangesetV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/changesets/v3 [get]
func changesetsV3List(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}
	var filter ChangesetV3
	if err = ctx.ShouldBindQuery(&filter); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	modelFilter, err := filter.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) unable to process filter: %v", errors.BadRequest, err))
		return
	}

	var results []models.Changeset

	idStrings := ctx.QueryArray("id")
	if len(idStrings) > 0 {
		ids := make([]uint, len(idStrings))
		for i, idString := range idStrings {
			if ids[i], err = utils.ParseUint(idString); err != nil {
				errors.AbortRequest(ctx, fmt.Errorf("(%s) couldn't parse '%s' to an ID: %v", errors.BadRequest, idString, err))
				return
			}
		}

		// We wipe out the ID field of the modelFilter. Either we have one ID and it being in the filter is superfluous,
		// or we have multiple and it being in the filter is actually a problem (because the filter would have just one ID)
		modelFilter.ID = 0

		if err = db.
			// In most cases, this modelFilter will be empty, but we pass it anyway in case someone does something like
			// id=1&id=2&chartRelease=leonardo-dev, where they're trying to filter a known set of IDs down.
			Where(&modelFilter).
			Scopes(models.ReadChangesetScope).
			Order("created_at desc").
			Find(&results, ids).Error; err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
	} else {
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

		if err = db.
			Where(&modelFilter).
			Scopes(models.ReadChangesetScope).
			Limit(limit).
			Offset(offset).
			Order("created_at desc").
			Find(&results).Error; err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
	}

	ctx.JSON(http.StatusOK, utils.Map(results, changesetFromModel))
}
