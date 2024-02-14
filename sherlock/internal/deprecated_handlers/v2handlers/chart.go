package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func registerChartHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ChartController) {
	routerGroup.POST("/charts", createChart(controller))
	routerGroup.GET("/charts", listChart(controller))
	routerGroup.GET("/charts/*selector", getChart(controller))
	routerGroup.PATCH("/charts/*selector", editChart(controller))
	routerGroup.PUT("/charts/*selector", upsertChart(controller))
	routerGroup.DELETE("/charts/*selector", deleteChart(controller))
	routerGroup.GET("/selectors/charts/*selector", listChartSelectors(controller))
}

func createChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

func listChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleList(controller)
}

func getChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleGet(controller)
}

func editChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

func upsertChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

func deleteChart(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

func listChartSelectors(controller *v2controllers.ChartController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
