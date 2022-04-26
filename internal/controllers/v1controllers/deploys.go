package v1controllers

import (
	"errors"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"

	"gorm.io/gorm"
)

var (
	// ErrServiceMismatch is an error returned when creating a deploy where the build and service instance
	// reference different service entities.
	ErrServiceMismatch = errors.New("service referenced by build and service instance do not match")
)

// DeployController is a type used to contain all the top level functionality for managing
// Deploy entities.
type DeployController struct {
	store            v1models.DeployStore
	ServiceInstances *ServiceInstanceController
	Builds           *BuildController
}

// NewDeployController accepts a gorm db connection and returns
// a controller struct used for the management of deploy entities
func NewDeployController(dbConn *gorm.DB) *DeployController {
	return &DeployController{
		store:            v1models.NewDeployStore(dbConn),
		ServiceInstances: NewServiceInstanceController(dbConn),
		Builds:           NewBuildController(dbConn),
	}
}

// CreateDeployRequest is a struct used to contain all the information
// that is necessary to provision a new deploy
type CreateDeployRequest struct {
	EnvironmentName    string
	ServiceName        string
	BuildVersionString string
}

// CreateNew is used to create a new deploy based on a service name, environment name and build
// version string
func (dc *DeployController) CreateNew(newDeployRequest CreateDeployRequest) (v1models.Deploy, error) {
	// look up the service instance associated with this deploy
	serviceInstanceID, err := dc.ServiceInstances.FindOrCreate(newDeployRequest.EnvironmentName, newDeployRequest.ServiceName)

	if err != nil {
		return v1models.Deploy{}, err
	}

	// look up the build based on provided version string
	build, err := dc.Builds.GetByVersionString(newDeployRequest.BuildVersionString)
	// for now just error if not exists
	if err != nil {
		// create the build if not exists
		newBuild := CreateBuildRequest{
			VersionString: newDeployRequest.BuildVersionString,
			ServiceName:   newDeployRequest.ServiceName,
		}

		build, err = dc.Builds.CreateNew(newBuild)
		if err != nil {
			return v1models.Deploy{}, err
		}
	}

	return dc.store.CreateDeploy(build.ID, serviceInstanceID)
}

// GetDeploysByEnvironmentAndService will retrieve the deploy history for a given service instance with the associated names
func (dc *DeployController) GetDeploysByEnvironmentAndService(environmentName, serviceName string) ([]v1models.Deploy, error) {
	// look up the service instance for the provided service and environment names
	serviceInstance, err := dc.ServiceInstances.GetByEnvironmentAndServiceName(environmentName, serviceName)
	if err != nil {
		return []v1models.Deploy{}, v1models.ErrServiceInstanceNotFound
	}

	return dc.store.GetDeploysByServiceInstance(serviceInstance.ID)
}

// GetMostRecentDeploy will look up the most recent ie currently active deploy for a given service instance
func (dc *DeployController) GetMostRecentDeploy(environmentName, serviceName string) (v1models.Deploy, error) {
	serviceInstance, err := dc.ServiceInstances.GetByEnvironmentAndServiceName(environmentName, serviceName)
	if err != nil {
		return v1models.Deploy{}, v1models.ErrServiceInstanceNotFound
	}

	return dc.store.GetMostRecentDeployByServiceInstance(serviceInstance.ID)
}

func (dc *DeployController) ListServiceInstances() ([]v1models.ServiceInstance, error) {
	return dc.ServiceInstances.ListAll()
}

// Serialize takes a variable number of deploy entities and serializes them into types suitable for use in
// client responses
func (dc *DeployController) Serialize(deploy ...v1models.Deploy) []v1serializers.DeployResponse {
	var deployList []v1models.Deploy
	deployList = append(deployList, deploy...)

	serializer := v1serializers.DeploysSerializer{deployList}
	return serializer.Response()
}
