package v2handlers

import (
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterChartVersionHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ChartVersionController) {
	routerGroup.POST("/create", createChartVersion(controller))
	routerGroup.GET("/get/*selector", getChartVersion(controller))
	routerGroup.GET("/selectors/*selector", listChartVersionSelectors(controller))
	routerGroup.GET("/list", listChartVersion(controller))
	routerGroup.POST("/list", listChartVersionWithFilter(controller))
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
// @router       /api/v2/chart-versions/create [post]
func createChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// getChartVersion godoc
// @summary      Get a ChartVersion entry
// @description  Get an existing ChartVersion entry via one its "selector"--its numeric ID.
// @tags         ChartVersions
// @produce      json
// @param        selector                 path      string  true  "The ChartVersion to get's selector: name or numeric ID"
// @success      200                      {object}  v2controllers.ChartVersion
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/chart-versions/get/{selector} [get]
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
// @router       /api/v2/chart-versions/selectors/{selector} [get]
func listChartVersionSelectors(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}

// listChartVersion godoc
// @summary      List ChartVersion entries
// @description  List existing ChartVersion entries, ordered by most recently updated.
// @tags         ChartVersions
// @produce      json
// @param        limit                    query     int  false  "An optional limit to the number of entries returned"
// @success      200                      {array}   v2controllers.ChartVersion
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/chart-versions/list [get]
func listChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleList(controller)
}

// listChartVersionWithFilter godoc
// @summary      List ChartVersion entries with field filters
// @description  List existing ChartVersion entries, ordered by most recently updated. Entries will be filtered to only return ones matching the provided non-empty fields in the body.
// @tags         ChartVersions
// @accept       json
// @produce      json
// @param        limit                    query     int                         false  "An optional limit to the number of entries returned"
// @param        chart-version            body      v2controllers.ChartVersion  true   "The fields and values to filter on (omit a field to not filter based on it)"
// @success      200                      {array}   v2controllers.ChartVersion
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/chart-versions/list [post]
func listChartVersionWithFilter(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleListWithFilter(controller)
}
