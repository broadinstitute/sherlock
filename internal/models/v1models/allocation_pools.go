// db-level managemenet for AllocationPool Struct
// APIs should not interact with this file and should user AllocationPoolController instead
// all gorm related methods should live in this file.

package v1models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ErrAllocationPoolNotFound is the error to represent a failed lookup of a allocationPool db record
var ErrAllocationPoolNotFound = gorm.ErrRecordNotFound

// allocationPoolStore is a wrapper around a gorm postgres client
// which can be used to implement the allocationPoolRepository interface
type allocationPoolStore struct {
	*gorm.DB
}

// AllocationPool is the data structure that v1models a persisted to a database via gorm
type AllocationPool struct {
	ID           int    `gorm:"primaryKey;uniqueIndex"`
	Name         string `gorm:"not null;default:null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Environments []Environment
}

// allocationPoolStore is the interface defining allowed db actions for AllocationPool
type AllocationPoolStore interface {
	ListAll() ([]AllocationPool, error)
	CreateNew(CreateAllocationPoolRequest) (AllocationPool, error)
	GetByID(int) (AllocationPool, error)
	GetByName(string) (AllocationPool, error)
	AddEnvironmentByID(AllocationPool, int) (AllocationPool, error)
}

// NewAllocationPoolStore creates a db connection via gorm
func NewAllocationPoolStore(dbconn *gorm.DB) allocationPoolStore {
	return allocationPoolStore{dbconn}
}

// CreateAllocationPoolRequest struct defines the data required to create a new allocationPool in db
type CreateAllocationPoolRequest struct {
	Name         string `json:"name" binding:"required"`
	Environments []Environment
}

// creates a allocationPool entity object to be persisted with the database from a
// request to create a allocationPool
func (createAllocationPoolRequest CreateAllocationPoolRequest) AllocationPoolReq() AllocationPool {
	return AllocationPool{
		Name: createAllocationPoolRequest.Name,
	}
}

//
// db methods
//

// Returns ALL AllocationPools in the db
func (db allocationPoolStore) ListAll() ([]AllocationPool, error) {
	allocationPools := []AllocationPool{}

	err := db.Find(&allocationPools).Error
	if err != nil {
		return []AllocationPool{}, fmt.Errorf("error retreiving allocationPools: %v", err)
	}

	return allocationPools, nil
}

// Saves an AllocationPool object to the db, returns the object if successful, nil otherwise
func (db allocationPoolStore) CreateNew(newAllocationPoolReq CreateAllocationPoolRequest) (AllocationPool, error) {
	newAllocationPool := newAllocationPoolReq.AllocationPoolReq()

	if err := db.Create(&newAllocationPool).Error; err != nil {
		return AllocationPool{}, fmt.Errorf("error saving to database: %v", err)
	}
	return newAllocationPool, nil
}

// Get is used to retrieve a specific allocationPool entity from a postgres database using
// id (primary key) as the lookup mechanism
func (db allocationPoolStore) GetByID(id int) (AllocationPool, error) {
	allocationPool := AllocationPool{}

	if err := db.First(&allocationPool, id).Error; err != nil {
		return AllocationPool{}, err
	}
	return allocationPool, nil
}

// get an AllocationPool by name column
func (db allocationPoolStore) GetByName(name string) (AllocationPool, error) {
	allocationPool := AllocationPool{}

	if err := db.Where(&AllocationPool{Name: name}).First(&allocationPool).Error; err != nil {
		return AllocationPool{}, err
	}
	return allocationPool, nil
}

// Take an existing environment and add it to the allocationPool.
func (db allocationPoolStore) AddEnvironmentByID(allocationPool AllocationPool, environmentID int) (AllocationPool, error) {
	environment := Environment{}

	//get the existing environment to add
	if err := db.Where("id = ?", environmentID).First(&environment).Error; err != nil {
		return AllocationPool{}, err
	}

	if err := db.Model(&allocationPool).Association("Environments").Append(&environment); err != nil {
		return AllocationPool{}, err
	}

	return allocationPool, nil
}
