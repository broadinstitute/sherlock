package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func registerEnvironmentHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.EnvironmentController) {
	routerGroup.POST("/environments", createEnvironment(controller))
	routerGroup.GET("/environments", listEnvironment(controller))
	routerGroup.GET("/environments/*selector", getEnvironment(controller))
	routerGroup.PATCH("/environments/*selector", editEnvironment(controller))
	routerGroup.PUT("/environments/*selector", upsertEnvironment(controller))
	routerGroup.DELETE("/environments/*selector", deleteEnvironment(controller))
	routerGroup.GET("/selectors/environments/*selector", listEnvironmentSelectors(controller))
	routerGroup.POST("/procedures/environments/trigger-incident/*selector", triggerEnvironmentPagerdutyIncident(controller))
}

func createEnvironment(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

func listEnvironment(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleList(controller)
}

func getEnvironment(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleGet(controller)
}

func editEnvironment(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

func upsertEnvironment(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

func deleteEnvironment(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

func listEnvironmentSelectors(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}

func triggerEnvironmentPagerdutyIncident(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleTriggerPagerdutyIncident(controller)
}
