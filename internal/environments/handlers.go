package environments

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrBadCreateRequest is an error type used when a create service request fails validation checks
var ErrBadCreateRequest error = errors.New("error invalid create environment request. environment name is required")

// RegisterHandlers accepts a routergroup and will attach all the handlers for
// working with environment entities to it
func (ec *EnvironmentController) RegisterHandlers(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", ec.getEnvironments)
	routerGroup.GET("/:name", ec.getEnvironmentByName)
	routerGroup.POST("", ec.createEnvironment)
}

func (ec *EnvironmentController) getEnvironments(c *gin.Context) {
	environments, err := ec.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Environments: environments})
}

func (ec *EnvironmentController) getEnvironmentByName(c *gin.Context) {
	name := c.Param("name")

	environment, err := ec.GetByName(name)
	if err != nil {
		switch err {
		case ErrEnvironmentNotFound:
			c.JSON(http.StatusNotFound, Response{Error: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, Response{Environments: []EnvironmentResponse{environment}})
}

func (ec *EnvironmentController) createEnvironment(c *gin.Context) {
	var newEnvironment CreateEnvironmentRequest

	if err := c.BindJSON(&newEnvironment); err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: ErrBadCreateRequest.Error()})
		return
	}

	savedEnvironment, err := ec.CreateNew(newEnvironment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Response{Environments: []EnvironmentResponse{savedEnvironment}})
}
