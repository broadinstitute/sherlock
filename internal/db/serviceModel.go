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

func (s *ServiceModel) ListAll() ([]services.Service, error) {
	services := []services.Service{}

	err := s.DB.Find(&services).Error
	if err != nil {
		return nil, fmt.Errorf("Error retriving services: %v", err)
	}

	return services, nil
}

func NewServiceModel(dbConn *gorm.DB) *ServiceModel {
	return &ServiceModel{DB: dbConn}
}
