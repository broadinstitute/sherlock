package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func registerCiIdentifierHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.CiIdentifierController) {
	routerGroup.POST("/ci-identifiers", createCiIdentifier(controller))
	routerGroup.GET("/ci-identifiers", listCiIdentifier(controller))
	routerGroup.GET("/ci-identifiers/*selector", getCiIdentifier(controller))
	routerGroup.PATCH("/ci-identifiers/*selector", editCiIdentifier(controller))
	routerGroup.PUT("/ci-identifiers/*selector", upsertCiIdentifier(controller))
	routerGroup.GET("/selectors/ci-identifiers/*selector", listCiIdentifierSelectors(controller))
}

func createCiIdentifier(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

func listCiIdentifier(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleList(controller)
}

func getCiIdentifier(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleGet(controller)
}

func editCiIdentifier(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

func upsertCiIdentifier(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

func listCiIdentifierSelectors(controller *v2controllers.CiIdentifierController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
