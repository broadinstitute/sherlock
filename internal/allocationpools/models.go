// db-level managemenet for AllocationPool Struct
// APIs should not interact with this file and should user AllocationPoolController instead
// all gorm related methods should live in this file.

package allocationpools

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ErrAllocationPoolNotFound is the error to represent a failed lookup of a allocationPool db record
var ErrAllocationPoolNotFound = gorm.ErrRecordNotFound

// dataStore is a wrapper around a gorm postgres client
// which can be used to implement the allocationPoolRepository interface
type dataStore struct {
	*gorm.DB
}

// AllocationPool is the data structure that models a persisted to a database via gorm
type AllocationPool struct {
	ID        int    `gorm:"primaryKey;uniqueIndex"`
	Name      string `gorm:"not null;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// allocationPoolStore is the interface defining allowed db actions for AllocationPool
type allocationPoolStore interface {
	listAll() ([]AllocationPool, error)
	createNew(CreateAllocationPoolRequest) (AllocationPool, error)
	getByID(int) (AllocationPool, error)
	getByName(string) (AllocationPool, error)
}

// creates a db connection via gorm
func newAllocationPoolStore(dbconn *gorm.DB) dataStore {
	return dataStore{dbconn}
}

// CreateAllocationPoolRequest struct defines the data required to create a new allocationPool in db
type CreateAllocationPoolRequest struct {
	Name string `json:"name" binding:"required"`
}

// creates a allocationPool entity object to be persisted with the database from a
// request to create a allocationPool
func (createAllocationPoolRequest CreateAllocationPoolRequest) allocationPoolReq() AllocationPool {
	return AllocationPool{
		Name: createAllocationPoolRequest.Name,
	}
}

//
// db methods
//

// Returns ALL AllocationPools in the db
func (db dataStore) listAll() ([]AllocationPool, error) {
	allocationPools := []AllocationPool{}

	err := db.Find(&allocationPools).Error
	if err != nil {
		return []AllocationPool{}, fmt.Errorf("error retreiving allocationPools: %v", err)
	}

	return allocationPools, nil
}

// Saves an AllocationPool object to the db, returns the object if successful, nil otherwise
func (db dataStore) createNew(newAllocationPoolReq CreateAllocationPoolRequest) (AllocationPool, error) {
	newAllocationPool := newAllocationPoolReq.allocationPoolReq()

	if err := db.Create(&newAllocationPool).Error; err != nil {
		return newAllocationPool, fmt.Errorf("error saving to database: %v", err)
	}
	return newAllocationPool, nil
}

// Get is used to retrieve a specific allocationPool entity from a postgres database using
// id (primary key) as the lookup mechanism
func (db dataStore) getByID(id int) (AllocationPool, error) {
	allocationPool := AllocationPool{}

	if err := db.First(allocationPool, id).Error; err != nil {
		return allocationPool, err
	}
	return allocationPool, nil
}

// get an AllocationPool by name column
func (db dataStore) getByName(name string) (AllocationPool, error) {
	allocationPool := AllocationPool{}

	if err := db.Where(&AllocationPool{Name: name}).First(&allocationPool).Error; err != nil {
		return allocationPool, err
	}
	return allocationPool, nil
}
