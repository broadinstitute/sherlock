package deploys

// ServiceInstance is a datastructure representing an association
// between an environment and a service. They are an internal mechanism
// that is used to build the association between a build, service, and environment
// at a specific point in time which is needed to represent a deploy

import (
	"errors"
	"fmt"

	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/services"
	"gorm.io/gorm"
)

var (
	errServiceDoesNotExist     = errors.New("unable to process Service instance create request. service does not exist")
	errEnvironmentDoesNotExist = errors.New("unable to process Service instance create request. environment does not exist")
)

// ServiceInstanceController is the type used to manage logic related to working with
// ServiceInstance entities
type ServiceInstanceController struct {
	store        serviceInstanceStore
	services     *services.ServiceController
	environments *environments.EnvironmentController
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

// CreateServiceInstanceRequest is the type used to validate the creation of a new service instance
// entity before it is persisted to the database.
type CreateServiceInstanceRequest struct {
	EnvironmentName string
	ServiceName     string
}

// CreateNew is used for persisting a new service instance to Sherlock's database. It performs some validation checks
// before actually executing the persistence  operation. It makes sure that both the referenced service and environment
// entities actually exist in the DB and will create them if not
func (sic *ServiceInstanceController) CreateNew(serviceName, environmentName string) (ServiceInstance, error) {
	createRequest := CreateServiceInstanceRequest{
		EnvironmentName: environmentName,
		ServiceName:     serviceName,
	}

	newServiceInstance, err := sic.validateCreateRequest(createRequest)
	if err != nil {
		return ServiceInstance{}, err
	}

	// if processing the create request succeeds then attempt to save it to the db
	savedServiceInstance, err := sic.store.createNew(newServiceInstance)
	if err != nil {
		return ServiceInstance{}, fmt.Errorf("error saving service instance: %v", err)
	}

	return savedServiceInstance, nil
}

func (sic *ServiceInstanceController) validateCreateRequest(newServiceInstance CreateServiceInstanceRequest) (ServiceInstance, error) {
	// make sure the service entity already exists and retrieve its id
	serviceID, ok := sic.services.DoesServiceExist(newServiceInstance.ServiceName)
	if !ok {
		return ServiceInstance{}, errServiceDoesNotExist
	}

	// make sure the environment entity already exists and retrieve its id
	environmentID, ok := sic.environments.DoesEnvironmentExist(newServiceInstance.EnvironmentName)
	if !ok {
		return ServiceInstance{}, errEnvironmentDoesNotExist
	}

	return ServiceInstance{
		ServiceID:     serviceID,
		EnvironmentID: environmentID,
	}, nil
}

// Serialize takes a variable number of service instance entities and serializes them into types suitable for use in
// client responses
func (sic *ServiceInstanceController) Serialize(serviceInstances ...ServiceInstance) []ServiceInstanceResponse {
	var serviceInstancesList []ServiceInstance
	serviceInstancesList = append(serviceInstancesList, serviceInstances...)

	serializer := ServiceInstancesSerializer{serviceInstancesList}
	return serializer.Response()
}
