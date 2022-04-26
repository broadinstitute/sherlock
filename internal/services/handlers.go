package services

// handlers.go contains all the logic for parsing requests and sending responses for
// the /builds api group. No business logic or database logic should be present in this file.

import (
	"errors"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
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
	services, err := sc.ListAll()

	if err != nil {
		// send error response to client
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Services: sc.serialize(services...)})
}

func (sc *ServiceController) getServiceByName(c *gin.Context) {
	name := c.Param("name")
	service, err := sc.GetByName(name)
	if err != nil {
		// return 404 if service is not found, return 500 if some other error
		switch err {
		case v1models.ErrServiceNotFound:
			c.JSON(http.StatusNotFound, Response{Error: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, Response{Services: sc.serialize(service)})
}

func (sc *ServiceController) createService(c *gin.Context) {
	var newService v1models.CreateServiceRequest

	// decode the post request body into a Service struct
	if err := c.BindJSON(&newService); err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: ErrBadCreateRequest.Error()})
		return
	}

	// the create method returns a service struct with the newly saved entity including fields
	// updated internally by the database such as ID
	savedService, err := sc.CreateNew(newService)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Response{Services: sc.serialize(savedService)})
}
