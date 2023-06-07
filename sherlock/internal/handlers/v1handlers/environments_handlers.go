package v1handlers

import (
	"errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/controllers/v1controllers"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/v1models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrBadEnvironmentCreateRequest is an error type used when a create environment request fails validation checks
var ErrBadEnvironmentCreateRequest error = errors.New("error invalid create environment request. environment name is required")

// RegisterEnvironmentHandlers accepts a routergroup and will attach all the handlers for
// working with environment entities to it
func RegisterEnvironmentHandlers(routerGroup *gin.RouterGroup, ec *v1controllers.EnvironmentController) {
	routerGroup.GET("", getEnvironments(ec))
	routerGroup.GET("/:name", getEnvironmentByName(ec))
	routerGroup.POST("", createEnvironment(ec))
}

func getEnvironments(ec *v1controllers.EnvironmentController) func(c *gin.Context) {
	return func(c *gin.Context) {
		environments, err := ec.ListAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, v1controllers.Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, v1controllers.Response{Environments: ec.Serialize(environments...)})
	}
}

func getEnvironmentByName(ec *v1controllers.EnvironmentController) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("name")

		environment, err := ec.GetByName(name)
		if err != nil {
			switch err {
			case v1models.ErrEnvironmentNotFound:
				c.JSON(http.StatusNotFound, v1controllers.Response{Error: err.Error()})
				return
			default:
				c.JSON(http.StatusInternalServerError, v1controllers.Response{Error: err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, v1controllers.Response{Environments: ec.Serialize(environment)})
	}
}

func createEnvironment(ec *v1controllers.EnvironmentController) func(c *gin.Context) {
	return func(c *gin.Context) {
		var newEnvironment v1models.CreateEnvironmentRequest

		if err := c.BindJSON(&newEnvironment); err != nil {
			c.JSON(http.StatusBadRequest, v1controllers.Response{Error: ErrBadEnvironmentCreateRequest.Error()})
			return
		}

		savedEnvironment, err := ec.CreateNew(newEnvironment)

		if err != nil {
			c.JSON(http.StatusInternalServerError, v1controllers.Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, v1controllers.Response{Environments: ec.Serialize(savedEnvironment)})
	}
}
