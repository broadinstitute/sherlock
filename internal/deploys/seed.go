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
func SeedServiceInstances(t *testing.T, db *gorm.DB) ([]ServiceInstance, error) {
	var serviceInstances []ServiceInstance
	// perform all the seeding operations in a transaction to avoid odd race conditions
	err := db.Transaction(func(tx *gorm.DB) error {
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

	require.NoError(t, err)

	err = db.Preload("Service").Preload("Environment").Find(&serviceInstances).Error
	t.Cleanup(func() {
		db.Delete(&serviceInstances)
		db.Delete(&environmentsToSeed)
		db.Delete(&servicesToSeed)
	})

	return serviceInstances, err
}
