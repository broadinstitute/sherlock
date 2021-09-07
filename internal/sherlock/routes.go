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

	// /services routes
	servicesGroup := router.Group("/services")
	a.Services.RegisterHandlers(servicesGroup)

	a.Handler = router
}
