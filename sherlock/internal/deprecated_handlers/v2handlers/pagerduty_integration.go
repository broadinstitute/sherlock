package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func registerPagerdutyIntegrationHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.PagerdutyIntegrationController) {
	routerGroup.POST("/pagerduty-integrations", createPagerdutyIntegration(controller))
	routerGroup.GET("/pagerduty-integrations", listPagerdutyIntegration(controller))
	routerGroup.GET("/pagerduty-integrations/*selector", getPagerdutyIntegration(controller))
	routerGroup.PATCH("/pagerduty-integrations/*selector", editPagerdutyIntegration(controller))
	routerGroup.PUT("/pagerduty-integrations/*selector", upsertPagerdutyIntegration(controller))
	routerGroup.DELETE("/pagerduty-integrations/*selector", deletePagerdutyIntegration(controller))
	routerGroup.GET("/selectors/pagerduty-integrations/*selector", listPagerdutyIntegrationSelectors(controller))
	routerGroup.POST("/procedures/pagerduty-integrations/trigger-incident/*selector", triggerPagerdutyIntegrationPagerdutyIncident(controller))
}

// createPagerdutyIntegration godoc
//
//	@summary		Create a new PagerdutyIntegration entry
//	@description	Create a new PagerdutyIntegration entry. Note that fields are immutable after creation.
//	@tags			PagerdutyIntegrations
//	@accept			json
//	@produce		json
//	@param			pagerduty-integration	body		v2controllers.CreatablePagerdutyIntegration	true	"The PagerdutyIntegration to create"
//	@success		200,201					{object}	v2controllers.PagerdutyIntegration
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/pagerduty-integrations [post]
func createPagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listPagerdutyIntegration godoc
//
//	@summary		List PagerdutyIntegration entries
//	@description	List existing PagerdutyIntegration entries, ordered by most recently updated.
//	@tags			PagerdutyIntegrations
//	@produce		json
//	@param			filter					query		v2controllers.PagerdutyIntegration	false	"Optional filters to apply to the returned entries"
//	@param			limit					query		int									false	"An optional limit to the number of entries returned"
//	@success		200						{array}		v2controllers.PagerdutyIntegration
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/pagerduty-integrations [get]
func listPagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getPagerdutyIntegration godoc
//
//	@summary		Get a PagerdutyIntegration entry
//	@description	Get an existing PagerdutyIntegration entry via one its "selectors": "pd-id/" + Pagerduty ID or numeric Sherlock ID.
//	@tags			PagerdutyIntegrations
//	@produce		json
//	@param			selector				path		string	true	"The PagerdutyIntegration to get's selector: chart/version or numeric ID"
//	@success		200						{object}	v2controllers.PagerdutyIntegration
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/pagerduty-integrations/{selector} [get]
func getPagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// editPagerdutyIntegration godoc
//
//	@summary		Edit a PagerdutyIntegration entry
//	@description	Edit an existing PagerdutyIntegration entry via one its "selectors": "pd-id/" + Pagerduty ID or numeric Sherlock ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
//	@tags			PagerdutyIntegrations
//	@accept			json
//	@produce		json
//	@param			selector				path		string										true	"The PagerdutyIntegration to edit's selector: chart/version or numeric ID"
//	@param			pagerduty-integration	body		v2controllers.EditablePagerdutyIntegration	true	"The edits to make to the PagerdutyIntegration"
//	@success		200						{object}	v2controllers.PagerdutyIntegration
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/pagerduty-integrations/{selector} [patch]
func editPagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

// upsertPagerdutyIntegration godoc
//
//	@summary		Create or edit a PagerdutyIntegration entry
//	@description	Create or edit a PagerdutyIntegration entry. Attempts to edit and will attempt to create upon an error.
//	@description	If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
//	@tags			PagerdutyIntegrations
//	@accept			json
//	@produce		json
//	@param			selector				path		string										true	"The PagerdutyIntegration to upsert's selector: "pd-id/" + Pagerduty ID or numeric Sherlock ID"
//	@param			pagerduty-integration	body		v2controllers.CreatablePagerdutyIntegration	true	"The PagerdutyIntegration to upsert"
//	@success		200,201					{object}	v2controllers.PagerdutyIntegration
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/pagerduty-integrations/{selector} [put]
func upsertPagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

// deletePagerdutyIntegration godoc
//
//	@summary		Delete a PagerdutyIntegration entry
//	@description	Delete an existing PagerdutyIntegration entry via one of its "selectors": "pd-id/" + Pagerduty ID or numeric Sherlock ID.
//	@tags			PagerdutyIntegrations
//	@produce		json
//	@param			selector				path		string	true	"The PagerdutyIntegration to delete's selector: "pd-id/" + Pagerduty ID or numeric Sherlock ID"
//	@success		200						{object}	v2controllers.PagerdutyIntegration
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/pagerduty-integrations/{selector} [delete]
func deletePagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

// listPagerdutyIntegrationSelectors godoc
//
//	@summary		List PagerdutyIntegration selectors
//	@description	Validate a given PagerdutyIntegration selector and provide any other selectors that would match the same PagerdutyIntegration.
//	@tags			PagerdutyIntegrations
//	@produce		json
//	@param			selector				path		string	true	"The selector of the PagerdutyIntegration to list other selectors for"
//	@success		200						{array}		string
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/selectors/pagerduty-integrations/{selector} [get]
func listPagerdutyIntegrationSelectors(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}

// triggerPagerdutyIntegrationPagerdutyIncident godoc
//
//	@summary		Trigger a Pagerduty incident for a given PagerdutyIntegration
//	@description	Trigger an alert via a PagerdutyIntegration itself.
//	@tags			PagerdutyIntegrations
//	@accept			json
//	@produce		json
//	@param			selector				path		string					true	"The PagerdutyIntegration's selector"
//	@param			summary					body		pagerduty.AlertSummary	true	"Summary of the incident"
//	@success		202						{object}	pagerduty.SendAlertResponse
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/procedures/pagerduty-integrations/trigger-incident/{selector} [post]
func triggerPagerdutyIntegrationPagerdutyIncident(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleTriggerPagerdutyIncident(controller)
}
