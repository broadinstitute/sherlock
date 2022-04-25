// Package services defines the control plane for sherlock's
// service entities and api routes for interating with those control
// plane methods
package services

// builds.go contains the "business" logic for operations relating to service entities.
// This could eventually be moved to it's own sub-folder if it becomes unwieldy

import (
	"errors"
	"fmt"

	"github.com/broadinstitute/sherlock/internal/models"
	"gorm.io/gorm"
)

// ServiceController is the management layer for CRUD operations for service entities
type ServiceController struct {
	store models.ServiceStore
}

// NewController accepts a gorm DB connection and returns a new instance
// of the service controller
func NewController(dbConn *gorm.DB) *ServiceController {
	serviceStore := models.NewServiceStore(dbConn)
	return &ServiceController{
		store: serviceStore,
	}
}

// DoesServiceExist is a helper method to check if a service with the given name
// already exists in sherlock's data storage
func (sc *ServiceController) DoesServiceExist(name string) (id int, ok bool) {
	svc, err := sc.GetByName(name)
	if errors.Is(err, models.ErrServiceNotFound) {
		return 0, false
	}
	return svc.ID, true
}

// CreateNew is the public api on the serviceController for persisting a new service entity to
// the data store
func (sc *ServiceController) CreateNew(newService models.CreateServiceRequest) (models.Service, error) {
	return sc.store.CreateNew(newService)
}

// ListAll is the public api for listing out all services tracked by sherlock
func (sc *ServiceController) ListAll() ([]models.Service, error) {
	return sc.store.ListAll()
}

// GetByName is the public API for looking up a service from the data store by name
func (sc *ServiceController) GetByName(name string) (models.Service, error) {
	return sc.store.GetByName(name)
}

// FindOrCreate will attempt to look an environment by name and return its ID if successful
// if unsuccessful it will create a new environment from the provider name and return that id
func (sc *ServiceController) FindOrCreate(name string) (int, error) {
	serviceID, exists := sc.DoesServiceExist(name)

	if !exists {
		// then make the new service
		newService := models.CreateServiceRequest{Name: name}
		createdService, err := sc.CreateNew(newService)
		if err != nil {
			return 0, fmt.Errorf("error creating service %s: %v", name, err)
		}
		serviceID = createdService.ID
	}
	return serviceID, nil
}

func (sc *ServiceController) serialize(services ...models.Service) []ServiceResponse {
	var serviceList []models.Service
	serviceList = append(serviceList, services...)

	serializer := ServicesSerializer{Services: serviceList}
	return serializer.Response()
}
