// Package services defines the control plane for sherlock's
// service entities and api routes for interating with those control
// plane methods
package services

// builds.go contains the "business" logic for operations relating to service entities.
// This could eventually be moved to it's own sub-folder if it becomes unwieldy

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// ServiceController is the management layer for CRUD operations for service entities
type ServiceController struct {
	store serviceStore
}

// NewController accepts a gorm DB connection and returns a new instance
// of the service controller
func NewController(dbConn *gorm.DB) *ServiceController {
	serviceStore := newServiceStore(dbConn)
	return &ServiceController{
		store: serviceStore,
	}
}

// DoesServiceExist is a helper method to check if a service with the given name
// already exists in sherlock's data storage
func (sc *ServiceController) DoesServiceExist(name string) (id int, ok bool) {
	svc, err := sc.GetByName(name)
	if errors.Is(err, ErrServiceNotFound) {
		return 0, false
	}
	return svc.ID, true
}

// CreateNew is the public api on the serviceController for persisting a new service entity to
// the data store
func (sc *ServiceController) CreateNew(newService CreateServiceRequest) (Service, error) {
	return sc.store.createNew(newService)
}

// ListAll is the public api for listing out all services tracked by sherlock
func (sc *ServiceController) ListAll() ([]Service, error) {
	return sc.store.listAll()
}

// GetByName is the public API for looking up a service from the data store by name
func (sc *ServiceController) GetByName(name string) (Service, error) {
	return sc.store.getByName(name)
}

// FindOrCreate will attempt to look an environment by name and return its ID if successful
// if unsuccessful it will create a new environment from the provider name and return that id
func (sc *ServiceController) FindOrCreate(name string) (int, error) {
	serviceID, exists := sc.DoesServiceExist(name)

	if !exists {
		// then make the new service
		newService := CreateServiceRequest{Name: name}
		createdService, err := sc.CreateNew(newService)
		if err != nil {
			return 0, fmt.Errorf("error creating service %s: %v", name, err)
		}
		serviceID = createdService.ID
	}
	return serviceID, nil
}

func (sc *ServiceController) serialize(services ...Service) []ServiceResponse {
	var serviceList []Service
	serviceList = append(serviceList, services...)

	serializer := ServicesSerializer{Services: serviceList}
	return serializer.Response()
}

// CreateServiceRequest is a type used to represent the information required to register a new service in sherlock
type CreateServiceRequest struct {
	Name    string `json:"name" binding:"required"`
	RepoURL string `json:"repo_url" binding:"required"`
}

// creates a service entity object to be persisted with the database from a
// request to create a service
func (cr *CreateServiceRequest) service() Service {
	return Service{
		Name:    cr.Name,
		RepoURL: cr.RepoURL,
	}
}
