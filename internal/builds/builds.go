package builds

import (
	"time"

	"gorm.io/gorm"
)

// BuildController is the management layer that processes requests
// to the /builds api group
type BuildController struct {
	store buildStore
}

// NewController returns an instance of the controller struct for
// interacting with build entities. It embeds a buildStore interface for
// operations on the build persistence layer
func NewController(dbConn *gorm.DB) *BuildController {
	buildStore := newBuildStore(dbConn)
	return &BuildController{
		store: buildStore,
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
	BuildURL      string    `json:"build_url"`
	BuiltAt       time.Time `json:"built_at"`
	ServiceName   string    `json:"service_name" binding:"required"`
}
