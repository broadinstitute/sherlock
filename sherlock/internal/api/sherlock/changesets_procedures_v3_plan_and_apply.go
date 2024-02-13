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

// changesetsProceduresV3PlanAndApply godoc
//
//	@summary		Plan and apply version changes in one step
//	@description	Like calling the plan procedure immediately followed by the apply procedure. See those endpoints for more information.
//	@tags			Changesets
//	@accept			json
//	@produce		json
//	@param			changeset-plan-request	body		ChangesetV3PlanRequest	true	"Info on what version changes or refreshes to apply."
//	@param			verbose-output			query		bool					false	"If full information about the changesets should be returned, defaults to true. If false, only the IDs will be returned."
//	@success		200,201					{array}		ChangesetV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/changesets/procedures/v3/plan-and-apply [post]
func changesetsProceduresV3PlanAndApply(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body ChangesetV3PlanRequest
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) JSON error parsing body to %T: %w", errors.BadRequest, body, err))
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

	var chartReleaseChangesets, environmentChangesets, recreateChangesets []models.Changeset

	if chartReleaseChangesets, err = body.parseChartReleaseEntries(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error handling chart release entries: %w", err))
		return
	}
	if environmentChangesets, err = body.parseEnvironmentEntries(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error handling environment entries: %w", err))
		return
	}
	if recreateChangesets, err = body.parseRecreateEntries(db); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error handling recreate entries: %w", err))
		return
	}
	changesets := append(append(chartReleaseChangesets, environmentChangesets...), recreateChangesets...)

	var ret []models.Changeset
	var createdChangesetIDs []uint
	if createdChangesetIDs, err = models.PlanChangesets(db, changesets); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error planning changesets: %w", err))
		return
	} else if len(createdChangesetIDs) == 0 {
		ctx.JSON(http.StatusOK, []ChangesetV3{})
		return
	} else if err = models.ApplyChangesets(db, createdChangesetIDs); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error applying changesets: %w", err))
		return
	} else if verboseOutput {
		if err = db.Scopes(models.ReadChangesetScope).Find(&ret, createdChangesetIDs).Error; err != nil {
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
		ctx.JSON(http.StatusOK, utils.Map(createdChangesetIDs, func(id uint) ChangesetV3 {
			return ChangesetV3{CommonFields: CommonFields{ID: id}}
		}))
	}
}
