// db-level managemenet for Environment Struct
// APIs should not interact with this Environment and should instead use EnvironmentController, thus all methods should be private
// all gorm related method should live in this file.

package environments

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	// "encoding/json" // TODO: deal with the jsonb
)

// ErrEnvironmentNotFound is the error to represent a failed lookup of a environment db record
var ErrEnvironmentNotFound = gorm.ErrRecordNotFound

// dataStore is a wrapper around a gorm postgres client
// which can be used to implement the environmentStore interface
type dataStore struct {
	*gorm.DB
}

// Environment is the data structure that models a persisted to a database via gorm
type Environment struct {
	ID               int    `gorm:"primaryKey"`
	Name             string `gorm:"not null;default:null"`
	IsPermanent      bool
	Requester        string
	DestroyedAt      time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	AllocationPoolID *int `gorm:"default:null"`
}

// environmentStore is the interface defining allowed db actions for Environment
type environmentStore interface {
	listAll() ([]Environment, error)
	createNew(CreateEnvironmentRequest) (Environment, error)
	getByID(int) (Environment, error)
	getByName(string) (Environment, error)
}

// creates a db connection via gorm
func newEnvironmentStore(dbconn *gorm.DB) dataStore {
	return dataStore{dbconn}
}

// CreateEnvironmentRequest struct defines the data required to create a new environment in db
type CreateEnvironmentRequest struct {
	Name string `json:"name" binding:"required"`
}

// creates a environment entity object to be persisted with the database from a
// request to create a environment
func (createEnvironmentRequest CreateEnvironmentRequest) EnvironmentReq() Environment {
	return Environment{
		Name: createEnvironmentRequest.Name,
	}
}

//
// db methods
//

// Returns ALL Environments in the db
func (db dataStore) listAll() ([]Environment, error) {
	environments := []Environment{}

	err := db.Find(&environments).Error
	if err != nil {
		return []Environment{}, fmt.Errorf("error retreiving environments: %v", err)
	}

	return environments, nil
}

// Saves an Environment object to the db, returns the object if successful, nil otherwise
func (db dataStore) createNew(newEnvironmentReq CreateEnvironmentRequest) (Environment, error) {
	newEnvironment := newEnvironmentReq.EnvironmentReq()

	if err := db.Create(&newEnvironment).Error; err != nil {
		return Environment{}, fmt.Errorf("error saving to database: %v", err)
	}
	return newEnvironment, nil
}

// Get is used to retrieve a specific environment entity from a postgres database using
// id (primary key) as the lookup mechanism
func (db dataStore) getByID(id int) (Environment, error) {
	environment := Environment{}

	if err := db.First(&environment, id).Error; err != nil {
		return Environment{}, err
	}
	return environment, nil
}

// get an Environment by name column
func (db dataStore) getByName(name string) (Environment, error) {
	environment := Environment{}

	if err := db.Where(&Environment{Name: name}).First(&environment).Error; err != nil {
		return Environment{}, err
	}
	return environment, nil
}
