package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterDatabaseInstanceHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.DatabaseInstanceController) {
	routerGroup.POST("/database-instances", createDatabaseInstance(controller))
	routerGroup.GET("/database-instances", listDatabaseInstance(controller))
	routerGroup.GET("/database-instances/*selector", getDatabaseInstance(controller))
	routerGroup.PATCH("/database-instances/*selector", editDatabaseInstance(controller))
	routerGroup.PUT("/database-instances/*selector", upsertDatabaseInstance(controller))
	routerGroup.DELETE("/database-instances/*selector", deleteDatabaseInstance(controller))
	routerGroup.GET("/selectors/database-instances/*selector", listDatabaseInstanceSelectors(controller))
}

// createDatabaseInstance godoc
//
//	@summary		Create a new DatabaseInstance entry
//	@description	Create a new DatabaseInstance entry. Note that some fields are immutable after creation; /edit lists mutable fields.
//	@tags			DatabaseInstances
//	@accept			json
//	@produce		json
//	@param			chart					body		v2controllers.CreatableDatabaseInstance	true	"The DatabaseInstance to create"
//	@success		200,201					{object}	v2controllers.DatabaseInstance
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/database-instances [post]
func createDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listDatabaseInstance godoc
//
//	@summary		List DatabaseInstance entries
//	@description	List existing DatabaseInstance entries, ordered by most recently updated.
//	@tags			DatabaseInstances
//	@produce		json
//	@param			filter					query		v2controllers.DatabaseInstance	false	"Optional filters to apply to the returned entries"
//	@param			limit					query		int								false	"An optional limit to the number of entries returned"
//	@success		200						{array}		v2controllers.DatabaseInstance
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/database-instances [get]
func listDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getDatabaseInstance godoc
//
//	@summary		Get a DatabaseInstance entry
//	@description	Get an existing DatabaseInstance entry via one of its "selectors": numeric ID or 'chart-release/' followed by a chart release selector.
//	@tags			DatabaseInstances
//	@produce		json
//	@param			selector				path		string	true	"The DatabaseInstance to get's selector: numeric ID or 'chart-release/' followed by a chart release selector"
//	@success		200						{object}	v2controllers.DatabaseInstance
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/database-instances/{selector} [get]
func getDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// editDatabaseInstance godoc
//
//	@summary		Edit a DatabaseInstance entry
//	@description	Edit an existing DatabaseInstance entry via one of its "selectors": numeric ID or 'chart-release/' followed by a chart release selector. Note that only mutable fields are available here, immutable fields can only be set using /create.
//	@tags			DatabaseInstances
//	@accept			json
//	@produce		json
//	@param			selector				path		string									true	"The DatabaseInstance to edit's selector: numeric ID or 'chart-release/' followed by a chart release selector"
//	@param			chart					body		v2controllers.EditableDatabaseInstance	true	"The edits to make to the DatabaseInstance"
//	@success		200						{object}	v2controllers.DatabaseInstance
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/database-instances/{selector} [patch]
func editDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

// upsertDatabaseInstance godoc
//
//	@summary		Create or edit a DatabaseInstance entry
//	@description	Create or edit a DatabaseInstance entry. Attempts to edit and will attempt to create upon an error.
//	@description	If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
//	@tags			DatabaseInstances
//	@accept			json
//	@produce		json
//	@param			selector				path		string									true	"The DatabaseInstance to upsert's selector: numeric ID or 'chart-release/' followed by a chart release selector"
//	@param			database-instance		body		v2controllers.CreatableDatabaseInstance	true	"The DatabaseInstance to upsert"
//	@success		200,201					{object}	v2controllers.DatabaseInstance
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/database-instances/{selector} [put]
func upsertDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

// deleteDatabaseInstance godoc
//
//	@summary		Delete a DatabaseInstance entry
//	@description	Delete an existing DatabaseInstance entry via one of its "selectors": numeric ID or 'chart-release/' followed by a chart release selector.
//	@tags			DatabaseInstances
//	@produce		json
//	@param			selector				path		string	true	"The DatabaseInstance to delete's selector: numeric ID or 'chart-release/' followed by a chart release selector"
//	@success		200						{object}	v2controllers.DatabaseInstance
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/database-instances/{selector} [delete]
func deleteDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

// listDatabaseInstanceSelectors godoc
//
//	@summary		List DatabaseInstance selectors
//	@description	Validate a given DatabaseInstance selector and provide any other selectors that would match the same DatabaseInstance.
//	@tags			DatabaseInstances
//	@produce		json
//	@param			selector				path		string	true	"The selector of the DatabaseInstance to list other selectors for"
//	@success		200						{array}		string
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/selectors/database-instances/{selector} [get]
func listDatabaseInstanceSelectors(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
