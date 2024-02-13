package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func registerClusterHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ClusterController) {
	routerGroup.POST("/clusters", createCluster(controller))
	routerGroup.GET("/clusters", listCluster(controller))
	routerGroup.GET("/clusters/*selector", getCluster(controller))
	routerGroup.PATCH("/clusters/*selector", editCluster(controller))
	routerGroup.PUT("/clusters/*selector", upsertCluster(controller))
	routerGroup.DELETE("/clusters/*selector", deleteCluster(controller))
	routerGroup.GET("/selectors/clusters/*selector", listClusterSelectors(controller))
}

func createCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

func listCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleList(controller)
}

func getCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleGet(controller)
}

func editCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

func upsertCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

func deleteCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

func listClusterSelectors(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
