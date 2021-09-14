package services

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ErrServiceNotFound is the error to represent a failed lookup of a service entity
var ErrServiceNotFound = gorm.ErrRecordNotFound

// dataStore is a wrapper around a gorm postgres client
// which can be used to implement the serviceRepository interface
type dataStore struct {
	*gorm.DB
}

// Service is the data structure that models a service entity persisted to a dataabase via gorm
type Service struct {
	ID        int
	Name      string
	RepoURL   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// serviceStore is the interface type that defines the methods required for implementing the persistence layer
// for services entities
type serviceStore interface {
	ListAll() ([]*Service, error)
	createNew(CreateServiceRequest) (*Service, error)
	GetByName(string) (*Service, error)
}

func newServiceStore(db *gorm.DB) dataStore {
	return dataStore{db}
}

// ListAll retrieves all service entities from a postgres database and
// returns them as a slice
func (db dataStore) ListAll() ([]*Service, error) {
	services := []*Service{}

	err := db.Find(&services).Error
	if err != nil {
		return []*Service{}, fmt.Errorf("Error retriving services: %v", err)
	}

	return services, nil
}

// Create will persist the service defined by newservice to a postgres database.
// It will return the service as stored in postgres for ease of testing if successful
func (db dataStore) createNew(newServiceReq CreateServiceRequest) (*Service, error) {
	newService := newServiceReq.service()
	if err := db.Create(newService).Error; err != nil {
		return nil, fmt.Errorf("error saving service to database: %v", err)
	}
	return newService, nil
}

// getByName retrives a service entity from persistence layer. It returns an ErrRecordNotFound
// if the requested name does not exist. This effectively equivalent to get service by id
// as the name field is indexed and enforced to be unique
func (db dataStore) GetByName(name string) (*Service, error) {
	service := &Service{}

	if err := db.Where(&Service{Name: name}).First(service).Error; err != nil {
		return nil, err
	}

	return service, nil
}
