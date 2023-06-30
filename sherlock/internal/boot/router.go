package boot

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/version"
	"github.com/broadinstitute/sherlock/sherlock/docs"
	"github.com/broadinstitute/sherlock/sherlock/html"
	"github.com/broadinstitute/sherlock/sherlock/internal/boot/middleware"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/controllers/v2controllers"
	"github.com/broadinstitute/sherlock/sherlock/internal/handlers/misc"
	"github.com/broadinstitute/sherlock/sherlock/internal/handlers/v2handlers"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/v2models"
	"github.com/gin-gonic/gin"
	swaggo_files "github.com/swaggo/files"
	swaggo_gin "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"net/http"
)

//	@title			Sherlock
//	@description	The Data Science Platform's source-of-truth service.
//	@description	Note: this API will try to load and return associations in responses, so clients won't need to make as many requests. This behavior isn't recursive, though, so associations of associations are *not* fully loaded (even if it might seem that way from looking at the data types).
//	@version		development
//	@schemes		https
//	@accept			json
//	@produce		json

//	@contact.name	DSP DevOps
//	@contact.email	dsp-devops@broadinstitute.org

//	@license.name	BSD-3-Clause
//	@license.url	https://github.com/broadinstitute/sherlock/blob/main/LICENSE.txt

func buildRouter(db *gorm.DB) (*gin.Engine, error) {
	// gin.DebugMode spews console output but can help resolve routing issues
	gin.SetMode(gin.ReleaseMode)

	docs.SwaggerInfo.Version = version.BuildVersion
	if config.Config.String("mode") == "debug" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	}

	router := gin.New()
	router.Use(gin.Recovery(), middleware.Logger(), middleware.Headers())
	authMiddleware := middleware.Auth(v2models.NewMiddlewareUserStore(db))

	router.GET("/version", misc.VersionHandler)
	router.GET("/status", misc.StatusHandler)
	router.GET("/my-user", authMiddleware, misc.MyUserHandler)

	metrics.RegisterPrometheusMetricsHandler(router.Group("/metrics"))

	router.StaticFS("/static", http.FS(html.StaticHtmlFiles))

	router.GET("/swagger/*any", swaggo_gin.WrapHandler(swaggo_files.Handler))
	router.GET("", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/swagger/index.html") })

	apiRoutes := router.Group("api", authMiddleware, middleware.DB(db))
	apiV2Routes := apiRoutes.Group("v2")
	controllers := v2controllers.NewControllerSet(v2models.NewStoreSet(db))
	v2handlers.RegisterClusterHandlers(apiV2Routes, controllers.ClusterController)
	v2handlers.RegisterEnvironmentHandlers(apiV2Routes, controllers.EnvironmentController)
	v2handlers.RegisterChartHandlers(apiV2Routes, controllers.ChartController)
	v2handlers.RegisterChartVersionHandlers(apiV2Routes, controllers.ChartVersionController)
	v2handlers.RegisterAppVersionHandlers(apiV2Routes, controllers.AppVersionController)
	v2handlers.RegisterChartReleaseHandlers(apiV2Routes, controllers.ChartReleaseController)
	v2handlers.RegisterChangesetHandlers(apiV2Routes, controllers.ChangesetController)
	v2handlers.RegisterPagerdutyIntegrationHandlers(apiV2Routes, controllers.PagerdutyIntegrationController)
	v2handlers.RegisterDatabaseInstanceHandlers(apiV2Routes, controllers.DatabaseInstanceController)
	v2handlers.RegisterUserHandlers(apiV2Routes, controllers.UserController)
	v2handlers.RegisterCiIdentifierHandlers(apiV2Routes, controllers.CiIdentifierController)
	v2handlers.RegisterCiRunHandlers(apiV2Routes, controllers.CiRunController)

	return router, nil
}
