package v2handlers

import (
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterChartVersionHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ChartVersionController) {
	routerGroup.POST("/chart-versions", createChartVersion(controller))
	routerGroup.GET("/chart-versions", listChartVersion(controller))
	routerGroup.GET("/chart-versions/*selector", getChartVersion(controller))
	routerGroup.PATCH("/chart-versions/*selector", editChartVersion(controller))
	routerGroup.PUT("/chart-versions/*selector", upsertChartVersion(controller))
	routerGroup.GET("/selectors/chart-versions/*selector", listChartVersionSelectors(controller))
}

// createChartVersion godoc
//
//	@summary		Create a new ChartVersion entry
//	@description	Create a new ChartVersion entry. Note that fields are immutable after creation.
//	@description	If the new entry is a duplicate of one already in the database, the database will not be altered and the call will return normally but with a 200 code.
//	@tags			ChartVersions
//	@accept			json
//	@produce		json
//	@param			chart-version			body		v2controllers.CreatableChartVersion	true	"The ChartVersion to create"
//	@success		200,201					{object}	v2controllers.ChartVersion
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/chart-versions [post]
func createChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listChartVersion godoc
//
//	@summary		List ChartVersion entries
//	@description	List existing ChartVersion entries, ordered by most recently updated.
//	@tags			ChartVersions
//	@produce		json
//	@param			filter					query		v2controllers.ChartVersion	false	"Optional filters to apply to the returned entries"
//	@param			limit					query		int							false	"An optional limit to the number of entries returned"
//	@success		200						{array}		v2controllers.ChartVersion
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/chart-versions [get]
func listChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getChartVersion godoc
//
//	@summary		Get a ChartVersion entry
//	@description	Get an existing ChartVersion entry via one its "selectors": chart/version or numeric ID.
//	@tags			ChartVersions
//	@produce		json
//	@param			selector				path		string	true	"The ChartVersion to get's selector: chart/version or numeric ID"
//	@success		200						{object}	v2controllers.ChartVersion
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/chart-versions/{selector} [get]
func getChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// editChartVersion godoc
//
//	@summary		Edit a ChartVersion entry
//	@description	Edit an existing ChartVersion entry via one its "selectors": chart/version or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
//	@tags			ChartVersions
//	@accept			json
//	@produce		json
//	@param			selector				path		string								true	"The ChartVersion to edit's selector: chart/version or numeric ID"
//	@param			chart-version			body		v2controllers.EditableChartVersion	true	"The edits to make to the ChartVersion"
//	@success		200						{object}	v2controllers.ChartVersion
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/chart-versions/{selector} [patch]
func editChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

// upsertChartVersion godoc
//
//	@summary		Create or edit a ChartVersion entry
//	@description	Create or edit a ChartVersion entry. Attempts to edit and will attempt to create upon an error.
//	@description	If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
//	@tags			ChartVersions
//	@accept			json
//	@produce		json
//	@param			selector				path		string								true	"The ChartVersion to upsert's selector: chart/version or numeric ID"
//	@param			chart-version			body		v2controllers.CreatableChartVersion	true	"The ChartVersion to upsert"
//	@success		200,201					{object}	v2controllers.ChartVersion
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/chart-versions/{selector} [put]
func upsertChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

// listChartVersionSelectors godoc
//
//	@summary		List ChartVersion selectors
//	@description	Validate a given ChartVersion selector and provide any other selectors that would match the same ChartVersion.
//	@tags			ChartVersions
//	@produce		json
//	@param			selector				path		string	true	"The selector of the ChartVersion to list other selectors for"
//	@success		200						{array}		string
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/selectors/chart-versions/{selector} [get]
func listChartVersionSelectors(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
