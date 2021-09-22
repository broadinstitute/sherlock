package sherlock

import (
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
	}

	a.Handler = router
}
