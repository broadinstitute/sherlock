package sherlock

import (
	"github.com/broadinstitute/sherlock/docs"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/handlers/misc"
	"github.com/broadinstitute/sherlock/internal/handlers/v1handlers"
	"github.com/broadinstitute/sherlock/internal/handlers/v2handlers"
	"github.com/broadinstitute/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/internal/version"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title        Sherlock
// @description  The Data Science Platform's source-of-truth service
// @version      development
// @schemes      https
// @accept       json
// @produce      json

// @contact.name   DSP DevOps
// @contact.email  dsp-devops@broadinstitute.org

// @license.name  BSD-3-Clause
// @license.url   https://github.com/broadinstitute/sherlock/blob/main/LICENSE.txt

// buildRouter attaches a Gin router with API, Swagger, and other endpoints to
// an existing Application instance. This exists outside of Application itself
// so that Application instances can be more easily tested without the complexity
// of running a full server.
func (a *Application) buildRouter() {
	authMiddleware := auth.IdentityAwareProxyAuthentication

	docs.SwaggerInfo.Version = version.BuildVersion
	if version.BuildVersion == version.DevelopmentVersionString {
		// if a dev build, allow http on Swagger page for localhost usage
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		// if a dev build, skip IAP
		authMiddleware = auth.DummyAuthentication
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logger())

	// register generic handlers just on /*
	router.GET("/version", misc.VersionHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/swagger/index.html") })
	router.GET("/user", authMiddleware(), misc.UserHandler)

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
	v2handlers.RegisterClusterHandlers(v2api.Group("/clusters"), a.V2Clusters)
	v2handlers.RegisterEnvironmentHandlers(v2api.Group("/environments"), a.V2Environments)
	v2handlers.RegisterChartHandlers(v2api.Group("/charts"), a.V2Charts)
	v2handlers.RegisterChartVersionHandlers(v2api.Group("/chart-versions"), a.V2ChartVersions)
	v2handlers.RegisterAppVersionHandlers(v2api.Group("/app-versions"), a.V2AppVersions)
	v2handlers.RegisterChartReleaseHandlers(v2api.Group("/chart-releases"), a.V2ChartReleases)
	v2handlers.RegisterChartDeployRecordHandlers(v2api.Group("/chart-deploy-records"), a.V2ChartDeployRecords)

	a.Handler = router
}
