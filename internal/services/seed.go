package services

import (
	"github.com/broadinstitute/sherlock/internal/models/v1_models"
	"gorm.io/gorm"
)

// Seed is a test utility that will populate a database with a predetermined list of "services"
// to be used for running integration tests against a real database
func Seed(db *gorm.DB) ([]v1_models.Service, error) {
	services := []v1_models.Service{
		{
			Name:    "cromwell",
			RepoURL: "https://github.com/broadinstitute/cromwell",
		},
		{
			Name:    "leonardo",
			RepoURL: "https://github.com/DataBiosphere/leonardo",
		},
		{
			Name:    "workspacemanager",
			RepoURL: "https://github.com/DataBiosphere/terra-workspace-manager",
		},
	}

	err := db.Create(&services).Error
	if err != nil {
		return nil, err
	}

	err = db.Find(&services).Error
	return services, err
}
