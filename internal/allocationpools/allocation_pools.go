// AllocationPoolController, public interface for AllocationPool objects

// Package allocationPools defines data structure representing
// a allocationPool instance and methods for interacting with them
// it is left to concrete implementations in package db or others to
// implement these interfaces
package allocationpools

import (
	"errors"

	"github.com/broadinstitute/sherlock/internal/models"
	"gorm.io/gorm"
)

// AllocationPoolController is the management layer for allocationPools
type AllocationPoolController struct {
	store models.AllocationPoolStore
}

// NewController accepts a gorm DB connection and returns a new instance
// of the allocationPool controller
func NewController(dbConn *gorm.DB) *AllocationPoolController {
	allocationPoolStore := models.NewAllocationPoolStore(dbConn)
	return &AllocationPoolController{
		store: allocationPoolStore,
	}
}

// DoesAllocationPoolExist is a helper method to check if a allocationPool with the given name
// already exists in sherlock's data storage
func (allocationPoolController AllocationPoolController) DoesAllocationPoolExist(name string) (int, bool) {
	allocationPool, err := allocationPoolController.GetByName(name)
	if errors.Is(err, models.ErrAllocationPoolNotFound) {
		return 0, false
	}
	return allocationPool.ID, true
}

// CreateNew is the public api on the allocationPoolController for persisting a new service entity to
// the data store
func (allocationPoolController *AllocationPoolController) CreateNew(newAllocationPool models.CreateAllocationPoolRequest) (models.AllocationPool, error) {
	return allocationPoolController.store.CreateNew(newAllocationPool)

}

// ListAll is the public api for listing out all allocationPools tracked by sherlock
func (allocationPoolController *AllocationPoolController) ListAll() ([]models.AllocationPool, error) {
	return allocationPoolController.store.ListAll()

}

// GetByName is the public API for looking up a allocationPool from the data store by name
func (allocationPoolController *AllocationPoolController) GetByName(name string) (models.AllocationPool, error) {
	return allocationPoolController.store.GetByName(name)
}

// GetByID is the public API for looking up a allocationPool from the data store by name
func (allocationPoolController *AllocationPoolController) GetByID(id int) (models.AllocationPool, error) {
	return allocationPoolController.store.GetByID(id)
}

// AddEnvironmentByID takes a AllocationPoolObject and associates an existing environment to it.
func (allocationPoolController *AllocationPoolController) AddEnvironmentByID(currentAllocationPool models.AllocationPool, environmentID int) (models.AllocationPool, error) {
	return allocationPoolController.store.AddEnvironmentByID(currentAllocationPool, environmentID)
}
