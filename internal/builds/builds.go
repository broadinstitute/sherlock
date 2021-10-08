// Package builds contains the definitions for a control plan for sherlock's
// build management systems. It also defines api routes for that control plane
package builds

// builds.go contains the "business" logic for operations relating to build entities.
// Thhis could eventually be moved to it's own sub-folder if it becomes unwieldy

import (
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

// CreateBuildRequest is the type used to validate a request for a new build
type CreateBuildRequest struct {
	VersionString string    `json:"version_string" binding:"required"`
	CommitSha     string    `json:"commit_sha" binding:"required"`
	BuildURL      string    `json:"build_url,omitempty"`
	BuiltAt       time.Time `json:"built_at,omitempty"`
	ServiceName   string    `json:"service_name" binding:"required"`
	ServiceRepo   string    `json:"service_repo"`
}

func (bc *BuildController) validateAndCreateNewBuild(newBuild CreateBuildRequest) (Build, error) {
	var serviceID int
	serviceID, err := bc.services.FindOrCreate(newBuild.ServiceName)
	if err != nil {
		return Build{}, ErrBadCreateRequest
	}
	build := Build{
		VersionString: newBuild.VersionString,
		CommitSha:     newBuild.CommitSha,
		BuildURL:      newBuild.BuildURL,
		BuiltAt:       newBuild.BuiltAt,
		ServiceID:     serviceID,
	}

	return bc.store.createNew(build)
}

// ListAll is the public API on the build controller for listing out all builds
func (bc *BuildController) ListAll() ([]Build, error) {
	return bc.store.listAll()

}

// CreateNew is the Public API on the build controller for saving a new build entity
// to persistent storage
func (bc *BuildController) CreateNew(newBuild CreateBuildRequest) (Build, error) {
	return bc.validateAndCreateNewBuild(newBuild)
}

// GetByID is the public api on the build controller for performing a lookup of
// a build entity by ID
func (bc *BuildController) GetByID(id int) (Build, error) {
	return bc.store.getByID(id)
}

// GetByVersionString will perform a look up of a build entity using it's unique version string
// ie image repo + tag
func (bc *BuildController) GetByVersionString(versionString string) (Build, error) {
	return bc.store.getByVersionString(versionString)
}

func (bc *BuildController) serialize(builds ...Build) []BuildResponse {
	var buildsList []Build
	buildsList = append(buildsList, builds...)

	serializer := BuildsSerializer{Builds: buildsList}
	return serializer.Response()
}
