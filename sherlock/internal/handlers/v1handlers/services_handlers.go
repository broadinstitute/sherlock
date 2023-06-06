package v1handlers

// environments_handlers.go contains all the logic for parsing requests and sending responses for
// the /builds api group. No business logic or database logic should be present in this file.

import (
	"errors"
	"github.com/broadinstitute/sherlock/internal/controllers/v1controllers"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrBadServiceCreateRequest is an error type used when a create servie request fails validation checks
var ErrBadServiceCreateRequest error = errors.New("error invalid create service request. service name and repo url are required")

// RegisterServiceHandlers accepts a gin router  and will attach handlers for working with
// Service entities to it
func RegisterServiceHandlers(routerGroup *gin.RouterGroup, sc *v1controllers.ServiceController) {
	routerGroup.GET("", getServices(sc))
	routerGroup.GET("/:name", getServiceByName(sc))
	routerGroup.POST("", createService(sc))
}

func getServices(sc *v1controllers.ServiceController) func(c *gin.Context) {
	return func(c *gin.Context) {
		services, err := sc.ListAll()

		if err != nil {
			// send error response to client
			c.JSON(http.StatusInternalServerError, v1serializers.Response{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, v1serializers.Response{Services: sc.Serialize(services...)})
	}
}

func getServiceByName(sc *v1controllers.ServiceController) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("name")
		service, err := sc.GetByName(name)
		if err != nil {
			// return 404 if service is not found, return 500 if some other error
			switch err {
			case v1models.ErrServiceNotFound:
				c.JSON(http.StatusNotFound, v1serializers.Response{Error: err.Error()})
				return
			default:
				c.JSON(http.StatusInternalServerError, v1serializers.Response{Error: err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, v1serializers.Response{Services: sc.Serialize(service)})
	}
}

func createService(sc *v1controllers.ServiceController) func(c *gin.Context) {
	return func(c *gin.Context) {
		var newService v1models.CreateServiceRequest

		// decode the post request body into a Service struct
		if err := c.BindJSON(&newService); err != nil {
			c.JSON(http.StatusBadRequest, v1serializers.Response{Error: ErrBadServiceCreateRequest.Error()})
			return
		}

		// the create method returns a service struct with the newly saved entity including fields
		// updated internally by the database such as ID
		savedService, err := sc.CreateNew(newService)

		if err != nil {
			c.JSON(http.StatusInternalServerError, v1serializers.Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, v1serializers.Response{Services: sc.Serialize(savedService)})
	}
}
