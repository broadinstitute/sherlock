package deploys

import (
	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/services"
)

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
