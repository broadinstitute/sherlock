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

	// /services routes
	services := router.Group("/services")
	// The empty strings here mean these handlers process requests to /services path
	// the group feature is nice for organization as the more endpoints are added
	services.GET("", a.getServices)
	services.GET("/:id", a.getServiceByID)
	services.POST("", a.createService)

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

func (a *Application) getServiceByID(c *gin.Context) {
	c.JSON(http.StatusOK, services.Service{})
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
