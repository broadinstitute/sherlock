package deploys

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/models/v1models"

	"gorm.io/gorm"
)

// SeedServiceInstances is used to populate the database with Service Instance entities
// solely intended for use in testing
func SeedServiceInstances(db *gorm.DB) ([]v1models.ServiceInstance, error) {
	var (
		services     []v1models.Service
		environments []v1models.Environment
	)

	if err := db.Find(&services).Error; err != nil {
		return nil, fmt.Errorf("error retrieving existing services: %v", err)
	}

	if err := db.Find(&environments).Error; err != nil {
		return nil, fmt.Errorf("error retrieving existing environments: %v", err)
	}

	var serviceInstances []v1models.ServiceInstance
	for _, service := range services {
		for _, environment := range environments {
			serviceInstances = append(serviceInstances, v1models.ServiceInstance{
				ServiceID:     service.ID,
				EnvironmentID: environment.ID,
			})
		}
	}

	if err := db.Create(&serviceInstances).Error; err != nil {
		return []v1models.ServiceInstance{}, err
	}

	err := db.Preload("Service").Preload("Environment").Find(&serviceInstances).Error
	return serviceInstances, err
}
