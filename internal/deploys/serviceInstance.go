package deploys

// ServiceInstance is a datastructure representing an association
// between an environment and a service. They are an internal mechanism
// that is used to build the association between a build, service, and environment
// at a specific point in time which is needed to represent a deploy

import (
	"errors"

	"github.com/broadinstitute/sherlock/internal/controllers"
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
	clusters     *controllers.ClusterController
}

// CreateServiceInstanceRequest is a type containing the name of an environment and service
// that sherlock uses to create a new association between the two.
type CreateServiceInstanceRequest struct {
	EnvironmentName string
	ServiceName     string
	ClusterName     string
}

// NewServiceInstanceController expects a gorm.DB connection and will provision
// a new controller instance
func NewServiceInstanceController(dbConn *gorm.DB) *ServiceInstanceController {
	store := newServiceInstanceStore(dbConn)
	return &ServiceInstanceController{
		store:        store,
		services:     services.NewController(dbConn),
		environments: environments.NewController(dbConn),
		clusters:     controllers.NewClusterController(dbConn),
	}
}

// ListAll retrieves all service_instance entities from the backing data store
func (sic *ServiceInstanceController) ListAll() ([]ServiceInstance, error) {
	return sic.store.listAll()
}

// CreateNew accepts the name of an environment and a service. It will perform find or create operations for both
// and their create an association between them
func (sic *ServiceInstanceController) CreateNew(newServiceInstance CreateServiceInstanceRequest) (ServiceInstance, error) {
	// technically that can create orphaned objects as you could create an environment and then never attach it b/c something later failed.

	// check if the service already exists
	clusterID, err := sic.clusters.FindOrCreate(newServiceInstance.ClusterName)
	if err != nil {
		return ServiceInstance{}, err
	}

	// check if the environment already exists
	environmentID, err := sic.environments.FindOrCreate(newServiceInstance.EnvironmentName)
	if err != nil {
		return ServiceInstance{}, err
	}

	// check if the service already exists
	serviceID, err := sic.services.FindOrCreate(newServiceInstance.ServiceName)
	if err != nil {
		return ServiceInstance{}, err
	}

	return sic.store.createNew(clusterID, environmentID, serviceID)
}

// GetByEnvironmentAndServiceName accepts environment and service names as strings and will return the Service_Instance entity
// representing the association between them if it exists
func (sic *ServiceInstanceController) GetByEnvironmentAndServiceName(environmentName, serviceName string) (ServiceInstance, error) {
	// retrieve the service id if exists
	serviceID, exists := sic.services.DoesServiceExist(serviceName)
	if !exists {
		return ServiceInstance{}, ErrServiceInstanceNotFound
	}

	// retrieve environmentID if exists
	environmentID, exists := sic.environments.DoesEnvironmentExist(environmentName)
	if !exists {
		return ServiceInstance{}, ErrServiceInstanceNotFound
	}

	return sic.store.getByEnvironmentAndServiceID(environmentID, serviceID)
}

// FindOrCreate will check if a service instance with the given name and environment already exists
// if so it returns the id. If not it will create it and then return the id
func (sic *ServiceInstanceController) FindOrCreate(environmentName, serviceName string) (int, error) {
	// attempt to look up the serviceInstance
	serviceInstance, err := sic.GetByEnvironmentAndServiceName(environmentName, serviceName)

	// if it doesn't exist create
	if err != nil {
		if errors.Is(err, ErrServiceInstanceNotFound) {
			newServiceInstanceReq := CreateServiceInstanceRequest{
				EnvironmentName: environmentName,
				ServiceName:     serviceName,
			}

			serviceInstance, err = sic.CreateNew(newServiceInstanceReq)
			if err != nil {
				// got some error trying to create the service instance
				return 0, err
			}
		} else {
			// got some other error
			return 0, err
		}
	}
	return serviceInstance.ID, nil
}

// Serialize takes a variable number of service instance entities and serializes them into types suitable for use in
// client responses
func (sic *ServiceInstanceController) Serialize(serviceInstances ...ServiceInstance) []ServiceInstanceResponse {
	var serviceInstancesList []ServiceInstance
	serviceInstancesList = append(serviceInstancesList, serviceInstances...)

	serializer := ServiceInstancesSerializer{serviceInstancesList}
	return serializer.Response()
}
