package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func registerAppVersionHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.AppVersionController) {
	routerGroup.POST("/app-versions", createAppVersion(controller))
	routerGroup.GET("/app-versions", listAppVersion(controller))
	routerGroup.GET("/app-versions/*selector", getAppVersion(controller))
	routerGroup.PATCH("/app-versions/*selector", editAppVersion(controller))
	routerGroup.PUT("/app-versions/*selector", upsertAppVersion(controller))
	routerGroup.GET("/selectors/app-versions/*selector", listAppVersionSelectors(controller))
	routerGroup.GET("/procedures/app-versions/children-path-to-parent", getAppVersionChildrenPathToParent(controller))
}

func createAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleCreate(controller.ModelController)
}

func listAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleList(controller.ModelController)
}

func getAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleGet(controller.ModelController)
}

func editAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleEdit(controller.ModelController)
}

func upsertAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleUpsert(controller.ModelController)
}

func listAppVersionSelectors(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleSelectorList(controller.ModelController)
}

func getAppVersionChildrenPathToParent(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleGetChildrenPathToParent(controller)
}
