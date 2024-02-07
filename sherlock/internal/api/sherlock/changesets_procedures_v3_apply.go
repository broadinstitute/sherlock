package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"strings"
)

// changesetsProceduresV3Apply godoc
//
//	@summary		Apply previously planned version changes to Chart Releases
//	@description	Looks up and applies previously-planned version diffs given by the ID. Other stored plans against the same Chart Releases are marked as superseded.
//	@description	Multiple Changesets can be specified simply by passing multiple IDs in the list.
//	@tags			Changesets
//	@accept			json
//	@produce		json
//	@param			apply-request			body		[]string	true	"String IDs of the Changesets to apply"
//	@param			verbose-output			query		bool		false	"If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned."
//	@success		200						{array}		ChangesetV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/changesets/procedures/v3/apply [post]
func changesetsProceduresV3Apply(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body []string
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) JSON error parsing body to %T: %w", errors.BadRequest, body, err))
		return
	}
	changesetIDs := make([]uint, len(body))
	for idx, idString := range body {
		if changesetIDs[idx], err = utils.ParseUint(idString); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) couldn't parse '%s' to an ID: %v", errors.BadRequest, idString, err))
			return
		} else if changesetIDs[idx] == 0 {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) invalid ID: %d", errors.BadRequest, changesetIDs[idx]))
			return
		}
	}
	if len(changesetIDs) == 0 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) no changesets specified", errors.BadRequest))
		return
	}

	var verboseOutput bool
	if verboseOutputString := ctx.DefaultQuery("verbose-output", "true"); verboseOutputString == "true" {
		verboseOutput = true
	} else if verboseOutputString == "false" {
		verboseOutput = false
	} else {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) couldn't parse 'verbose-output' to a boolean: %s", errors.BadRequest, verboseOutputString))
		return
	}

	var ret []models.Changeset
	if err = models.ApplyChangesets(db, changesetIDs); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error applying changesets: %w", err))
		return
	} else if verboseOutput {
		if err = db.Scopes(models.ReadChangesetScope).Find(&ret, changesetIDs).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error querying applied changesets: %w", err))
			return
		} else {
			slices.SortFunc(ret, func(a, b models.Changeset) int {
				if a.ChartRelease != nil || b.ChartRelease != nil {
					return strings.Compare(a.ChartRelease.Name, b.ChartRelease.Name)
				} else {
					return 0
				}
			})
			ctx.JSON(http.StatusOK, utils.Map(ret, changesetFromModel))
		}
	} else {
		ctx.JSON(http.StatusOK, utils.Map(changesetIDs, func(id uint) ChangesetV3 {
			return ChangesetV3{CommonFields: CommonFields{ID: id}}
		}))
	}
}
