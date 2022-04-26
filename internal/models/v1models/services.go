package v1models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ErrServiceNotFound is the error to represent a failed lookup of a service entity
var (
	ErrServiceNotFound = gorm.ErrRecordNotFound
)

// dataStore is a wrapper around a gorm postgres client
// which can be used to implement the serviceRepository interface
type serviceStore struct {
	*gorm.DB
}

// Service is the data structure that v1models a service entity persisted to a database via gorm
type Service struct {
	ID        int    `gorm:"primaryKey" faker:"unique"`
	Name      string `gorm:"not null;default:null"`
	RepoURL   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// serviceStore is the interface type that defines the methods required for implementing the persistence layer
// for services entities
type ServiceStore interface {
	ListAll() ([]Service, error)
	CreateNew(CreateServiceRequest) (Service, error)
	GetByName(string) (Service, error)
}

func NewServiceStore(db *gorm.DB) serviceStore {
	return serviceStore{db}
}

// CreateServiceRequest is a type used to represent the information required to register a new service in sherlock
type CreateServiceRequest struct {
	Name    string `json:"name" binding:"required"`
	RepoURL string `json:"repo_url" binding:"required"`
}

// creates a service entity object to be persisted with the database from a
// request to create a service
func (cr CreateServiceRequest) Service() Service {
	return Service{
		Name:    cr.Name,
		RepoURL: cr.RepoURL,
	}
}

// ListAll retrieves all service entities from a postgres database and
// returns them as a slice
func (db serviceStore) ListAll() ([]Service, error) {
	services := []Service{}

	err := db.Find(&services).Error
	if err != nil {
		return []Service{}, fmt.Errorf("error retriving services: %v", err)
	}

	return services, nil
}

// Create will persist the service defined by newservice to a postgres database.
// It will return the service as stored in postgres for ease of testing if successful
func (db serviceStore) CreateNew(newServiceReq CreateServiceRequest) (Service, error) {
	newService := newServiceReq.Service()

	if err := validateNotEmpty(newServiceReq.Name); err != nil {
		return Service{}, fmt.Errorf("error saving service to database: %v", err)
	}

	if err := db.Create(&newService).Error; err != nil {
		return Service{}, fmt.Errorf("error saving service to database: %v", err)
	}
	return newService, nil
}

// getByName retrives a service entity from persistence layer. It returns an ErrRecordNotFound
// if the requested name does not exist. This effectively equivalent to get service by id
// as the name field is indexed and enforced to be unique
func (db serviceStore) GetByName(name string) (Service, error) {
	service := Service{}

	if err := db.Where("name = ?", name).First(&service).Error; err != nil {
		return Service{}, err
	}

	return service, nil
}

// validates that a string is not empty, important for any situation where go will default "",
// which technically passes a NULL database check but is not wanted behavior
func validateNotEmpty(stringToValidate string) error {
	if stringToValidate == "" {
		err := errors.New("database field cannot be empty")
		return err
	}

	return nil
}
