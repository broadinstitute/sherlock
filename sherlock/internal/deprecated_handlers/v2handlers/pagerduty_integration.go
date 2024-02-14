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

func createPagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

func listPagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleList(controller)
}

func getPagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleGet(controller)
}

func editPagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

func upsertPagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

func deletePagerdutyIntegration(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

func listPagerdutyIntegrationSelectors(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}

func triggerPagerdutyIntegrationPagerdutyIncident(controller *v2controllers.PagerdutyIntegrationController) func(ctx *gin.Context) {
	return handleTriggerPagerdutyIncident(controller)
}
