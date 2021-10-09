package deploys

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrBadCreateRequest is an error type used when a create servie request fails validation checks
var ErrBadCreateRequest error = errors.New("error invalid create deploy request")

// RegisterHandlers accepts a gin router group and attaches handlers for working
// with deploy entities
func (dc *DeployController) RegisterHandlers(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/:environment/:service", dc.getDeploysByEnvironmentAndService)
	routerGroup.POST("/:environment/:service", dc.createDeploy)
}

func (dc *DeployController) getDeploysByEnvironmentAndService(c *gin.Context) {
	environment := c.Param("environment")
	service := c.Param("service")

	deploys, err := dc.GetDeploysByEnvironmentAndService(environment, service)
	if err != nil {
		switch err {
		case ErrServiceInstanceNotFound:
			c.JSON(http.StatusNotFound, Response{Error: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, Response{Deploys: dc.Serialize(deploys...)})
}

// CreateDeployRequestBody is used to extract any additional
// data needed to create a new deploy from the request body.
// currently this is just the version string as the environment
// and service are determined from path params
type CreateDeployRequestBody struct {
	VersionString string `json:"version_string" binding:"required"`
}

func (dc *DeployController) createDeploy(c *gin.Context) {
	environment := c.Param("environment")
	service := c.Param("service")

	var deployRequestBody CreateDeployRequestBody
	if err := c.BindJSON(&deployRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: ErrBadCreateRequest.Error()})
		return
	}

	// construct the create deploy request from params and body
	newDeployRequest := CreateDeployRequest{
		EnvironmentName:    environment,
		ServiceName:        service,
		BuildVersionString: deployRequestBody.VersionString,
	}

	deploy, err := dc.CreateNew(newDeployRequest)
	if err != nil {
		if errors.Is(err, ErrServiceMismatch) {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Response{Deploys: dc.Serialize(deploy)})
}
