// Package services defines data structure representing
// a service instance and methods for interacting with them
// it is left to concrete implementations in package db or others to
// implement these interfaces
package services

import "gorm.io/gorm"

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

// CreateRequest is a type used to represent the information required to register a new service in sherlock
type CreateRequest struct {
	Name    string `json:"name" binding:"required"`
	RepoURL string `json:"repo_url" binding:"required"`
}

// creates a service entity object to be persisted with the database from a
// request to create a service
func (cr *CreateRequest) service() *Service {
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
