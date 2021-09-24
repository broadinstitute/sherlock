package deploys

// ServiceInstance is a datastructure representing an association
// between an environment and a service. They are an internal mechanism
// that is used to build the association between a build, service, and environment
// at a specific point in time which is needed to represent a deploy

import "gorm.io/gorm"

// ServiceInstanceController is the type used to manage logic related to working with
// ServiceInstance entities
type ServiceInstanceController struct {
	store serviceInstanceStore
}

// NewServiceInstanceController expects a gorm.DB connection and will provision
// a new controller instance
func NewServiceInstanceController(dbConn *gorm.DB) *ServiceInstanceController {
	store := newServiceInstanceStore(dbConn)
	return &ServiceInstanceController{
		store: store,
	}
}

// ListAll retrieves all service_instance entities from the backing data store
func (sic *ServiceInstanceController) ListAll() ([]ServiceInstance, error) {
	return sic.store.listAll()
}

// Serialize takes a variable number of service instance entities and serializes them into types suitable for use in
// client responses
func (sic *ServiceInstanceController) Serialize(serviceInstances ...ServiceInstance) []ServiceInstanceResponse {
	var serviceInstancesList []ServiceInstance
	serviceInstancesList = append(serviceInstancesList, serviceInstances...)

	serializer := ServiceInstancesSerializer{serviceInstancesList}
	return serializer.Response()
}
