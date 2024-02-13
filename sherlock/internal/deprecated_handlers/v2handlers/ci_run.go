package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func registerCiRunHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.CiRunController) {
	routerGroup.POST("/ci-runs", createCiRun(controller))
	routerGroup.GET("/ci-runs", listCiRun(controller))
	routerGroup.GET("/ci-runs/*selector", getCiRun(controller))
	routerGroup.PATCH("/ci-runs/*selector", editCiRun(controller))
	routerGroup.PUT("/ci-runs/*selector", upsertCiRun(controller))
	routerGroup.DELETE("/ci-runs/*selector", deleteCiRun(controller))
	routerGroup.GET("/selectors/ci-runs/*selector", listCiRunSelectors(controller))
}

func createCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

func listCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleList(controller)
}

func getCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleGet(controller)
}

func editCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

func upsertCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

func deleteCiRun(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

func listCiRunSelectors(controller *v2controllers.CiRunController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
