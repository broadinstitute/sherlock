package v2handlers

import (
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterEnvironmentHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.EnvironmentController) {
	routerGroup.POST("/environments", createEnvironment(controller))
	routerGroup.GET("/environments", listEnvironment(controller))
	routerGroup.GET("/environments/*selector", getEnvironment(controller))
	routerGroup.PATCH("/environments/*selector", editEnvironment(controller))
	routerGroup.DELETE("/environments/*selector", deleteEnvironment(controller))
	routerGroup.GET("/selectors/environments/*selector", listEnvironmentSelectors(controller))
}

// createEnvironment godoc
// @summary      Create a new Environment entry
// @description  Create a new Environment entry. Note that some fields are immutable after creation; /edit lists mutable fields.
// @description  Creating a dynamic environment based on a template will also copy ChartReleases from the template.
// @tags         Environments
// @accept       json
// @produce      json
// @param        environment              body      v2controllers.CreatableEnvironment  true  "The Environment to create"
// @success      200                      {object}  v2controllers.Environment
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/environments [post]
func createEnvironment(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listEnvironment godoc
// @summary      List Environment entries
// @description  List existing Environment entries, ordered by most recently updated.
// @tags         Environments
// @produce      json
// @param        filter                   query     v2controllers.Environment  false  "Optional filters to apply to the returned entries"
// @param        limit                    query     int                        false  "An optional limit to the number of entries returned"
// @success      200                      {array}   v2controllers.Environment
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/environments [get]
func listEnvironment(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getEnvironment godoc
// @summary      Get a Environment entry
// @description  Get an existing Environment entry via one of its "selectors": name or numeric ID.
// @tags         Environments
// @produce      json
// @param        selector                 path      string  true  "The Environment to get's selector: name or numeric ID"
// @success      200                      {object}  v2controllers.Environment
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/environments/{selector} [get]
func getEnvironment(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// editEnvironment godoc
// @summary      Edit a Environment entry
// @description  Edit an existing Environment entry via one of its "selectors": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
// @tags         Environments
// @accept       json
// @produce      json
// @param        selector                 path      string                             true  "The Environment to edit's selector: name or numeric ID"
// @param        environment              body      v2controllers.EditableEnvironment  true  "The edits to make to the Environment"
// @success      200                      {object}  v2controllers.Environment
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/environments/{selector} [patch]
func editEnvironment(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleEdit(controller)
}

// deleteEnvironment godoc
// @summary      Delete a Environment entry
// @description  Delete an existing Environment entry via one of its "selectors": name or numeric ID.
// @tags         Environments
// @produce      json
// @param        selector                 path      string  true  "The Environment to delete's selector: name or numeric ID"
// @success      200                      {object}  v2controllers.Environment
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/environments/{selector} [delete]
func deleteEnvironment(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleDelete(controller)
}

// listEnvironmentSelectors godoc
// @summary      List Environment selectors
// @description  Validate a given Environment selector and provide any other selectors that would match the same Environment.
// @tags         Environments
// @produce      json
// @param        selector                 path      string  true  "The selector of the Environment to list other selectors for"
// @success      200                      {array}   string
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/selectors/environments/{selector} [get]
func listEnvironmentSelectors(controller *v2controllers.EnvironmentController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
