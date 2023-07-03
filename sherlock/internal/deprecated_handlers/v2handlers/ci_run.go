package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterCiRunHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.CiRunController) {
	routerGroup.POST("/ci-runs", createCiRun(controller))
	routerGroup.GET("/ci-runs", listCiRun(controller))
	routerGroup.GET("/ci-runs/*selector", getCiRun(controller))
	routerGroup.PATCH("/ci-runs/*selector", editCiRun(controller))
	routerGroup.PUT("/ci-runs/*selector", upsertCiRun(controller))
	routerGroup.DELETE("/ci-runs/*selector", deleteCiRun(controller))
	routerGroup.GET("/selectors/ci-runs/*selector", listCiRunSelectors(controller))
}

// createCiRun godoc
//
//	@summary		Create a new CiRun entry
//	@description	Create a new CiRun entry. Note that some fields are immutable after creation; /edit lists mutable fields.
//	@tags			CiRuns
//	@accept			json
//	@produce		json
//	@param			ci-run					body		v2controllers.CreatableCiRun	true	"The CiRun to create"
//	@success		200,201					{object}	v2controllers.CiRun
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/ci-runs [post]
func createCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listCiRun godoc
//
//	@summary		List CiRun entries
//	@description	List existing CiRun entries, ordered by most recently updated.
//	@tags			CiRuns
//	@produce		json
//	@param			filter					query		v2controllers.CiRun	false	"Optional filters to apply to the returned entries"
//	@param			limit					query		int					false	"An optional limit to the number of entries returned"
//	@success		200						{array}		v2controllers.CiRun
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/ci-runs [get]
func listCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getCiRun godoc
//
//	@summary		Get a CiRun entry
//	@description	Get an existing CiRun entry via one of its "selectors": ID, 'github-actions/' + owner + repo + run ID + attempt number, or 'argo-workflows/' + namespace + name.
//	@tags			CiRuns
//	@produce		json
//	@param			selector				path		string	true	"The CiRun to get's selector: ID, 'github-actions/' + owner + repo + run ID + attempt number, or 'argo-workflows/' + namespace + name"
//	@success		200						{object}	v2controllers.CiRun
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/ci-runs/{selector} [get]
func getCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// editCiRun godoc
//
//	@summary		Edit a CiRun entry
//	@description	Edit an existing CiRun entry via one of its "selectors": ID, 'github-actions/' + owner + repo + run ID + attempt number, or 'argo-workflows/' + namespace + name. Note that only mutable fields are available here, immutable fields can only be set using /create.
//	@tags			CiRuns
//	@accept			json
//	@produce		json
//	@param			selector				path		string						true	"The CiRun to edit's selector: ID, 'github-actions/' + owner + repo + run ID + attempt number, or 'argo-workflows/' + namespace + name"
//	@param			ci-run					body		v2controllers.EditableCiRun	true	"The edits to make to the CiRun"
//	@success		200						{object}	v2controllers.CiRun
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/ci-runs/{selector} [patch]
func editCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

// upsertCiRun godoc
//
//	@summary		Create or edit a CiRun entry
//	@description	Create or edit a CiRun entry. Attempts to edit and will attempt to create upon an error.
//	@description	If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
//	@tags			CiRuns
//	@accept			json
//	@produce		json
//	@param			selector				path		string							true	"The CiRun to upsert's selector: ID, 'github-actions/' + owner + repo + run ID + attempt number, or 'argo-workflows/' + namespace + name"
//	@param			ci-run					body		v2controllers.CreatableCiRun	true	"The CiRun to upsert"
//	@success		200,201					{object}	v2controllers.CiRun
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/ci-runs/{selector} [put]
func upsertCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

// deleteCiRun godoc
//
//	@summary		Delete a CiRun entry
//	@description	Delete an existing CiRun entry via one of its "selectors": ID, 'github-actions/' + owner + repo + run ID + attempt number, or 'argo-workflows/' + namespace + name.
//	@tags			CiRuns
//	@produce		json
//	@param			selector				path		string	true	"The CiRun to delete's selector: ID, 'github-actions/' + owner + repo + run ID + attempt number, or 'argo-workflows/' + namespace + name"
//	@success		200						{object}	v2controllers.CiRun
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/ci-runs/{selector} [delete]
func deleteCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

// listCiRunSelectors godoc
//
//	@summary		List CiRun selectors
//	@description	Validate a given CiRun selector and provide any other selectors that would match the same CiRun.
//	@tags			CiRuns
//	@produce		json
//	@param			selector				path		string	true	"The selector of the CiRun to list other selectors for"
//	@success		200						{array}		string
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/selectors/ci-runs/{selector} [get]
func listCiRunSelectors(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
