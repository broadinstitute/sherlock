package v1serializers

// builds_serializers.go contains logic for building
// http responses from the builds data base model while avoiding
// dependencies on the database model in the route handling logic directly.
// This is an essentially an abstraction layer to give use more control over
// what is returned from api endpoints

import (
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"time"
)

// BuildsResponse is a type that allows all data returned from the /builds api group to share a consistent structure
type BuildsResponse struct {
	Builds []BuildResponse `json:"builds"`
	Error  string          `json:"error,omitempty"`
}

// BuildResponse is the type used to ensure that response bodies from the /builds api group have a consistent
// structure
type BuildResponse struct {
	ID            int             `json:"id"`
	VersionString string          `json:"version_string"`
	CommitSha     string          `json:"commit_sha"`
	BuildURL      string          `json:"build_url,omitempty"`
	BuiltAt       time.Time       `json:"built_at,omitempty"`
	Service       ServiceResponse `json:"service"`
}

// BuildSerializer takes a Build model entity and translates it into a response
type BuildSerializer struct {
	Build v1models.Build
}

// Response method performs the serialization from a Build entity to BuildsResponse
func (bs *BuildSerializer) Response() BuildResponse {
	service := ServiceSerializer{Service: bs.Build.Service}
	return BuildResponse{
		ID:            bs.Build.ID,
		VersionString: bs.Build.VersionString,
		CommitSha:     bs.Build.CommitSha,
		BuildURL:      bs.Build.BuildURL,
		BuiltAt:       bs.Build.BuiltAt,
		Service:       service.Response(),
	}
}

// BuildsSerializer is used to transform a slice of builds into the BuildsResponse type
type BuildsSerializer struct {
	Builds []v1models.Build
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
