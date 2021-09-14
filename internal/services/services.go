// Package services defines data structure representing
// a service instance and methods for interacting with them
// it is left to concrete implementations in package db or others to
// implement these interfaces
package services

import (
	"errors"

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
func (sc *ServiceController) CreateNew(newService CreateServiceRequest) (*Service, error) {
	return sc.store.createNew(newService)
}

// ListAll is the public api for listing out all services tracked by sherlock
func (sc *ServiceController) ListAll() ([]*Service, error) {
	return sc.store.listAll()
}

// GetByName is the public API for looking up a service from the data store by name
func (sc *ServiceController) GetByName(name string) (*Service, error) {
	return sc.store.getByName(name)
}

// CreateServiceRequest is a type used to represent the information required to register a new service in sherlock
type CreateServiceRequest struct {
	Name    string `json:"name" binding:"required"`
	RepoURL string `json:"repo_url" binding:"required"`
}

// creates a service entity object to be persisted with the database from a
// request to create a service
func (cr *CreateServiceRequest) service() *Service {
	return &Service{
		Name:    cr.Name,
		RepoURL: cr.RepoURL,
	}
}

// Response is a type that allows all data returned from the /service api group to share a consistent structure
type Response struct {
	Services []*Service `json:"services"`
	Error    string     `json:"error,omitempty"`
}
