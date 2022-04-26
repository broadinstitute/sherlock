package v1controllers

// Builds.go contains the "business" logic for operations relating to build entities.
// Thhis could eventually be moved to it's own sub-folder if it becomes unwieldy

import (
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"
	"time"

	"gorm.io/gorm"
)

// BuildController is the management layer that processes requests
// to the /Builds api group
type BuildController struct {
	Store v1models.BuildStore
	// this is needed so that we can automatically create a new service entity
	// if a build is reported for a service not tracked by sherlock
	Services *ServiceController
}

// NewBuildController returns an instance of the controller struct for
// interacting with build entities. It embeds a buildStore interface for
// operations on the build persistence layer
func NewBuildController(dbConn *gorm.DB) *BuildController {
	buildStore := v1models.NewBuildStore(dbConn)
	return &BuildController{
		Store:    buildStore,
		Services: NewServiceController(dbConn),
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

func (bc *BuildController) validateAndCreateNewBuild(newBuild CreateBuildRequest) (v1models.Build, error) {
	var serviceID int
	serviceID, err := bc.Services.FindOrCreate(newBuild.ServiceName)
	if err != nil {
		return v1models.Build{}, v1models.ErrBadCreateRequest
	}
	build := v1models.Build{
		VersionString: newBuild.VersionString,
		CommitSha:     newBuild.CommitSha,
		BuildURL:      newBuild.BuildURL,
		BuiltAt:       newBuild.BuiltAt,
		ServiceID:     serviceID,
	}

	return bc.Store.CreateNew(build)
}

// ListAll is the public API on the build controller for listing out all Builds
func (bc *BuildController) ListAll() ([]v1models.Build, error) {
	return bc.Store.ListAll()

}

// CreateNew is the Public API on the build controller for saving a new build entity
// to persistent storage
func (bc *BuildController) CreateNew(newBuild CreateBuildRequest) (v1models.Build, error) {
	return bc.validateAndCreateNewBuild(newBuild)
}

// GetByID is the public api on the build controller for performing a lookup of
// a build entity by ID
func (bc *BuildController) GetByID(id int) (v1models.Build, error) {
	return bc.Store.GetByID(id)
}

// GetByVersionString will perform a look up of a build entity using it's unique version string
// ie image repo + tag
func (bc *BuildController) GetByVersionString(versionString string) (v1models.Build, error) {
	return bc.Store.GetByVersionString(versionString)
}

func (bc *BuildController) Serialize(builds ...v1models.Build) []v1serializers.BuildResponse {
	var buildsList []v1models.Build
	buildsList = append(buildsList, builds...)

	serializer := v1serializers.BuildsSerializer{Builds: buildsList}
	return serializer.Response()
}
