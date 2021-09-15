package builds

import (
	"time"

	"github.com/broadinstitute/sherlock/internal/services"
)

// Response is a type that allows all data returned from the /builds api group to share a consistent structure
type Response struct {
	Builds []BuildResponse `json:"builds"`
	Error  string          `json:"error,omitempty"`
}

// BuildResponse is the type used to ensure that response bodies from the /builds api group have a consistent
// structure
type BuildResponse struct {
	ID            int                      `json:"id"`
	VersionString string                   `json:"version_string"`
	CommitSha     string                   `json:"commit_sha"`
	BuildURL      string                   `json:"build_url,omitempty"`
	BuiltAt       time.Time                `json:"built_at,omitempty"`
	Service       services.ServiceResponse `json:"service"`
}

// BuildSerializer takes a Build model entity and translates it into a response
type BuildSerializer struct {
	Build
}

// Response method performs the serialization from a Build entity to Build Response
func (bs *BuildSerializer) Response() BuildResponse {
	service := services.ServiceSerializer{bs.Service}
	return BuildResponse{
		ID:            bs.ID,
		VersionString: bs.VersionString,
		CommitSha:     bs.CommitSha,
		BuildURL:      bs.BuildURL,
		BuiltAt:       bs.BuiltAt,
		Service:       service.Response(),
	}
}

// BuildsSerializer is used to transform a slice of builds into the Response type
type BuildsSerializer struct {
	Builds []Build
}

// Response transforms a list of build model entities into BuildResponse's
func (bs *BuildsSerializer) Response() []BuildResponse {
	builds := []BuildResponse{}
	for _, build := range bs.Builds {
		serializer := BuildSerializer{build}
		builds = append(builds, serializer.Response())
	}
	return builds
}
