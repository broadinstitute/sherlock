package services

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrBadCreateRequest is an error type used when a create servie request fails validation checks
var ErrBadCreateRequest error = errors.New("error invalid create service request. service name and repo url are required")

// RegisterHandlers accepts a gin router  and will attach handlers for working with
// Service entities to it
func (sc *ServiceController) RegisterHandlers(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", sc.getServices)
	routerGroup.GET("/:name", sc.getServiceByName)
	routerGroup.POST("", sc.createService)
}

func (sc *ServiceController) getServices(c *gin.Context) {

	services, err := sc.store.listAll()
	if err != nil {
		// send error response to client
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Services: services})
}

func (sc *ServiceController) getServiceByName(c *gin.Context) {
	name := c.Param("name")
	service, err := sc.store.getByName(name)
	if err != nil {
		// return 404 if service is not found, return 500 if some other error
		switch err {
		case ErrServiceNotFound:
			c.JSON(http.StatusNotFound, Response{Error: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, Response{Services: []*Service{service}})
}

func (sc *ServiceController) createService(c *gin.Context) {
	var newService CreateServiceRequest

	// decode the post request body into a Service struct
	if err := c.BindJSON(&newService); err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: ErrBadCreateRequest.Error()})
		return
	}

	// the create method returns a service struct with the newly saved entity including fields
	// updated internally by the database such as ID
	savedService, err := sc.store.createNew(newService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Response{Services: []*Service{savedService}})
}
