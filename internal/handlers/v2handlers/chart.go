package v2handlers

import (
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterChartHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ChartController) {
	routerGroup.POST("/create", createChart(controller))
	routerGroup.GET("/get/*selector", getChart(controller))
	routerGroup.PATCH("/edit/*selector", editChart(controller))
	routerGroup.DELETE("/delete/*selector", deleteChart(controller))
	routerGroup.GET("/selectors/*selector", listChartSelectors(controller))
	routerGroup.GET("/list", listChart(controller))
	routerGroup.POST("/list", listChartWithFilter(controller))
}

// createChart godoc
// @summary      Create a new Chart entry
// @description  Create a new Chart entry. Note that some fields are immutable after creation; /edit lists mutable fields.
// @tags         Charts
// @accept       json
// @produce      json
// @param        chart                    body      v2controllers.CreatableChart  true  "The Chart to create"
// @success      200                      {object}  v2controllers.Chart
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/charts/create [post]
func createChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleCreate(&controller.ImmutableModelController)
}

// getChart godoc
// @summary      Get a Chart entry
// @description  Get an existing Chart entry via one of its "selectors": name or numeric ID.
// @tags         Charts
// @produce      json
// @param        selector                 path      string  true  "The Chart to get's selector: name or numeric ID"
// @success      200                      {object}  v2controllers.Chart
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/charts/get/{selector} [get]
func getChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleGet(&controller.ImmutableModelController)
}

// editChart godoc
// @summary      Edit a Chart entry
// @description  Edit an existing Chart entry via one of its "selectors": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
// @tags         Charts
// @accept       json
// @produce      json
// @param        selector                 path      string                       true  "The Chart to edit's selector: name or numeric ID"
// @param        chart                    body      v2controllers.EditableChart  true  "The edits to make to the Chart"
// @success      200                      {object}  v2controllers.Chart
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/charts/edit/{selector} [patch]
func editChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

// deleteChart godoc
// @summary      Delete a Chart entry
// @description  Delete an existing Chart entry via one of its "selectors": name or numeric ID.
// @tags         Charts
// @produce      json
// @param        selector                 path      string  true  "The Chart to delete's selector: name or numeric ID"
// @success      200                      {object}  v2controllers.Chart
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/charts/delete/{selector} [delete]
func deleteChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

// listChartSelectors godoc
// @summary      List Chart selectors
// @description  Validate a given Chart selector and provide any other selectors that would match the same Chart.
// @tags         Charts
// @produce      json
// @param        selector                 path      string  true  "The selector of the Chart to list other selectors for"
// @success      200                      {array}   string
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/charts/selectors/{selector} [get]
func listChartSelectors(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleSelectorList(&controller.ImmutableModelController)
}

// listChart godoc
// @summary      List Chart entries
// @description  List existing Chart entries, ordered by most recently updated.
// @tags         Charts
// @produce      json
// @param        limit                    query     int  false  "An optional limit to the number of entries returned"
// @success      200                      {array}   v2controllers.Chart
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/charts/list [get]
func listChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleList(&controller.ImmutableModelController)
}

// listChartWithFilter godoc
// @summary      List Chart entries with field filters
// @description  List existing Chart entries, ordered by most recently updated. Entries will be filtered to only return ones matching the provided non-empty fields in the body.
// @tags         Charts
// @accept       json
// @produce      json
// @param        limit                    query     int                  false  "An optional limit to the number of entries returned"
// @param        chart                    body      v2controllers.Chart  true   "The fields and values to filter on (omit a field to not filter based on it)"
// @success      200                      {array}   v2controllers.Chart
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/charts/list [post]
func listChartWithFilter(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleListWithFilter(&controller.ImmutableModelController)
}
