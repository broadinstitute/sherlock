package v2handlers

import (
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterChartReleaseHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ChartReleaseController) {
	routerGroup.POST("/chart-releases", createChartRelease(controller))
	routerGroup.GET("/chart-releases", listChartRelease(controller))
	routerGroup.GET("/chart-releases/*selector", getChartRelease(controller))
	routerGroup.PATCH("/chart-releases/*selector", editChartRelease(controller))
	routerGroup.DELETE("/chart-releases/*selector", deleteChartRelease(controller))
	routerGroup.GET("/selectors/chart-releases/*selector", listChartReleaseSelectors(controller))
}

// createChartRelease godoc
// @summary     Create a new ChartRelease entry
// @description Create a new ChartRelease entry. Note that some fields are immutable after creation; /edit lists mutable fields.
// @tags        ChartReleases
// @accept      json
// @produce     json
// @param       chart-release           body     v2controllers.CreatableChartRelease true "The ChartRelease to create"
// @success     201                     {object} v2controllers.ChartRelease
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/chart-releases [post]
func createChartRelease(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listChartRelease godoc
// @summary     List ChartRelease entries
// @description List existing ChartRelease entries, ordered by most recently updated.
// @tags        ChartReleases
// @produce     json
// @param       filter                  query    v2controllers.ChartRelease false "Optional filters to apply to the returned entries"
// @param       limit                   query    int                        false "An optional limit to the number of entries returned"
// @success     200                     {array}  v2controllers.ChartRelease
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/chart-releases [get]
func listChartRelease(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getChartRelease godoc
// @summary     Get a ChartRelease entry
// @description Get an existing ChartRelease entry via one of its "selectors": name, numeric ID, environment/chart, or cluster/namespace/chart.
// @tags        ChartReleases
// @produce     json
// @param       selector                path     string true "The ChartRelease to get's selector: name or numeric ID"
// @success     200                     {object} v2controllers.ChartRelease
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/chart-releases/{selector} [get]
func getChartRelease(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// editChartRelease godoc
// @summary     Edit a ChartRelease entry
// @description Edit an existing ChartRelease entry via one of its "selectors": name, numeric ID, environment/chart, or cluster/namespace/chart. Note that only mutable fields are available here, immutable fields can only be set using /create.
// @tags        ChartReleases
// @accept      json
// @produce     json
// @param       selector                path     string                             true "The ChartRelease to edit's selector: name or numeric ID"
// @param       chart-release           body     v2controllers.EditableChartRelease true "The edits to make to the ChartRelease"
// @success     200                     {object} v2controllers.ChartRelease
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/chart-releases/{selector} [patch]
func editChartRelease(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

// deleteChartRelease godoc
// @summary     Delete a ChartRelease entry
// @description Delete an existing ChartRelease entry via one of its "selectors": name, numeric ID, environment/chart, or cluster/namespace/chart.
// @tags        ChartReleases
// @produce     json
// @param       selector                path     string true "The ChartRelease to delete's selector: name or numeric ID"
// @success     200                     {object} v2controllers.ChartRelease
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/chart-releases/{selector} [delete]
func deleteChartRelease(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

// listChartReleaseSelectors godoc
// @summary     List ChartRelease selectors
// @description Validate a given ChartRelease selector and provide any other selectors that would match the same ChartRelease.
// @tags        ChartReleases
// @produce     json
// @param       selector                path     string true "The selector of the ChartRelease to list other selectors for"
// @success     200                     {array}  string
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/selectors/chart-releases/{selector} [get]
func listChartReleaseSelectors(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
