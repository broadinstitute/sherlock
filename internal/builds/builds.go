package builds

import (
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/internal/services"
	"gorm.io/gorm"
)

// BuildController is the management layer that processes requests
// to the /builds api group
type BuildController struct {
	store buildStore
	// this is needed so that we can automatically create a new service entity
	// if a build is reported for a service not tracked by sherlock
	services *services.ServiceController
}

// NewController returns an instance of the controller struct for
// interacting with build entities. It embeds a buildStore interface for
// operations on the build persistence layer
func NewController(dbConn *gorm.DB) *BuildController {
	buildStore := newBuildStore(dbConn)
	return &BuildController{
		store:    buildStore,
		services: services.NewController(dbConn),
	}
}

// Response is a type that allows all data returned from the /builds api group to share a consistent structure
type Response struct {
	Builds []Build `json:"builds"`
	Error  string  `json:"error,omitempty"`
}

// CreateBuildRequest is the type used to validate a request for a new build
type CreateBuildRequest struct {
	VersionString string    `json:"version_string" binding:"required"`
	CommitSha     string    `json:"commit_sha" binding:"required"`
	BuildURL      string    `json:"build_url,omitempty"`
	BuiltAt       time.Time `json:"built_at,omitempty"`
	ServiceName   string    `json:"service_name" binding:"required"`
	ServiceRepo   string    `json:"service_repo"`
}

func (bc *BuildController) validateAndCreateNewBuild(newBuild CreateBuildRequest) (*Build, error) {
	var serviceID int
	serviceID, ok := bc.services.DoesServiceExist(newBuild.ServiceName)
	// handle case where service doesn't exist already
	if !ok {
		var err error
		serviceID, err = bc.createNewServiceFromBuildRequest(newBuild.ServiceName, newBuild.ServiceRepo)
		if err != nil {
			return nil, err
		}
	}
	build := &Build{
		VersionString: newBuild.VersionString,
		CommitSha:     newBuild.CommitSha,
		BuildURL:      newBuild.BuildURL,
		BuiltAt:       newBuild.BuiltAt,
		ServiceID:     serviceID,
	}

	return bc.CreateNew(build)
}

// when receiving a new build for a service sherlock isn't aware of, create that service
func (bc *BuildController) createNewServiceFromBuildRequest(name, repoURL string) (int, error) {
	newServiceRequest := services.CreateServiceRequest{
		Name:    name,
		RepoURL: repoURL,
	}
	newService, err := bc.services.CreateNew(newServiceRequest)
	if err != nil {
		return 0, fmt.Errorf("error creating new service from build request: %v", err)
	}
	return newService.ID, nil
}

// ListAll is the public API on the build controller for listing out all builds
func (bc *BuildController) ListAll() ([]Build, error) {
	return bc.store.listAll()
}

// CreateNew is the Public API on the build controller for saving a new build entity
// to persistent storage
func (bc *BuildController) CreateNew(build *Build) (*Build, error) {
	return bc.store.createNew(build)
}

// GetByID is the public api on the build controller for performing a lookup of
// a build entity by ID
func (bc *BuildController) GetByID(id int) (*Build, error) {
	return bc.store.getByID(id)
}
