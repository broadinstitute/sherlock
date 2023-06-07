package v2handlers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterClusterHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.ClusterController) {
	routerGroup.POST("/clusters", createCluster(controller))
	routerGroup.GET("/clusters", listCluster(controller))
	routerGroup.GET("/clusters/*selector", getCluster(controller))
	routerGroup.PATCH("/clusters/*selector", editCluster(controller))
	routerGroup.PUT("/clusters/*selector", upsertCluster(controller))
	routerGroup.DELETE("/clusters/*selector", deleteCluster(controller))
	routerGroup.GET("/selectors/clusters/*selector", listClusterSelectors(controller))
}

// createCluster godoc
//
//	@summary		Create a new Cluster entry
//	@description	Create a new Cluster entry. Note that some fields are immutable after creation; /edit lists mutable fields.
//	@tags			Clusters
//	@accept			json
//	@produce		json
//	@param			cluster					body		v2controllers.CreatableCluster	true	"The Cluster to create"
//	@success		200,201					{object}	v2controllers.Cluster
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/clusters [post]
func createCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listCluster godoc
//
//	@summary		List Cluster entries
//	@description	List existing Cluster entries, ordered by most recently updated.
//	@tags			Clusters
//	@produce		json
//	@param			filter					query		v2controllers.Cluster	false	"Optional filters to apply to the returned entries"
//	@param			limit					query		int						false	"An optional limit to the number of entries returned"
//	@success		200						{array}		v2controllers.Cluster
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/clusters [get]
func listCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getCluster godoc
//
//	@summary		Get a Cluster entry
//	@description	Get an existing Cluster entry via one of its "selectors": name or numeric ID.
//	@tags			Clusters
//	@produce		json
//	@param			selector				path		string	true	"The Cluster to get's selector: name or numeric ID"
//	@success		200						{object}	v2controllers.Cluster
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/clusters/{selector} [get]
func getCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// editCluster godoc
//
//	@summary		Edit a Cluster entry
//	@description	Edit an existing Cluster entry via one of its "selectors": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
//	@tags			Clusters
//	@accept			json
//	@produce		json
//	@param			selector				path		string							true	"The Cluster to edit's selector: name or numeric ID"
//	@param			cluster					body		v2controllers.EditableCluster	true	"The edits to make to the Cluster"
//	@success		200						{object}	v2controllers.Cluster
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/clusters/{selector} [patch]
func editCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

// upsertCluster godoc
//
//	@summary		Create or edit a Cluster entry
//	@description	Create or edit a Cluster entry. Attempts to edit and will attempt to create upon an error.
//	@description	If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
//	@tags			Clusters
//	@accept			json
//	@produce		json
//	@param			selector				path		string							true	"The Cluster to upsert's selector: name or numeric ID"
//	@param			cluster					body		v2controllers.CreatableCluster	true	"The Cluster to upsert"
//	@success		200,201					{object}	v2controllers.Cluster
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/clusters/{selector} [put]
func upsertCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleUpsert(controller)
}

// deleteCluster godoc
//
//	@summary		Delete a Cluster entry
//	@description	Delete an existing Cluster entry via one of its "selectors": name or numeric ID.
//	@tags			Clusters
//	@produce		json
//	@param			selector				path		string	true	"The Cluster to delete's selector: name or numeric ID"
//	@success		200						{object}	v2controllers.Cluster
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/clusters/{selector} [delete]
func deleteCluster(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

// listClusterSelectors godoc
//
//	@summary		List Cluster selectors
//	@description	Validate a given Cluster selector and provide any other selectors that would match the same Cluster.
//	@tags			Clusters
//	@produce		json
//	@param			selector				path		string	true	"The selector of the Cluster to list other selectors for"
//	@success		200						{array}		string
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/v2/selectors/clusters/{selector} [get]
func listClusterSelectors(controller *v2controllers.ClusterController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
