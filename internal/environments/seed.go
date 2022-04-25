package environments

import (
	"github.com/broadinstitute/sherlock/internal/models"
	"gorm.io/gorm"
)

// Seed takes a gorm DB connection and will seed a db
// with some fake environment data for use in integration testing
func Seed(db *gorm.DB) ([]models.Environment, error) {
	seededEnvironments := []models.Environment{
		{
			Name: "dev",
		},
		{
			Name: "perf",
		},
		{
			Name: "alpha",
		},
		{
			Name: "staging",
		},
		{
			Name: "prod",
		},
	}

	if err := db.Create(&seededEnvironments).Error; err != nil {
		return []models.Environment{}, err
	}

	err := db.Find(&seededEnvironments).Error
	return seededEnvironments, err
}
