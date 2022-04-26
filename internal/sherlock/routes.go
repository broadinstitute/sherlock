package sherlock

import (
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
	api := router.Group("api/v1")

	// register handlers for both /* and /api/v1/*
	for _, group := range []*gin.RouterGroup{&router.RouterGroup, api} {
		// /services routes
		servicesGroup := group.Group("/services")
		a.Services.RegisterHandlers(servicesGroup)

		// /builds routes
		buildsGroup := group.Group("/builds")
		a.Builds.RegisterHandlers(buildsGroup)

		// environments routes
		environmentsGroup := group.Group("/environments")
		a.Environments.RegisterHandlers(environmentsGroup)

		// deploys routes
		deploysGroup := group.Group("/deploys")
		a.Deploys.RegisterHandlers(deploysGroup)

		// metrics routes
		metricsGroup := group.Group("/metrics")
		metrics.RegisterPrometheusMetricsHandler(metricsGroup)

	}

	v2api := router.Group("api/v2")
	v2api.Handle("GET", "/version", func(c *gin.Context) {
		c.String(200, "%s", version.BuildVersion)
	})
	a.Handler = router
}
