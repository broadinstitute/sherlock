package v2handlers

import (
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterCiIdentifierHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.CiIdentifierController) {
	routerGroup.POST("/ci-identifiers", createCiIdentifier(controller))
	routerGroup.GET("/ci-identifiers", listCiIdentifier(controller))
	routerGroup.GET("/ci-identifiers/*selector", getCiIdentifier(controller))
	routerGroup.PATCH("/ci-identifiers/*selector", editCiIdentifier(controller))
	routerGroup.PUT("/ci-identifiers/*selector", upsertCiIdentifier(controller))
	routerGroup.GET("/selectors/ci-identifiers/*selector", listCiIdentifierSelectors(controller))
}

// createCiIdentifier godoc
//
//	@summary		Create a new CiIdentifier entry
//	@description	Create a new CiIdentifier entry. Note that some fields are immutable after creation; /edit lists mutable fields.
//	@tags			CiIdentifiers
//	@accept			json
//	@produce		json
//	@param			ci-identifier			body		v2controllers.CreatableCiIdentifier	true	"The CiIdentifier to create"
//	@success		200,201					{object}	v2controllers.CiIdentifier
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/ci-identifiers [post]
func createCiIdentifier(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listCiIdentifier godoc
//
//	@summary		List CiIdentifier entries
//	@description	List existing CiIdentifier entries, ordered by most recently updated.
//	@tags			CiIdentifiers
//	@produce		json
//	@param			filter					query		v2controllers.CiIdentifier	false	"Optional filters to apply to the returned entries"
//	@param			limit					query		int							false	"An optional limit to the number of entries returned"
//	@success		200						{array}		v2controllers.CiIdentifier
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/ci-identifiers [get]
func listCiIdentifier(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getCiIdentifier godoc
//
//	@summary		Get a CiIdentifier entry
//	@description	Get an existing CiIdentifier entry via one of its "selectors": ID or type + '/' + selector of the referenced type.
//	@tags			CiIdentifiers
//	@produce		json
//	@param			selector				path		string	true	"The CiIdentifier to get's selector: ID or type + '/' + selector of the referenced type"
//	@success		200						{object}	v2controllers.CiIdentifier
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/ci-identifiers/{selector} [get]
func getCiIdentifier(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// editCiIdentifier godoc
//
//	@summary		Edit a CiIdentifier entry
//	@description	Edit an existing CiIdentifier entry via one of its "selectors": ID or type + '/' + selector of the referenced type. Note that only mutable fields are available here, immutable fields can only be set using /create.
//	@tags			CiIdentifiers
//	@accept			json
//	@produce		json
//	@param			selector				path		string								true	"The CiIdentifier to edit's selector: ID or type + '/' + selector of the referenced type"
//	@param			ci-identifier			body		v2controllers.EditableCiIdentifier	true	"The edits to make to the CiIdentifier"
//	@success		200						{object}	v2controllers.CiIdentifier
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/ci-identifiers/{selector} [patch]
func editCiIdentifier(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

// upsertCiIdentifier godoc
//
//	@summary		Create or edit a CiIdentifier entry
//	@description	Create or edit a CiIdentifier entry. Attempts to edit and will attempt to create upon an error.
//	@description	If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
//	@tags			CiIdentifiers
//	@accept			json
//	@produce		json
//	@param			selector				path		string								true	"The CiIdentifier to upsert's selector: ID or type + '/' + selector of the referenced type"
//	@param			ci-identifier			body		v2controllers.CreatableCiIdentifier	true	"The CiIdentifier to upsert"
//	@success		200,201					{object}	v2controllers.CiIdentifier
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/ci-identifiers/{selector} [put]
func upsertCiIdentifier(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

// listCiIdentifierSelectors godoc
//
//	@summary		List CiIdentifier selectors
//	@description	Validate a given CiIdentifier selector and provide any other selectors that would match the same CiIdentifier.
//	@tags			CiIdentifiers
//	@produce		json
//	@param			selector				path		string	true	"The selector of the CiIdentifier to list other selectors for"
//	@success		200						{array}		string
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/selectors/ci-identifiers/{selector} [get]
func listCiIdentifierSelectors(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
