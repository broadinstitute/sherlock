package deploys

import (
	"errors"
	"net/http"

	"github.com/broadinstitute/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/internal/models"
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
		case models.ErrServiceInstanceNotFound:
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

	// flag used to track if this deploy event is a no-op redeploy in which case lead time shouldn't change
	shouldUpdateLeadTime := true
	// look up the current active deployment for the give service instance
	currentDeploy, err := dc.GetMostRecentDeploy(newDeployRequest.EnvironmentName, newDeployRequest.ServiceName)
	if err != nil && err != models.ErrDeployNotFound {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
	}

	// if doing a redeploy of the current build, don't update leadtime
	if currentDeploy.Build.VersionString == newDeployRequest.BuildVersionString {
		shouldUpdateLeadTime = false
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

	metrics.RecordDeployFrequency(c, environment, service)
	// calculate lead time
	if shouldUpdateLeadTime {
		leadTime := deploy.CalculateLeadTimeHours()
		metrics.RecordLeadTime(c, leadTime, environment, service)
	}

	c.JSON(http.StatusCreated, Response{Deploys: dc.Serialize(deploy)})
}
