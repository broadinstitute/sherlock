package deploys

import (
	"time"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/services"
)

// Response is a type that allows all data returned from the /builds api group to share a consistent structure
type Response struct {
	Deploys []DeployResponse `json:"deploys"`
	Error   string           `json:"error,omitempty"`
}

// ServiceInstanceResponse is the type that is used to represent data
// about a ServiceInstance entity in response to clients. Its purpose
// is to decouple responses from the database model
type ServiceInstanceResponse struct {
	ID          int                              `json:"id"`
	Service     services.ServiceResponse         `json:"service"`
	Environment environments.EnvironmentResponse `json:"environment"`
}

// ServiceInstanceSerializer is an intermediate type used to
// convert a Service instance into its response type
type ServiceInstanceSerializer struct {
	serviceInstance ServiceInstance
}

// Response consumes a ServiceInstanceSerializer and generated a response type
func (sis *ServiceInstanceSerializer) Response() ServiceInstanceResponse {
	service := services.ServiceSerializer{Service: sis.serviceInstance.Service}
	environment := environments.EnvironmentSerializer{Environment: sis.serviceInstance.Environment}

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
	ServiceInstances []ServiceInstance
}

// Response Will generate a slice of Service Instance Response from
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
	Build           builds.BuildResponse    `json:"build"`
	CreatedAt       time.Time               `json:"deployed_at"`
}

type deploySerializer struct {
	deploy Deploy
}

func (ds *deploySerializer) Response() DeployResponse {
	serviceInstance := ServiceInstanceSerializer{serviceInstance: ds.deploy.ServiceInstance}
	build := builds.BuildSerializer{Build: ds.deploy.Build}
	return DeployResponse{
		ID:              ds.deploy.ID,
		ServiceInstance: serviceInstance.Response(),
		Build:           build.Response(),
		CreatedAt:       ds.deploy.CreatedAt,
	}
}

// DeploysSerializer is used to transform a slice of Deploy models into
// into deploy responses and can supply additional data to attach to the response
type DeploysSerializer struct {
	deploys []Deploy
}

// Response is used to seralize a slice of deploy database models
// into a slice of deploy api responses
func (ds *DeploysSerializer) Response() []DeployResponse {
	deploys := []DeployResponse{}
	for _, Deploy := range ds.deploys {
		serializer := deploySerializer{Deploy}
		deploys = append(deploys, serializer.Response())
	}
	return deploys
}
