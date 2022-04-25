package deploys

import (
	"errors"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/broadinstitute/sherlock/internal/models"
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
	store            models.DeployStore
	serviceInstances *ServiceInstanceController
	builds           *builds.BuildController
}

// NewDeployController accepts a gorm db connection and returns
// a controller struct used for the management of deploy entities
func NewDeployController(dbConn *gorm.DB) *DeployController {
	return &DeployController{
		store:            models.NewDeployStore(dbConn),
		serviceInstances: NewServiceInstanceController(dbConn),
		builds:           builds.NewController(dbConn),
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
func (dc *DeployController) CreateNew(newDeployRequest CreateDeployRequest) (models.Deploy, error) {
	// look up the service instance associated with this deploy
	serviceInstanceID, err := dc.serviceInstances.FindOrCreate(newDeployRequest.EnvironmentName, newDeployRequest.ServiceName)

	if err != nil {
		return models.Deploy{}, err
	}

	// look up the build based on provided version string
	build, err := dc.builds.GetByVersionString(newDeployRequest.BuildVersionString)
	// for now just error if not exists
	if err != nil {
		// create the build if not exists
		newBuild := builds.CreateBuildRequest{
			VersionString: newDeployRequest.BuildVersionString,
			ServiceName:   newDeployRequest.ServiceName,
		}

		build, err = dc.builds.CreateNew(newBuild)
		if err != nil {
			return models.Deploy{}, err
		}
	}

	return dc.store.CreateDeploy(build.ID, serviceInstanceID)
}

// GetDeploysByEnvironmentAndService will retrieve the deploy history for a given service instance with the associated names
func (dc *DeployController) GetDeploysByEnvironmentAndService(environmentName, serviceName string) ([]models.Deploy, error) {
	// look up the service instance for the provided service and environment names
	serviceInstance, err := dc.serviceInstances.GetByEnvironmentAndServiceName(environmentName, serviceName)
	if err != nil {
		return []models.Deploy{}, models.ErrServiceInstanceNotFound
	}

	return dc.store.GetDeploysByServiceInstance(serviceInstance.ID)
}

// GetMostRecentDeploy will look up the most recent ie currently active deploy for a given service instance
func (dc *DeployController) GetMostRecentDeploy(environmentName, serviceName string) (models.Deploy, error) {
	serviceInstance, err := dc.serviceInstances.GetByEnvironmentAndServiceName(environmentName, serviceName)
	if err != nil {
		return models.Deploy{}, models.ErrServiceInstanceNotFound
	}

	return dc.store.GetMostRecentDeployByServiceInstance(serviceInstance.ID)
}

func (dc *DeployController) ListServiceInstances() ([]models.ServiceInstance, error) {
	return dc.serviceInstances.ListAll()
}

// Serialize takes a variable number of deploy entities and serializes them into types suitable for use in
// client responses
func (dc *DeployController) Serialize(deploy ...models.Deploy) []DeployResponse {
	var deployList []models.Deploy
	deployList = append(deployList, deploy...)

	serializer := DeploysSerializer{deployList}
	return serializer.Response()
}
