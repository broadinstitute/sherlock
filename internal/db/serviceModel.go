package db

import (
	"fmt"

	"github.com/broadinstitute/sherlock/internal/services"
	"gorm.io/gorm"
)

// ServiceModel is an implementation of services.ServiceModel interface
//  which supports the service operations with a postgres backend
type ServiceModel struct {
	DB *gorm.DB
}

// ListAll retrieves all service entities from a postgres database and
// returns them as a slice
func (s *ServiceModel) ListAll() ([]services.Service, error) {
	services := []services.Service{}

	err := s.DB.Find(&services).Error
	if err != nil {
		return nil, fmt.Errorf("Error retriving services: %v", err)
	}

	return services, nil
}

// Create will persist the service defined by newservice to a postgres database.
// It will return the service as stored in postgres for ease of testing if successful
func (s *ServiceModel) Create(newService *services.Service) (*services.Service, error) {
	if err := s.DB.Create(newService).Error; err != nil {
		return nil, fmt.Errorf("error saving service to database: %v", err)
	}
	return newService, nil
}

// Get is used to retrieve a specific service entity from a postgres database using
// id (primary key) as the lookup mechanism
func (s *ServiceModel) Get(id string) (*services.Service, error) {
	service := &services.Service{}

	if err := s.DB.First(service, id).Error; err != nil {
		return nil, err
	}
	return service, nil
}

// NewServiceModel constructs an object that can perform
// crud operations on service entities stored in a postgres database
func NewServiceModel(dbConn *gorm.DB) *ServiceModel {
	return &ServiceModel{DB: dbConn}
}
