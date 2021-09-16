// EnvironmentController, public interface for Environment objects

// Package environments defines data structure representing
// a environment instance and methods for interacting with them
// it is left to concrete implementations in package db or others to
// implement these interfaces
package environments

import (
	"errors"

	"gorm.io/gorm"
)

// EnvironmentController is the management layer for environments
type EnvironmentController struct {
	store environmentStore
	// store jobRequestStore // TODO: placeholder for other store environment controller can touch
}

// NewController accepts a gorm DB connection and returns a new instance
// of the environment controller
func NewController(dbConn *gorm.DB) *EnvironmentController {
	environmentStore := newEnvironmentStore(dbConn)
	return &EnvironmentController{
		store: environmentStore,
	}
}

// DoesEnvironmentExist is a helper method to check if a environment with the given name
// already exists in sherlock's data storage
func (environmentController EnvironmentController) DoesEnvironmentExist(name string) (id int, ok bool) {
	environment, err := environmentController.GetByName(name)
	if errors.Is(err, ErrEnvironmentNotFound) {
		return 0, false
	}
	return environment.ID, true
}

// CreateNew is the public api on the environmentController for persisting a new service entity to
// the data store
func (environmentController *EnvironmentController) CreateNew(newEnvironment CreateEnvironmentRequest) (Environment, error) {
	return environmentController.store.createNew(newEnvironment)
}

// ListAll is the public api for listing out all environments tracked by sherlock
func (environmentController *EnvironmentController) ListAll() ([]Environment, error) {
	return environmentController.store.listAll()
}

// GetByName is the public API for looking up a environment from the data store by name
func (environmentController *EnvironmentController) GetByName(name string) (Environment, error) {
	return environmentController.store.getByName(name)
}

// Response is a type that allows all data returned from the /environment api group to share a consistent structure
type Response struct {
	Environments []Environment `json:"environments"`
	Error        string        `json:"error,omitempty"`
}
