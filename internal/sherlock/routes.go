package sherlock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// helper function to take an existing sherlock application instance
// then build and attach a gin router to it.
// this decouples building the router from instanting a sherlock Application
// which makes testing easier
func (a *Application) buildRouter() {
	router := gin.Default()
	router.GET("/services", a.getServices)
	a.Handler = router
}

func (a *Application) getServices(c *gin.Context) {

	services, err := a.ServiceModel.ListAll()
	if err != nil {
		// send error response to client
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}
