package v1serializers

import (
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"time"

	"github.com/broadinstitute/sherlock/internal/services"
)

// DeploysResponse is a type that allows all data returned from the /builds api group to share a consistent structure
type DeploysResponse struct {
	Deploys []DeployResponse `json:"deploys"`
	Error   string           `json:"error,omitempty"`
}

// ServiceInstanceResponse is the type that is used to represent data
// about a ServiceInstance entity in response to clients. Its purpose
// is to decouple responses from the database model
type ServiceInstanceResponse struct {
	ID          int                      `json:"id"`
	Service     services.ServiceResponse `json:"service"`
	Environment EnvironmentResponse      `json:"environment"`
}

// ServiceInstanceSerializer is an intermediate type used to
// convert a Service instance into its response type
type ServiceInstanceSerializer struct {
	serviceInstance v1models.ServiceInstance
}

// DeploysResponse consumes a ServiceInstanceSerializer and generated a response type
func (sis *ServiceInstanceSerializer) Response() ServiceInstanceResponse {
	service := services.ServiceSerializer{Service: sis.serviceInstance.Service}
	environment := EnvironmentSerializer{Environment: sis.serviceInstance.Environment}

	return ServiceInstanceResponse{
		ID:          sis.serviceInstance.ID,
		Service:     service.Response(),
		Environment: environment.Response(),
	}
}

// ServiceInstancesSerializer is a wrapper around
// ServiceInstanceSerializer that supports serialization of
// mulitple ServiceInstance entities
type ServiceInstancesSerializer struct {
	ServiceInstances []v1models.ServiceInstance
}

// DeploysResponse Will generate a slice of Service Instance DeploysResponse from
// ServiceInstancesSerializer
func (sis *ServiceInstancesSerializer) Response() []ServiceInstanceResponse {
	serviceInstances := []ServiceInstanceResponse{}
	for _, serviceInstance := range sis.ServiceInstances {
		serializer := ServiceInstanceSerializer{serviceInstance}
		serviceInstances = append(serviceInstances, serializer.Response())
	}
	return serviceInstances
}

// DeployResponse is the type used for generating api responses
// containing information about deploy(s)
type DeployResponse struct {
	ID              int                     `json:"id"`
	ServiceInstance ServiceInstanceResponse `json:"service_instance"`
	Build           BuildResponse           `json:"build"`
	CreatedAt       time.Time               `json:"deployed_at"`
}

type deploySerializer struct {
	deploy v1models.Deploy
}

func (ds *deploySerializer) Response() DeployResponse {
	serviceInstance := ServiceInstanceSerializer{serviceInstance: ds.deploy.ServiceInstance}
	build := BuildSerializer{Build: ds.deploy.Build}
	return DeployResponse{
		ID:              ds.deploy.ID,
		ServiceInstance: serviceInstance.Response(),
		Build:           build.Response(),
		CreatedAt:       ds.deploy.CreatedAt,
	}
}

// DeploysSerializer is used to transform a slice of Deploy v1models into
// into deploy responses and can supply additional data to attach to the response
type DeploysSerializer struct {
	Deploys []v1models.Deploy
}

// DeploysResponse is used to seralize a slice of deploy database v1models
// into a slice of deploy api responses
func (ds *DeploysSerializer) Response() []DeployResponse {
	deploys := []DeployResponse{}
	for _, Deploy := range ds.Deploys {
		serializer := deploySerializer{Deploy}
		deploys = append(deploys, serializer.Response())
	}
	return deploys
}
