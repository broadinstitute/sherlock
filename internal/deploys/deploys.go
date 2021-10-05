package deploys

import (
	"errors"

	"github.com/broadinstitute/sherlock/internal/builds"
	"gorm.io/gorm"
)

var (
	// ErrBuildNotFound is returned by the create deploy methhod when trying to create
	// a deploy of a build that doesn't already exist
	ErrBuildNotFound = errors.New("unable to create deploy, build does not exist")
)

// DeployController is a type used to contain all the top level functionality for managing
// Deploy entities.
type DeployController struct {
	store           deployStore
	serviceInsances *ServiceInstanceController
	builds          *builds.BuildController
}

// NewDeployController accepts a gorm db connection and returns
// a controller struct used for the management of deploy entities
func NewDeployController(dbConn *gorm.DB) *DeployController {
	return &DeployController{
		store:           newDeployStore(dbConn),
		serviceInsances: NewServiceInstanceController(dbConn),
		builds:          builds.NewController(dbConn),
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
func (dc *DeployController) CreateNew(newDeployRequest CreateDeployRequest) (Deploy, error) {
	// look up the service instance associated with this deploy
	serviceInstanceID, err := dc.serviceInsances.FindOrCreate(newDeployRequest.EnvironmentName, newDeployRequest.ServiceName)
	if err != nil {
		return Deploy{}, err
	}

	// look up the build based on provided version string
	build, err := dc.builds.GetByVersionString(newDeployRequest.BuildVersionString)
	// for now just error if not exists
	if err != nil {
		return Deploy{}, err
	}

	return dc.store.createDeploy(build.ID, serviceInstanceID)
}
