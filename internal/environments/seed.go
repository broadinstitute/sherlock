package environments

import "gorm.io/gorm"

// Seed takes a gorm DB connection and will seed a db
// with some fake environment data for use in integration testing
func Seed(db *gorm.DB) ([]Environment, error) {
	seededEnvironments := []Environment{
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
		return []Environment{}, err
	}

	err := db.Find(&seededEnvironments).Error
	return seededEnvironments, err
}
