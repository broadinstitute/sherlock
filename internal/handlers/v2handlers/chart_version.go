package v2handlers

import (
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterChartVersionHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ChartVersionController) {
	routerGroup.POST("/chart-versions", createChartVersion(controller))
	routerGroup.GET("/chart-versions", listChartVersion(controller))
	routerGroup.GET("/chart-versions/*selector", getChartVersion(controller))
	routerGroup.GET("/selectors/chart-versions/*selector", listChartVersionSelectors(controller))
}

// createChartVersion godoc
// @summary      Create a new ChartVersion entry
// @description  Create a new ChartVersion entry. Note that fields are immutable after creation.
// @tags         ChartVersions
// @accept       json
// @produce      json
// @param        chart-version            body      v2controllers.CreatableChartVersion  true  "The ChartVersion to create"
// @success      200                      {object}  v2controllers.ChartVersion
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/chart-versions [post]
func createChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listChartVersion godoc
// @summary      List ChartVersion entries
// @description  List existing ChartVersion entries, ordered by most recently updated.
// @tags         ChartVersions
// @produce      json
// @param        filter                   query     v2controllers.ChartVersion  false  "Optional filters to apply to the returned entries"
// @param        limit                    query     int                         false  "An optional limit to the number of entries returned"
// @success      200                      {array}   v2controllers.ChartVersion
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/chart-versions [get]
func listChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getChartVersion godoc
// @summary      Get a ChartVersion entry
// @description  Get an existing ChartVersion entry via one its "selector"--its numeric ID.
// @tags         ChartVersions
// @produce      json
// @param        selector                 path      string  true  "The ChartVersion to get's selector: name or numeric ID"
// @success      200                      {object}  v2controllers.ChartVersion
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/chart-versions/{selector} [get]
func getChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// listChartVersionSelectors godoc
// @summary      List ChartVersion selectors
// @description  Validate a given ChartVersion selector and provide any other selectors that would match the same ChartVersion.
// @tags         ChartVersions
// @produce      json
// @param        selector                 path      string  true  "The selector of the ChartVersion to list other selectors for"
// @success      200                      {array}   string
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/selectors/chart-versions/{selector} [get]
func listChartVersionSelectors(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
