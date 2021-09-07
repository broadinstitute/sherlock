package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (sc *ServiceController) RegisterHandlers(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", sc.getServices)
	routerGroup.GET("/:id", sc.getServiceByID)
	routerGroup.POST("", sc.createService)
}

func (sc *ServiceController) getServices(c *gin.Context) {

	services, err := sc.store.listAll()
	if err != nil {
		// send error response to client
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

func (sc *ServiceController) getServiceByID(c *gin.Context) {
	id := c.Param("id")
	service, err := sc.store.getByID(id)
	if err != nil {
		// return 404 if service is not found, return 500 if some other error
		switch err {
		case ErrServiceNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, service)
}

func (sc *ServiceController) createService(c *gin.Context) {
	var newService Service

	// decode the post request body into a Service struct
	if err := c.BindJSON(&newService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// the create method returns a service struct with the newly saved entity including fields
	// updated internally by the database such as ID
	savedService, err := sc.store.createNew(&newService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, savedService)
}
