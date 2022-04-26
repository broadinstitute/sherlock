package v1models

import (
	"fmt"
	"gorm.io/gorm"
)

// SeedServiceInstances is used to populate the database with Service Instance entities
// solely intended for use in testing
func SeedServiceInstances(db *gorm.DB) ([]ServiceInstance, error) {
	var (
		services     []Service
		environments []Environment
	)

	if err := db.Find(&services).Error; err != nil {
		return nil, fmt.Errorf("error retrieving existing services: %v", err)
	}

	if err := db.Find(&environments).Error; err != nil {
		return nil, fmt.Errorf("error retrieving existing environments: %v", err)
	}

	var serviceInstances []ServiceInstance
	for _, service := range services {
		for _, environment := range environments {
			serviceInstances = append(serviceInstances, ServiceInstance{
				ServiceID:     service.ID,
				EnvironmentID: environment.ID,
			})
		}
	}

	if err := db.Create(&serviceInstances).Error; err != nil {
		return []ServiceInstance{}, err
	}

	err := db.Preload("Service").Preload("Environment").Find(&serviceInstances).Error
	return serviceInstances, err
}
