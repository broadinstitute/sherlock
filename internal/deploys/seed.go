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
func SeedServicesAndEnvironments(t *testing.T, db *gorm.DB) {
	err := db.Create(&servicesToSeed).Error
	require.NoError(t, err)

	err = db.Create(&environmentsToSeed).Error
	require.NoError(t, err)
}
