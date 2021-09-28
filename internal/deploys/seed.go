package deploys

import (
	"testing"
	"time"

	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var (
	servicesToSeed = []services.Service{
		{
			Name:      "sam",
			RepoURL:   "blah.blah.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "buffer",
			RepoURL:   "github.com/databiosphere/buffer",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "rawls",
			RepoURL:   "github.com/broadinstitute/rawls",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	environmentsToSeed = []environments.Environment{
		{
			Name:        "dev",
			IsPermanent: true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "alpha",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			IsPermanent: true,
		},
		{
			Name:        "prod",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			IsPermanent: true,
		},
	}
)

// SeedServiceInstances is used to populate the database with Service Instance entities
// solely intended for use in testing
func SeedServiceInstances(t *testing.T, db *gorm.DB) []ServiceInstance {
	var serviceInstances []ServiceInstance

	// use a tranaction to populate service instance test data is it
	// is a complex operation requiring multiple associations
	err := db.Transaction(func(tx *gorm.DB) error {
		// use a nested transaction for populating services and environments to
		// guarantee they exist before trying to reference them in service instances
		err := tx.Transaction(func(tx2 *gorm.DB) error {
			if err := tx2.Create(&servicesToSeed).Error; err != nil {
				return err
			}

			if err := tx2.Create(&environmentsToSeed).Error; err != nil {
				return err
			}

			return nil
		})
		require.NoError(t, err)

		for _, service := range servicesToSeed {
			for _, environment := range environmentsToSeed {
				serviceInstances = append(serviceInstances, ServiceInstance{
					ServiceID:     service.ID,
					EnvironmentID: environment.ID,
				})
			}
		}

		if err := tx.Create(&serviceInstances).Error; err != nil {
			return err
		}

		return nil
	})

	// load all the data just seeded to compare against it in tests
	err = db.Preload("Service").Preload("Environment").Find(&serviceInstances).Error
	require.NoError(t, err)

	return serviceInstances
}
