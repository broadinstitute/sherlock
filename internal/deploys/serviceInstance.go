package deploys

// ServiceInstance is a datastructure representing an association
// between an environment and a service. They are an internal mechanism
// that is used to build the association between a build, service, and environment
// at a specific point in time which is needed to represent a deploy

import (
	"fmt"

	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/services"
	"gorm.io/gorm"
)

// ServiceInstanceController is the type used to manage logic related to working with
// ServiceInstance entities
type ServiceInstanceController struct {
	store        serviceInstanceStore
	services     *services.ServiceController
	environments *environments.EnvironmentController
}

// CreateServiceInstanceRequest is a type containing the name of an environment and service
// that sherlock uses to create a new association between the two.
type CreateServiceInstanceRequest struct {
	EnvironmentName string
	ServiceName     string
}

// NewServiceInstanceController expects a gorm.DB connection and will provision
// a new controller instance
func NewServiceInstanceController(dbConn *gorm.DB) *ServiceInstanceController {
	store := newServiceInstanceStore(dbConn)
	return &ServiceInstanceController{
		store:        store,
		services:     services.NewController(dbConn),
		environments: environments.NewController(dbConn),
	}
}

// ListAll retrieves all service_instance entities from the backing data store
func (sic *ServiceInstanceController) ListAll() ([]ServiceInstance, error) {
	return sic.store.listAll()
}

// CreateNew accepts the name of an environment and a service. It will perform find or create operations for both
// and their create an association between them
func (sic *ServiceInstanceController) CreateNew(newServiceInstance CreateServiceInstanceRequest) (ServiceInstance, error) {
	// check if the environment already exists
	environmentID, err := sic.environments.FindOrCreate(newServiceInstance.EnvironmentName)
	if err != nil {
		return ServiceInstance{}, err
	}

	// check if the service already exists
	serviceID, doesExist := sic.services.DoesServiceExist(newServiceInstance.ServiceName)
	if !doesExist {
		return ServiceInstance{}, fmt.Errorf("service: %s does not exist", newServiceInstance.ServiceName)
	}

	return sic.store.createNew(environmentID, serviceID)
}

// GetByEnvironmentAndServiceName accepts environment and service names as strings and will return the Service_Instance entity
// representing the association between them if it exists
func (sic *ServiceInstanceController) GetByEnvironmentAndServiceName(environmentName, serviceName string) (ServiceInstance, error) {
	return sic.store.getByEnvironmentAndServiceName(environmentName, serviceName)
}

// Serialize takes a variable number of service instance entities and serializes them into types suitable for use in
// client responses
func (sic *ServiceInstanceController) Serialize(serviceInstances ...ServiceInstance) []ServiceInstanceResponse {
	var serviceInstancesList []ServiceInstance
	serviceInstancesList = append(serviceInstancesList, serviceInstances...)

	serializer := ServiceInstancesSerializer{serviceInstancesList}
	return serializer.Response()
}
