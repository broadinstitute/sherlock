// EnvironmentController, public interface for Environment objects

// Package environments defines data structure representing
// a environment instance and methods for interacting with them
// it is left to concrete implementations in package db or others to
// implement these interfaces
package v1controllers

import (
	"errors"
	"fmt"
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"

	"gorm.io/gorm"
)

// EnvironmentController is the management layer for environments
type EnvironmentController struct {
	store v1models.EnvironmentStore
}

// NewEnvironmentController accepts a gorm DB connection and returns a new instance
// of the environment controller
func NewEnvironmentController(dbConn *gorm.DB) *EnvironmentController {
	environmentStore := v1models.NewEnvironmentStore(dbConn)
	return &EnvironmentController{
		store: environmentStore,
	}
}

// DoesEnvironmentExist is a helper method to check if a environment with the given name
// already exists in sherlock's data storage
func (environmentController EnvironmentController) DoesEnvironmentExist(name string) (id int, ok bool) {
	environment, err := environmentController.GetByName(name)
	if errors.Is(err, v1models.ErrEnvironmentNotFound) {
		return 0, false
	}
	return environment.ID, true
}

// CreateNew is the public api on the environmentController for persisting a new service entity to
// the data store
func (environmentController *EnvironmentController) CreateNew(newEnvironment v1models.CreateEnvironmentRequest) (v1models.Environment, error) {
	return environmentController.store.CreateNew(newEnvironment)

}

// ListAll is the public api for listing out all environments tracked by sherlock
func (environmentController *EnvironmentController) ListAll() ([]v1models.Environment, error) {
	return environmentController.store.ListAll()

}

// GetByName is the public API for looking up a environment from the data store by name
func (environmentController *EnvironmentController) GetByName(name string) (v1models.Environment, error) {
	return environmentController.store.GetByName(name)
}

// GetByName is the public API for looking up a environment from the data store by name
func (environmentController *EnvironmentController) GetByID(ID int) (v1models.Environment, error) {
	return environmentController.store.GetByID(ID)
}

// FindOrCreate will attempt to look an environment by name and return its ID if successful
// if unsuccessful it will create a new environment from the provider name and return that id
func (environmentController *EnvironmentController) FindOrCreate(name string) (int, error) {
	environmentID, exists := environmentController.DoesEnvironmentExist(name)

	if !exists {
		newEnvironment := v1models.CreateEnvironmentRequest{Name: name}
		createdEnvironment, err := environmentController.CreateNew(newEnvironment)
		if err != nil {
			return 0, fmt.Errorf("error creating environment %s: %v", name, err)
		}
		environmentID = createdEnvironment.ID
	}
	return environmentID, nil
}

// Takes an GORM Environment object and returns a JSON for environment
func (environmentController *EnvironmentController) Serialize(environments ...v1models.Environment) []v1serializers.EnvironmentResponse {
	// collect arguments into a slice to be serialized into a single response
	var environmentList []v1models.Environment
	environmentList = append(environmentList, environments...)

	serializer := v1serializers.EnvironmentsSerializer{Environments: environmentList}
	return serializer.Response()
}

// Response is a type that allows all data returned from the /environment api group to share a consistent structure
type Response struct {
	Environments []v1serializers.EnvironmentResponse `json:"environments"`
	Error        string                              `json:"error,omitempty"`
}
