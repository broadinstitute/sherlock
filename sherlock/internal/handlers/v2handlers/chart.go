package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterChartHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ChartController) {
	routerGroup.POST("/charts", createChart(controller))
	routerGroup.GET("/charts", listChart(controller))
	routerGroup.GET("/charts/*selector", getChart(controller))
	routerGroup.PATCH("/charts/*selector", editChart(controller))
	routerGroup.PUT("/charts/*selector", upsertChart(controller))
	routerGroup.DELETE("/charts/*selector", deleteChart(controller))
	routerGroup.GET("/selectors/charts/*selector", listChartSelectors(controller))
}

// createChart godoc
//
//	@summary		Create a new Chart entry
//	@description	Create a new Chart entry. Note that some fields are immutable after creation; /edit lists mutable fields.
//	@tags			Charts
//	@accept			json
//	@produce		json
//	@param			chart					body		v2controllers.CreatableChart	true	"The Chart to create"
//	@success		200,201					{object}	v2controllers.Chart
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/charts [post]
func createChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listChart godoc
//
//	@summary		List Chart entries
//	@description	List existing Chart entries, ordered by most recently updated.
//	@tags			Charts
//	@produce		json
//	@param			filter					query		v2controllers.Chart	false	"Optional filters to apply to the returned entries"
//	@param			limit					query		int					false	"An optional limit to the number of entries returned"
//	@success		200						{array}		v2controllers.Chart
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/charts [get]
func listChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getChart godoc
//
//	@summary		Get a Chart entry
//	@description	Get an existing Chart entry via one of its "selectors": name or numeric ID.
//	@tags			Charts
//	@produce		json
//	@param			selector				path		string	true	"The Chart to get's selector: name or numeric ID"
//	@success		200						{object}	v2controllers.Chart
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/charts/{selector} [get]
func getChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// editChart godoc
//
//	@summary		Edit a Chart entry
//	@description	Edit an existing Chart entry via one of its "selectors": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
//	@tags			Charts
//	@accept			json
//	@produce		json
//	@param			selector				path		string						true	"The Chart to edit's selector: name or numeric ID"
//	@param			chart					body		v2controllers.EditableChart	true	"The edits to make to the Chart"
//	@success		200						{object}	v2controllers.Chart
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/charts/{selector} [patch]
func editChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

// upsertChart godoc
//
//	@summary		Create or edit a Chart entry
//	@description	Create or edit a Chart entry. Attempts to edit and will attempt to create upon an error.
//	@description	If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
//	@tags			Charts
//	@accept			json
//	@produce		json
//	@param			selector				path		string							true	"The Chart to upsert's selector: name or numeric ID"
//	@param			chart					body		v2controllers.CreatableChart	true	"The Chart to upsert"
//	@success		200,201					{object}	v2controllers.Chart
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/charts/{selector} [put]
func upsertChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

// deleteChart godoc
//
//	@summary		Delete a Chart entry
//	@description	Delete an existing Chart entry via one of its "selectors": name or numeric ID.
//	@tags			Charts
//	@produce		json
//	@param			selector				path		string	true	"The Chart to delete's selector: name or numeric ID"
//	@success		200						{object}	v2controllers.Chart
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/charts/{selector} [delete]
func deleteChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

// listChartSelectors godoc
//
//	@summary		List Chart selectors
//	@description	Validate a given Chart selector and provide any other selectors that would match the same Chart.
//	@tags			Charts
//	@produce		json
//	@param			selector				path		string	true	"The selector of the Chart to list other selectors for"
//	@success		200						{array}		string
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/selectors/charts/{selector} [get]
func listChartSelectors(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
