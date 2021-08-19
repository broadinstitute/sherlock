package sherlock

import (
	"net/http"

	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/gin-gonic/gin"
)

// helper function to take an existing sherlock application instance
// then build and attach a gin router to it.
// this decouples building the router from instanting a sherlock Application
// which makes testing easier
func (a *Application) buildRouter() {
	router := gin.Default()
	router.GET("/services", a.getServices)
	router.POST("/services", a.createService)
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

func (a *Application) createService(c *gin.Context) {
	var newService services.Service

	// decode the post request body into a Service struct
	if err := c.BindJSON(&newService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// the create method returns a service struct with the newly saved entity including fields
	// updated internally by the database such as ID
	savedService, err := a.ServiceModel.Create(&newService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, savedService)
}
