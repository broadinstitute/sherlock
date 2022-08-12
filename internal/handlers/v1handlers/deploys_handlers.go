package v1handlers

import (
	"errors"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/controllers/v1controllers"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"
	"net/http"

	"github.com/broadinstitute/sherlock/internal/metrics"
	"github.com/gin-gonic/gin"
)

// ErrBadDeployCreateRequest is an error type used when a create service request fails validation checks
var ErrBadDeployCreateRequest error = errors.New("error invalid create deploy request")

// RegisterDeployHandlers accepts a gin router group and attaches handlers for working
// with deploy entities
func RegisterDeployHandlers(routerGroup *gin.RouterGroup, dc *v1controllers.DeployController) {
	routerGroup.GET("/:environment/:service", getDeploysByEnvironmentAndService(dc))
	routerGroup.POST("/:environment/:service", createDeploy(dc))
}

func getDeploysByEnvironmentAndService(dc *v1controllers.DeployController) func(c *gin.Context) {
	return func(c *gin.Context) {
		environment := c.Param("environment")
		service := c.Param("service")

		deploys, err := dc.GetDeploysByEnvironmentAndService(environment, service)
		if err != nil {
			switch err {
			case v1models.ErrServiceInstanceNotFound:
				c.JSON(http.StatusNotFound, v1serializers.DeploysResponse{Error: err.Error()})
				return
			default:
				c.JSON(http.StatusInternalServerError, v1serializers.DeploysResponse{Error: err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, v1serializers.DeploysResponse{Deploys: dc.Serialize(deploys...)})
	}
}

// CreateDeployRequestBody is used to extract any additional
// data needed to create a new deploy from the request body.
// currently this is just the version string as the environment
// and service are determined from path params
type CreateDeployRequestBody struct {
	VersionString string `json:"version_string" binding:"required"`
}

func createDeploy(dc *v1controllers.DeployController) func(c *gin.Context) {
	return func(c *gin.Context) {
		environment := c.Param("environment")
		service := c.Param("service")

		var deployRequestBody CreateDeployRequestBody
		if err := c.BindJSON(&deployRequestBody); err != nil {
			c.JSON(http.StatusBadRequest, v1serializers.DeploysResponse{Error: ErrBadDeployCreateRequest.Error()})
			return
		}

		// construct the create deploy request from params and body
		newDeployRequest := v1controllers.CreateDeployRequest{
			EnvironmentName:    environment,
			ServiceName:        service,
			BuildVersionString: deployRequestBody.VersionString,
		}

		// flag used to track if this deploy event is a no-op redeploy in which case lead time shouldn't change
		shouldUpdateLeadTime := true
		// look up the current active deployment for the give service instance
		currentDeploy, err := dc.GetMostRecentDeploy(newDeployRequest.EnvironmentName, newDeployRequest.ServiceName)
		if err != nil && err != v1models.ErrDeployNotFound {
			c.JSON(http.StatusInternalServerError, v1serializers.DeploysResponse{Error: err.Error()})
		}

		// if doing a redeploy of the current build, don't update leadtime
		if currentDeploy.Build.VersionString == newDeployRequest.BuildVersionString {
			shouldUpdateLeadTime = false
		}

		deploy, err := dc.CreateNew(newDeployRequest)
		if err != nil {
			if errors.Is(err, v1controllers.ErrServiceMismatch) {
				c.JSON(http.StatusBadRequest, v1serializers.DeploysResponse{Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, v1serializers.DeploysResponse{Error: err.Error()})
			return
		}

		if config.Config.String("metrics.accelerate.fromAPI") != "v2" {
			metrics.RecordDeployFrequency(c, environment, service)
			// calculate lead time
			if shouldUpdateLeadTime {
				leadTime := deploy.CalculateLeadTimeHours()
				metrics.RecordLeadTime(c, leadTime, environment, service)
			}
		}

		c.JSON(http.StatusCreated, v1serializers.DeploysResponse{Deploys: dc.Serialize(deploy)})
	}
}
