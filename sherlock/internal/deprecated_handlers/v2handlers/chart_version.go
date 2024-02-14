package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func registerChartVersionHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ChartVersionController) {
	routerGroup.POST("/chart-versions", createChartVersion(controller))
	routerGroup.GET("/chart-versions", listChartVersion(controller))
	routerGroup.GET("/chart-versions/*selector", getChartVersion(controller))
	routerGroup.PATCH("/chart-versions/*selector", editChartVersion(controller))
	routerGroup.PUT("/chart-versions/*selector", upsertChartVersion(controller))
	routerGroup.GET("/selectors/chart-versions/*selector", listChartVersionSelectors(controller))
	routerGroup.GET("/procedures/chart-versions/children-path-to-parent", getChartVersionChildrenPathToParent(controller))
}

func createChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleCreate(controller.ModelController)
}

func listChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleList(controller.ModelController)
}

func getChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleGet(controller.ModelController)
}

func editChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleEdit(controller.ModelController)
}

func upsertChartVersion(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleUpsert(controller.ModelController)
}

func listChartVersionSelectors(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleSelectorList(controller.ModelController)
}

func getChartVersionChildrenPathToParent(controller *v2controllers.ChartVersionController) func(ctx *gin.Context) {
	return handleGetChildrenPathToParent(controller)
}
