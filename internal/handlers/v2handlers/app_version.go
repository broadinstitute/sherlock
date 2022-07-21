package v2handlers

import (
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAppVersionHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.AppVersionController) {
	routerGroup.POST("/create", createAppVersion(controller))
	routerGroup.GET("/get/*selector", getAppVersion(controller))
	routerGroup.GET("/selectors/*selector", listAppVersionSelectors(controller))
	routerGroup.GET("/list", listAppVersion(controller))
	routerGroup.POST("/list", listAppVersionWithFilter(controller))
}

// createAppVersion godoc
// @summary      Create a new AppVersion entry
// @description  Create a new AppVersion entry. Note that fields are immutable after creation.
// @tags         AppVersions
// @accept       json
// @produce      json
// @param        app-version              body      v2controllers.CreatableAppVersion  true  "The AppVersion to create"
// @success      200                      {object}  v2controllers.AppVersion
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/app-versions/create [post]
func createAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// getAppVersion godoc
// @summary      Get a AppVersion entry
// @description  Get an existing AppVersion entry via one its "selector"--its numeric ID.
// @tags         AppVersions
// @produce      json
// @param        selector                 path      string  true  "The AppVersion to get's selector: name or numeric ID"
// @success      200                      {object}  v2controllers.AppVersion
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/app-versions/get/{selector} [get]
func getAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// listAppVersionSelectors godoc
// @summary      List AppVersion selectors
// @description  Validate a given AppVersion selector and provide any other selectors that would match the same AppVersion.
// @tags         AppVersions
// @produce      json
// @param        selector                 path      string  true  "The selector of the AppVersion to list other selectors for"
// @success      200                      {array}   string
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/app-versions/selectors/{selector} [get]
func listAppVersionSelectors(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}

// listAppVersion godoc
// @summary      List AppVersion entries
// @description  List existing AppVersion entries, ordered by most recently updated.
// @tags         AppVersions
// @produce      json
// @param        limit                    query     int  false  "An optional limit to the number of entries returned"
// @success      200                      {array}   v2controllers.AppVersion
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/app-versions/list [get]
func listAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleList(controller)
}

// listAppVersionWithFilter godoc
// @summary      List AppVersion entries with field filters
// @description  List existing AppVersion entries, ordered by most recently updated. Entries will be filtered to only return ones matching the provided non-empty fields in the body.
// @tags         AppVersions
// @accept       json
// @produce      json
// @param        limit                    query     int                       false  "An optional limit to the number of entries returned"
// @param        app-version              body      v2controllers.AppVersion  true   "The fields and values to filter on (omit a field to not filter based on it)"
// @success      200                      {array}   v2controllers.AppVersion
// @failure      400,403,404,407,409,500  {object}  errors.ErrorResponse
// @router       /api/v2/app-versions/list [post]
func listAppVersionWithFilter(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleListWithFilter(controller)
}
