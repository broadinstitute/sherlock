package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func registerDatabaseInstanceHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.DatabaseInstanceController) {
	routerGroup.POST("/database-instances", createDatabaseInstance(controller))
	routerGroup.GET("/database-instances", listDatabaseInstance(controller))
	routerGroup.GET("/database-instances/*selector", getDatabaseInstance(controller))
	routerGroup.PATCH("/database-instances/*selector", editDatabaseInstance(controller))
	routerGroup.PUT("/database-instances/*selector", upsertDatabaseInstance(controller))
	routerGroup.DELETE("/database-instances/*selector", deleteDatabaseInstance(controller))
	routerGroup.GET("/selectors/database-instances/*selector", listDatabaseInstanceSelectors(controller))
}

func createDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

func listDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleList(controller)
}

func getDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleGet(controller)
}

func editDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

func upsertDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

func deleteDatabaseInstance(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

func listDatabaseInstanceSelectors(controller *v2controllers.DatabaseInstanceController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
