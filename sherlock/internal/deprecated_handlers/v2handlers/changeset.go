package v2handlers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	v2controllers2 "github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerChangesetHandlers(routerGroup *gin.RouterGroup, controller *v2controllers2.ChangesetController) {
	routerGroup.GET("/changesets", listChangeset(controller))
	routerGroup.GET("/changesets/*selector", getChangeset(controller))
	routerGroup.GET("/selectors/changesets/*selector", listChangesetSelectors(controller))
	routerGroup.POST("/procedures/changesets/plan-and-apply", planAndApplyChangeset(controller))
	routerGroup.POST("/procedures/changesets/plan", planChangeset(controller))
	routerGroup.POST("/procedures/changesets/apply", applyChangeset(controller))
	routerGroup.GET("/procedures/changesets/query-applied-for-chart-release/*selector", queryAppliedChangesetForChartRelease(controller))
	routerGroup.GET("/procedures/changesets/query-applied-for-version/:version-type/:chart/:version", queryAppliedChangesetForVersion(controller))
}

// listChangeset godoc
//
//	@summary		List Changeset entries
//	@description	List existing Changeset entries, ordered by most recently updated.
//	@tags			Changesets
//	@produce		json
//	@param			filter					query		v2controllers.Changeset	false	"Optional filters to apply to the returned entries"
//	@param			limit					query		int						false	"An optional limit to the number of entries returned"
//	@success		200						{array}		v2controllers.Changeset
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/changesets [get]
func listChangeset(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return handleList(&controller.ModelController)
}

// getChangeset godoc
//
//	@summary		Get a Changeset entry
//	@description	Get an existing Changeset entry via its "selector"--its numeric ID.
//	@tags			Changesets
//	@produce		json
//	@param			selector				path		string	true	"The Changeset to get's selector--numeric ID"
//	@success		200						{object}	v2controllers.Changeset
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/changesets/{selector} [get]
func getChangeset(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return handleGet(&controller.ModelController)
}

// listChangesetSelectors godoc
//
//	@summary		List Changeset selectors
//	@description	Validate a given Changeset selector and provide any other selectors that would match the same Changeset.
//	@tags			Changesets
//	@produce		json
//	@param			selector				path		string	true	"The selector of the Changeset to list other selectors for"
//	@success		200						{array}		string
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/selectors/changesets/{selector} [get]
func listChangesetSelectors(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return handleSelectorList(&controller.ModelController)
}

// planAndApplyChangeset godoc
//
//	@summary		Plan and apply version changes in one step
//	@description	Like the plan and apply endpoints immediately in sequence.
//	@tags			Changesets
//	@accept			json
//	@produce		json
//	@param			changeset-plan-request	body		v2controllers.ChangesetPlanRequest	true	"Info on what version changes or refreshes to perform"
//	@success		200,201					{array}		v2controllers.Changeset
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/procedures/changesets/plan-and-apply [post]
func planAndApplyChangeset(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := authentication.ShouldUseUser(ctx)
		if err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		var request v2controllers2.ChangesetPlanRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) JSON error parsing to %T: %v", errors.BadRequest, request, err))
			return
		}
		result, err := controller.PlanAndApply(request, user)
		if err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		if len(result) > 0 {
			ctx.JSON(http.StatusCreated, result)
		} else {
			ctx.JSON(http.StatusOK, result)
		}
	}
}

// planChangeset godoc
//
//	@summary		Plan--but do not apply--version changes to Chart Releases
//	@description	Refreshes and calculates version diffs for Chart Releases. If there's a diff, the plan is stored and returned so it can be applied later.
//	@description	Multiple Chart Releases can be specified--as can groups of Chart Releases from multiple Environments.
//	@tags			Changesets
//	@accept			json
//	@produce		json
//	@param			changeset-plan-request	body		v2controllers.ChangesetPlanRequest	true	"Info on what version changes or refreshes to plan"
//	@success		200,201					{array}		v2controllers.Changeset
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/procedures/changesets/plan [post]
func planChangeset(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := authentication.ShouldUseUser(ctx)
		if err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		var request v2controllers2.ChangesetPlanRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) JSON error parsing to %T: %v", errors.BadRequest, request, err))
			return
		}
		result, err := controller.Plan(request, user)
		if err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		if len(result) > 0 {
			ctx.JSON(http.StatusCreated, result)
		} else {
			ctx.JSON(http.StatusOK, result)
		}
	}
}

// applyChangeset godoc
//
//	@summary		Apply previously planned version changes to Chart Releases
//	@description	Looks up and applies previously-planned version diffs given by the ID. Other stored plans against the same Chart Releases are marked as superseded.
//	@description	Multiple Changesets can be specified simply by passing multiple IDs in the list.
//	@tags			Changesets
//	@accept			json
//	@produce		json
//	@param			apply-request			body		[]string	true	"String IDs of the Changesets to apply"
//	@success		200						{array}		v2controllers.Changeset
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/procedures/changesets/apply [post]
func applyChangeset(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := authentication.ShouldUseUser(ctx)
		if err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		var request []string
		if err := ctx.ShouldBindJSON(&request); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) JSON error parsing to %T: %v", errors.BadRequest, request, err))
			return
		}
		result, err := controller.Apply(request, user)
		if err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

// queryAppliedChangesetForChartRelease godoc
//
//	@summary		List applied Changesets for a Chart Release
//	@description	List existing applied Changesets for a particular Chart Release, ordered by most recently applied.
//	@tags			Changesets
//	@produce		json
//	@param			selector				path		string	true	"Selector the Chart Release to find applied Changesets for"
//	@param			offset					query		int		false	"An optional offset to skip a number of latest Changesets"
//	@param			limit					query		int		false	"An optional limit to the number of entries returned"
//	@success		200						{array}		v2controllers.Changeset
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/procedures/changesets/query-applied-for-chart-release/{selector} [get]
func queryAppliedChangesetForChartRelease(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		offsetString := ctx.DefaultQuery("offset", "0")
		offset, err := strconv.Atoi(offsetString)
		if err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) error parsing offset parameter: %v", errors.BadRequest, err))
			return
		}
		limitString := ctx.DefaultQuery("limit", "0")
		limit, err := strconv.Atoi(limitString)
		if err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) error parsing limit parameter: %v", errors.BadRequest, err))
			return
		}
		result, err := controller.QueryAppliedForChartRelease(formatSelector(ctx.Param("selector")), offset, limit)
		if err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

// queryAppliedChangesetForVersion godoc
//
//	@summary		List applied Changesets for an App or Chart Version
//	@description	List existing applied Changesets that newly deployed a given App Version or Chart Version, ordered by most recently applied.
//	@tags			Changesets
//	@produce		json
//	@param			version-type			path		string	true	"The type of the version, either 'app' or 'chart'"	Enums(app, chart)
//	@param			chart					path		string	true	"The chart the version belongs to"
//	@param			version					path		string	true	"The version to look for"
//	@success		200						{array}		v2controllers.Changeset
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/procedures/changesets/query-applied-for-version/{version-type}/{chart}/{version} [get]
func queryAppliedChangesetForVersion(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if result, err := controller.QueryAppliedForVersion(
			ctx.Param("chart"),
			ctx.Param("version"),
			ctx.Param("version-type"),
		); err != nil {
			errors.AbortRequest(ctx, err)
		} else {
			ctx.JSON(http.StatusOK, result)
		}
	}
}
