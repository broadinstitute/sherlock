package v2handlers

import (
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterChartDeployRecordHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ChartDeployRecordController) {
	routerGroup.POST("/chart-deploy-records", createChartDeployRecord(controller))
	routerGroup.GET("/chart-deploy-records", listChartDeployRecord(controller))
	routerGroup.GET("/chart-deploy-records/*selector", getChartDeployRecord(controller))
	routerGroup.GET("/selectors/chart-deploy-records/*selector", listChartDeployRecordSelectors(controller))
}

// createChartDeployRecord godoc
// @summary     Create a new ChartDeployRecord entry
// @description Create a new ChartDeployRecord entry. Note that fields are immutable after creation.
// @tags        ChartDeployRecords
// @accept      json
// @produce     json
// @param       chart-deploy-record     body     v2controllers.CreatableChartDeployRecord true "The ChartDeployRecord to create"
// @success     200,201                 {object} v2controllers.ChartDeployRecord
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/chart-deploy-records [post]
func createChartDeployRecord(controller *v2controllers.ChartDeployRecordController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listChartDeployRecord godoc
// @summary     List ChartDeployRecord entries
// @description List existing ChartDeployRecord entries, ordered by most recently updated.
// @tags        ChartDeployRecords
// @produce     json
// @param       filter                  query    v2controllers.ChartDeployRecord false "Optional filters to apply to the returned entries"
// @param       limit                   query    int                             false "An optional limit to the number of entries returned"
// @success     200                     {array}  v2controllers.ChartDeployRecord
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/chart-deploy-records [get]
func listChartDeployRecord(controller *v2controllers.ChartDeployRecordController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getChartDeployRecord godoc
// @summary     Get a ChartDeployRecord entry
// @description Get an existing ChartDeployRecord entry via one its "selector"--its numeric ID.
// @tags        ChartDeployRecords
// @produce     json
// @param       selector                path     string true "The ChartDeployRecord to get's selector: name or numeric ID"
// @success     200                     {object} v2controllers.ChartDeployRecord
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/chart-deploy-records/{selector} [get]
func getChartDeployRecord(controller *v2controllers.ChartDeployRecordController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// listChartDeployRecordSelectors godoc
// @summary     List ChartDeployRecord selectors
// @description Validate a given ChartDeployRecord selector and provide any other selectors that would match the same ChartDeployRecord.
// @tags        ChartDeployRecords
// @produce     json
// @param       selector                path     string true "The selector of the ChartDeployRecord to list other selectors for"
// @success     200                     {array}  string
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/selectors/chart-deploy-records/{selector} [get]
func listChartDeployRecordSelectors(controller *v2controllers.ChartDeployRecordController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
