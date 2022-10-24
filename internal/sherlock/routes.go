package sherlock

import (
	"github.com/broadinstitute/sherlock/docs"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/handlers/misc"
	"github.com/broadinstitute/sherlock/internal/handlers/v1handlers"
	"github.com/broadinstitute/sherlock/internal/handlers/v2handlers"
	"github.com/broadinstitute/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/internal/version"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title       Sherlock
// @description The Data Science Platform's source-of-truth service.
// @description Note: this API will try to load and return associations in responses, so clients won't need to make as many requests. This behavior isn't recursive, though, so associations of associations are *not* fully loaded (even if it might seem that way from looking at the data types).
// @version     development
// @schemes     https
// @accept      json
// @produce     json

// @contact.name  DSP DevOps
// @contact.email dsp-devops@broadinstitute.org

// @license.name BSD-3-Clause
// @license.url  https://github.com/broadinstitute/sherlock/blob/main/LICENSE.txt

// buildRouter attaches a Gin router with API, Swagger, and other endpoints to
// an existing Application instance. This exists outside of Application itself
// so that Application instances can be more easily tested without the complexity
// of running a full server.
func (a *Application) buildRouter() {
	authMiddleware := auth.IapUserMiddleware

	docs.SwaggerInfo.Version = version.BuildVersion
	if config.Config.MustString("mode") == "debug" {
		// if a dev build, allow http on Swagger page for localhost usage
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		authMiddleware = auth.FakeUserMiddleware
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logger())
	// No origins allowed by default
	router.Use(cors.New(cors.DefaultConfig()))
	// Browsers shouldn't ever cache Sherlock responses
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-store")
	})

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/swagger/index.html") })

	// register generic handlers just on /*
	router.GET("/version", misc.VersionHandler)
	router.GET("/my-user", authMiddleware(), misc.MyUserHandler)
	router.GET("/status", misc.StatusHandler)

	// register v1 API handlers on both /* and /api/v1/*
	v1api := router.Group("api/v1")
	for _, group := range []*gin.RouterGroup{&router.RouterGroup, v1api} {
		// /services routes
		servicesGroup := group.Group("/services")
		v1handlers.RegisterServiceHandlers(servicesGroup, a.Services)

		// /builds routes
		buildsGroup := group.Group("/builds")
		v1handlers.RegisterBuildHandlers(buildsGroup, a.Builds)

		// /environments routes
		environmentsGroup := group.Group("/environments")
		v1handlers.RegisterEnvironmentHandlers(environmentsGroup, a.Environments)

		// /deploys routes
		deploysGroup := group.Group("/deploys")
		v1handlers.RegisterDeployHandlers(deploysGroup, a.Deploys)

		// /metrics route
		metricsGroup := group.Group("/metrics")
		metrics.RegisterPrometheusMetricsHandler(metricsGroup)
	}

	// register v2 API handlers on /api/v2/*
	v2api := router.Group("api/v2")
	v2api.Use(authMiddleware())
	v2handlers.RegisterClusterHandlers(v2api, a.v2controllers.ClusterController)
	v2handlers.RegisterEnvironmentHandlers(v2api, a.v2controllers.EnvironmentController)
	v2handlers.RegisterChartHandlers(v2api, a.v2controllers.ChartController)
	v2handlers.RegisterChartVersionHandlers(v2api, a.v2controllers.ChartVersionController)
	v2handlers.RegisterAppVersionHandlers(v2api, a.v2controllers.AppVersionController)
	v2handlers.RegisterChartReleaseHandlers(v2api, a.v2controllers.ChartReleaseController)
	v2handlers.RegisterChangesetHandlers(v2api, a.v2controllers.ChangesetController)

	a.Handler = router
}
