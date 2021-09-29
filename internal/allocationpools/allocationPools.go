// AllocationPoolController, public interface for AllocationPool objects

// Package allocationPools defines data structure representing
// a allocationPool instance and methods for interacting with them
// it is left to concrete implementations in package db or others to
// implement these interfaces
package allocationpools

import (
	"errors"

	"gorm.io/gorm"
)

// AllocationPoolController is the management layer for allocationPools
type AllocationPoolController struct {
	store allocationPoolStore
}

// NewController accepts a gorm DB connection and returns a new instance
// of the allocationPool controller
func NewController(dbConn *gorm.DB) *AllocationPoolController {
	allocationPoolStore := newAllocationPoolStore(dbConn)
	return &AllocationPoolController{
		store: allocationPoolStore,
	}
}

// DoesAllocationPoolExist is a helper method to check if a allocationPool with the given name
// already exists in sherlock's data storage
func (allocationPoolController AllocationPoolController) DoesAllocationPoolExist(name string) (id int, ok bool) {
	allocationPool, err := allocationPoolController.GetByName(name)
	if errors.Is(err, ErrAllocationPoolNotFound) {
		return 0, false
	}
	return allocationPool.ID, true
}

// CreateNew is the public api on the allocationPoolController for persisting a new service entity to
// the data store
func (allocationPoolController *AllocationPoolController) CreateNew(newAllocationPool CreateAllocationPoolRequest) (AllocationPool, error) {
	return allocationPoolController.store.createNew(newAllocationPool)

}

// ListAll is the public api for listing out all allocationPools tracked by sherlock
func (allocationPoolController *AllocationPoolController) ListAll() ([]AllocationPool, error) {
	return allocationPoolController.store.listAll()

}

// GetByName is the public API for looking up a allocationPool from the data store by name
func (allocationPoolController *AllocationPoolController) GetByName(name string) (AllocationPool, error) {
	return allocationPoolController.store.getByName(name)

}
