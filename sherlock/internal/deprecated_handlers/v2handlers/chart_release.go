package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func registerChartReleaseHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ChartReleaseController) {
	routerGroup.POST("/chart-releases", createChartRelease(controller))
	routerGroup.GET("/chart-releases", listChartRelease(controller))
	routerGroup.GET("/chart-releases/*selector", getChartRelease(controller))
	routerGroup.PATCH("/chart-releases/*selector", editChartRelease(controller))
	routerGroup.PUT("/chart-releases/*selector", upsertChartRelease(controller))
	routerGroup.DELETE("/chart-releases/*selector", deleteChartRelease(controller))
	routerGroup.GET("/selectors/chart-releases/*selector", listChartReleaseSelectors(controller))
	routerGroup.POST("/procedures/chart-releases/trigger-incident/*selector", triggerChartReleasePagerdutyIncident(controller))
}

func createChartRelease(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

func listChartRelease(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleList(controller)
}

func getChartRelease(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleGet(controller)
}

func editChartRelease(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

func upsertChartRelease(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

func deleteChartRelease(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

func listChartReleaseSelectors(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}

func triggerChartReleasePagerdutyIncident(controller *v2controllers.ChartReleaseController) func(ctx *gin.Context) {
	return handleTriggerPagerdutyIncident(controller)
}
