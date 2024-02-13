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

func listChangeset(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return handleList(&controller.ModelController)
}

func getChangeset(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return handleGet(&controller.ModelController)
}

func listChangesetSelectors(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return handleSelectorList(&controller.ModelController)
}

func planAndApplyChangeset(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := authentication.ShouldUseUser(ctx)
		if err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		var request v2controllers2.ChangesetPlanRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) JSON error parsing to %T: %w", errors.BadRequest, request, err))
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

func planChangeset(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := authentication.ShouldUseUser(ctx)
		if err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		var request v2controllers2.ChangesetPlanRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) JSON error parsing to %T: %w", errors.BadRequest, request, err))
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

func applyChangeset(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user, err := authentication.ShouldUseUser(ctx)
		if err != nil {
			errors.AbortRequest(ctx, err)
			return
		}
		var request []string
		if err := ctx.ShouldBindJSON(&request); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) JSON error parsing to %T: %w", errors.BadRequest, request, err))
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

func queryAppliedChangesetForChartRelease(controller *v2controllers2.ChangesetController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		offsetString := ctx.DefaultQuery("offset", "0")
		offset, err := strconv.Atoi(offsetString)
		if err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) error parsing offset parameter: %w", errors.BadRequest, err))
			return
		}
		limitString := ctx.DefaultQuery("limit", "0")
		limit, err := strconv.Atoi(limitString)
		if err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) error parsing limit parameter: %w", errors.BadRequest, err))
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
