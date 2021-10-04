// EnvironmentController, public interface for Environment objects

// Package environments defines data structure representing
// a environment instance and methods for interacting with them
// it is left to concrete implementations in package db or others to
// implement these interfaces
package environments

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// EnvironmentController is the management layer for environments
type EnvironmentController struct {
	store environmentStore
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

// FindOrCreate will attempt to look an environment by name and return its ID if successful
// if unsuccessful it will create a new environment from the provider name and return that id
func (environmentController *EnvironmentController) FindOrCreate(name string) (int, error) {
	environmentID, exists := environmentController.DoesEnvironmentExist(name)

	if !exists {
		newEnvironment := CreateEnvironmentRequest{Name: name}
		createdEnvironment, err := environmentController.CreateNew(newEnvironment)
		if err != nil {
			return 0, fmt.Errorf("error creating environment %s: %v", name, err)
		}
		environmentID = createdEnvironment.ID
	}
	return environmentID, nil
}

// Takes an GORM Environment object and returns a JSON for environment
func (environmentController *EnvironmentController) serialize(environments ...Environment) []EnvironmentResponse {
	// collect arguments into a slice to be serialized into a single response
	var environmentList []Environment
	environmentList = append(environmentList, environments...)

	serializer := EnvironmentsSerializer{Environments: environmentList}
	return serializer.Response()
}

// Response is a type that allows all data returned from the /environment api group to share a consistent structure
type Response struct {
	Environments []EnvironmentResponse `json:"environments"`
	Error        string                `json:"error,omitempty"`
}
