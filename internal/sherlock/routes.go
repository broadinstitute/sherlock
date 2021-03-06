package sherlock

import (
	"github.com/broadinstitute/sherlock/internal/handlers/v1handlers"
	"github.com/broadinstitute/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/internal/version"
	"github.com/gin-gonic/gin"
)

// helper function to take an existing sherlock application instance
// then build and attach a gin router to it.
// this decouples building the router from instanting a sherlock Application
// which makes testing easier
func (a *Application) buildRouter() {
	router := gin.Default()

	// register basic non-API handlers just on /*
	router.Handle("GET", "/version", func(c *gin.Context) {
		c.String(200, "%s", version.BuildVersion)
	})

	// register handlers for both /* and /api/v1/*
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
	a.Handler = router
}
