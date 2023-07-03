package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAppVersionHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.AppVersionController) {
	routerGroup.POST("/app-versions", createAppVersion(controller))
	routerGroup.GET("/app-versions", listAppVersion(controller))
	routerGroup.GET("/app-versions/*selector", getAppVersion(controller))
	routerGroup.PATCH("/app-versions/*selector", editAppVersion(controller))
	routerGroup.PUT("/app-versions/*selector", upsertAppVersion(controller))
	routerGroup.GET("/selectors/app-versions/*selector", listAppVersionSelectors(controller))
	routerGroup.GET("/procedures/app-versions/children-path-to-parent", getAppVersionChildrenPathToParent(controller))
}

// createAppVersion godoc
//
//	@summary		Create a new AppVersion entry
//	@description	Create a new AppVersion entry. Note that fields are immutable after creation.
//	@description	If the new entry is a duplicate of one already in the database, the database will not be altered and the call will return normally but with a 200 code.
//	@tags			AppVersions
//	@accept			json
//	@produce		json
//	@param			app-version				body		v2controllers.CreatableAppVersion	true	"The AppVersion to create"
//	@success		200,201					{object}	v2controllers.AppVersion
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/app-versions [post]
func createAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleCreate(controller.ModelController)
}

// listAppVersion godoc
//
//	@summary		List AppVersion entries
//	@description	List existing AppVersion entries, ordered by most recently updated.
//	@tags			AppVersions
//	@produce		json
//	@param			filter					query		v2controllers.AppVersion	false	"Optional filters to apply to the returned entries"
//	@param			limit					query		int							false	"An optional limit to the number of entries returned"
//	@success		200						{array}		v2controllers.AppVersion
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/app-versions [get]
func listAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleList(controller.ModelController)
}

// getAppVersion godoc
//
//	@summary		Get a AppVersion entry
//	@description	Get an existing AppVersion entry via one its "selectors": chart/version or numeric ID.
//	@tags			AppVersions
//	@produce		json
//	@param			selector				path		string	true	"The AppVersion to get's selector: chart/version or numeric ID"
//	@success		200						{object}	v2controllers.AppVersion
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/app-versions/{selector} [get]
func getAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleGet(controller.ModelController)
}

// editAppVersion godoc
//
//	@summary		Edit a AppVersion entry
//	@description	Edit an existing AppVersion entry via one its "selectors": chart/version or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
//	@tags			AppVersions
//	@accept			json
//	@produce		json
//	@param			selector				path		string								true	"The AppVersion to edit's selector: chart/version or numeric ID"
//	@param			app-version				body		v2controllers.EditableAppVersion	true	"The edits to make to the AppVersion"
//	@success		200						{object}	v2controllers.AppVersion
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/app-versions/{selector} [patch]
func editAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleEdit(controller.ModelController)
}

// upsertAppVersion godoc
//
//	@summary		Create or edit an AppVersion entry
//	@description	Create or edit an AppVersion entry. Attempts to edit and will attempt to create upon an error.
//	@description	If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
//	@tags			AppVersions
//	@accept			json
//	@produce		json
//	@param			selector				path		string								true	"The AppVersion to upsert's selector: chart/version or numeric ID"
//	@param			app-version				body		v2controllers.CreatableAppVersion	true	"The AppVersion to upsert"
//	@success		200,201					{object}	v2controllers.AppVersion
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/app-versions/{selector} [put]
func upsertAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleUpsert(controller.ModelController)
}

// listAppVersionSelectors godoc
//
//	@summary		List AppVersion selectors
//	@description	Validate a given AppVersion selector and provide any other selectors that would match the same AppVersion.
//	@tags			AppVersions
//	@produce		json
//	@param			selector				path		string	true	"The selector of the AppVersion to list other selectors for"
//	@success		200						{array}		string
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/selectors/app-versions/{selector} [get]
func listAppVersionSelectors(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleSelectorList(controller.ModelController)
}

// getAppVersionChildrenPathToParent godoc
//
//	@summary		Get a changelog between two AppVersions
//	@description	Get the path through parent references from a child AppVersion (inclusive) to a parent AppVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer AppVersion should be the child. The result will always exclude the parent. If the child can't be connected to the parent, just the child will be returned with a 204 code.
//	@tags			AppVersions
//	@produce		json
//	@param			child					query		string	true	"The selector of the newer AppVersion for the changelog"
//	@param			parent					query		string	true	"The selector of the older AppVersion for the changelog"
//	@success		200,204					{array}		v2controllers.AppVersion
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/procedures/app-versions/children-path-to-parent [get]
func getAppVersionChildrenPathToParent(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleGetChildrenPathToParent(controller)
}
