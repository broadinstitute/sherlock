package v2handlers

import (
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAppVersionHandlers(routerGroup *gin.RouterGroup, controller *v2controllers.AppVersionController) {
	routerGroup.POST("/app-versions", createAppVersion(controller))
	routerGroup.GET("/app-versions", listAppVersion(controller))
	routerGroup.GET("/app-versions/*selector", getAppVersion(controller))
	routerGroup.GET("/selectors/app-versions/*selector", listAppVersionSelectors(controller))
}

// createAppVersion godoc
// @summary     Create a new AppVersion entry
// @description Create a new AppVersion entry. Note that fields are immutable after creation.
// @description If the new entry is a duplicate of one already in the database, the database will not be altered and the call will return normally but with a 200 code.
// @tags        AppVersions
// @accept      json
// @produce     json
// @param       app-version             body     v2controllers.CreatableAppVersion true "The AppVersion to create"
// @success     200,201                 {object} v2controllers.AppVersion
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/app-versions [post]
func createAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleCreate(controller)
}

// listAppVersion godoc
// @summary     List AppVersion entries
// @description List existing AppVersion entries, ordered by most recently updated.
// @tags        AppVersions
// @produce     json
// @param       filter                  query    v2controllers.AppVersion false "Optional filters to apply to the returned entries"
// @param       limit                   query    int                      false "An optional limit to the number of entries returned"
// @success     200                     {array}  v2controllers.AppVersion
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/app-versions [get]
func listAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleList(controller)
}

// getAppVersion godoc
// @summary     Get a AppVersion entry
// @description Get an existing AppVersion entry via one its "selectors": chart/version or numeric ID.
// @tags        AppVersions
// @produce     json
// @param       selector                path     string true "The AppVersion to get's selector: chart/version or numeric ID"
// @success     200                     {object} v2controllers.AppVersion
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/app-versions/{selector} [get]
func getAppVersion(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleGet(controller)
}

// listAppVersionSelectors godoc
// @summary     List AppVersion selectors
// @description Validate a given AppVersion selector and provide any other selectors that would match the same AppVersion.
// @tags        AppVersions
// @produce     json
// @param       selector                path     string true "The selector of the AppVersion to list other selectors for"
// @success     200                     {array}  string
// @failure     400,403,404,407,409,500 {object} errors.ErrorResponse
// @router      /api/v2/selectors/app-versions/{selector} [get]
func listAppVersionSelectors(controller *v2controllers.AppVersionController) func(ctx *gin.Context) {
	return handleSelectorList(controller)
}
