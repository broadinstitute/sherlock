package services

import (
	"github.com/broadinstitute/sherlock/internal/models/v1models"
)

// builds_serializers.go contains logic for building
// http responses from the builds data base model while avoiding
// dependencies on the database model in the route handling logic directly.
// This is an essentially an abstraction layer to give use more control over
// what is returned from api endpoints

// ServiceResponse is the type used to serialize Service data that is returned
// to clients of the sherlock api
type ServiceResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	RepoURL string `json:"repo_url"`
}

// ServiceSerializer is used to serializer a single Service model
// to a used to generate responses from the /services api group
type ServiceSerializer struct {
	Service v1models.Service
}

// Response takes a Service Model entity and transforms it into a ServiceResponse
func (ss *ServiceSerializer) Response() ServiceResponse {
	return ServiceResponse{
		ID:      ss.Service.ID,
		Name:    ss.Service.Name,
		RepoURL: ss.Service.RepoURL,
	}
}

// ServicesSerializer is used to serialize a ServiceModel
// entity to a a format used in http response bodies
type ServicesSerializer struct {
	Services []v1models.Service
}

// Response serializes Service v1models into Service Responses
func (ss *ServicesSerializer) Response() []ServiceResponse {
	services := []ServiceResponse{}
	for _, service := range ss.Services {
		serializer := ServiceSerializer{service}
		services = append(services, serializer.Response())
	}
	return services
}

// Response is a type that allows all data returned from the /service api group to share a consistent structure
type Response struct {
	Services []ServiceResponse `json:"services"`
	Error    string            `json:"error,omitempty"`
}
