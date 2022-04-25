package deploys

import (
	"fmt"

	"github.com/broadinstitute/sherlock/internal/models"
	"gorm.io/gorm"
)

// SeedServiceInstances is used to populate the database with Service Instance entities
// solely intended for use in testing
func SeedServiceInstances(db *gorm.DB) ([]models.ServiceInstance, error) {
	var (
		services     []models.Service
		environments []models.Environment
	)

	if err := db.Find(&services).Error; err != nil {
		return nil, fmt.Errorf("error retrieving existing services: %v", err)
	}

	if err := db.Find(&environments).Error; err != nil {
		return nil, fmt.Errorf("error retrieving existing environments: %v", err)
	}

	var serviceInstances []models.ServiceInstance
	for _, service := range services {
		for _, environment := range environments {
			serviceInstances = append(serviceInstances, models.ServiceInstance{
				ServiceID:     service.ID,
				EnvironmentID: environment.ID,
			})
		}
	}

	if err := db.Create(&serviceInstances).Error; err != nil {
		return []models.ServiceInstance{}, err
	}

	err := db.Preload("Service").Preload("Environment").Find(&serviceInstances).Error
	return serviceInstances, err
}
