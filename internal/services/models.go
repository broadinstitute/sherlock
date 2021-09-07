package services

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// dataStore is a wrapper around a gorm postgres client
// which can be used to implement the serviceRepository interface
type dataStore struct {
	*gorm.DB
}

// Service is the data structure representing an used to model a service entity
// in databases
type Service struct {
	ID        int
	Name      string
	RepoURL   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// serviceStore is the interface type used
type serviceStore interface {
	ListAll() ([]Service, error)
	CreateNew(*Service) (*Service, error)
	GetByID(string) (*Service, error)
}

func newServiceStore(db *gorm.DB) dataStore {
	return dataStore{db}
}

// ListAll retrieves all service entities from a postgres database and
// returns them as a slice
func (db dataStore) ListAll() ([]Service, error) {
	services := []Service{}

	err := db.Find(&services).Error
	if err != nil {
		return nil, fmt.Errorf("Error retriving services: %v", err)
	}

	return services, nil
}

// Create will persist the service defined by newservice to a postgres database.
// It will return the service as stored in postgres for ease of testing if successful
func (db dataStore) CreateNew(newService *Service) (*Service, error) {
	if err := db.Create(newService).Error; err != nil {
		return nil, fmt.Errorf("error saving service to database: %v", err)
	}
	return newService, nil
}

// Get is used to retrieve a specific service entity from a postgres database using
// id (primary key) as the lookup mechanism
func (db dataStore) GetByID(id string) (*Service, error) {
	service := &Service{}

	if err := db.First(service, id).Error; err != nil {
		return nil, err
	}
	return service, nil
}
