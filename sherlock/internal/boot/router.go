package boot

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/version"
	"github.com/broadinstitute/sherlock/sherlock/docs"
	"github.com/broadinstitute/sherlock/sherlock/html"
	"github.com/broadinstitute/sherlock/sherlock/internal/api/misc"
	"github.com/broadinstitute/sherlock/sherlock/internal/api/sherlock"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/boot/middleware"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_controllers/v2controllers"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_handlers/v2handlers"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics"
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

func buildRouter(db *gorm.DB) *gin.Engine {
	// gin.DebugMode spews console output but can help resolve routing issues
	gin.SetMode(gin.ReleaseMode)

	docs.SwaggerInfo.Version = version.BuildVersion
	if config.Config.String("mode") == "debug" {
		// When running locally, make the Swagger page have a scheme dropdown with http as the default
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	}

	router := gin.New()
	router.Use(gin.Recovery(), middleware.Logger(), middleware.Headers())

	misc.ConfigureRoutes(&router.RouterGroup)

	router.GET("/metrics", metrics.PrometheusHandler())
	router.StaticFS("/static", http.FS(html.StaticHtmlFiles))
	router.GET("/swagger/*any", swaggo_gin.WrapHandler(swaggo_files.Handler))
	router.GET("", func(ctx *gin.Context) { ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html") })

	apiRouter := router.Group("api", authentication.UserMiddleware(db), authentication.DbMiddleware(db))
	sherlock.ConfigureRoutes(apiRouter)

	v2apiRouter := apiRouter.Group("v2")
	v2ControllerSet := v2controllers.NewControllerSet(v2models.NewStoreSet(db))
	v2handlers.RegisterClusterHandlers(v2apiRouter, v2ControllerSet.ClusterController)
	v2handlers.RegisterEnvironmentHandlers(v2apiRouter, v2ControllerSet.EnvironmentController)
	v2handlers.RegisterChartHandlers(v2apiRouter, v2ControllerSet.ChartController)
	v2handlers.RegisterChartVersionHandlers(v2apiRouter, v2ControllerSet.ChartVersionController)
	v2handlers.RegisterAppVersionHandlers(v2apiRouter, v2ControllerSet.AppVersionController)
	v2handlers.RegisterChartReleaseHandlers(v2apiRouter, v2ControllerSet.ChartReleaseController)
	v2handlers.RegisterChangesetHandlers(v2apiRouter, v2ControllerSet.ChangesetController)
	v2handlers.RegisterPagerdutyIntegrationHandlers(v2apiRouter, v2ControllerSet.PagerdutyIntegrationController)
	v2handlers.RegisterDatabaseInstanceHandlers(v2apiRouter, v2ControllerSet.DatabaseInstanceController)
	v2handlers.RegisterUserHandlers(v2apiRouter, v2ControllerSet.UserController)
	v2handlers.RegisterCiIdentifierHandlers(v2apiRouter, v2ControllerSet.CiIdentifierController)
	v2handlers.RegisterCiRunHandlers(v2apiRouter, v2ControllerSet.CiRunController)

	return router
}
