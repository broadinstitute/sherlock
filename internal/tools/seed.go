package tools

import (
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/services"
)

// SeedServices is a test utility that will populate a database with a predetermined list of "services"
// to be used for running integration tests against a real database
func SeedServices(repository *db.Repository) ([]services.Service, error) {
	services := []services.Service{
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

	err := repository.DB.Create(&services).Error
	if err != nil {
		return nil, err
	}
	err = repository.DB.Find(&services).Error
	return services, err
}

// Truncate cleans up tables after integration tests
func Truncate(repository *db.Repository) error {
	// gorm doesn't seem to support truncate operations which are essential to cleaning up after
	// integration tests (and the only use case of this function so doing it with raw sql)
	truncateStatement := "TRUNCATE TABLE services, builds, environments, service_instances, deploys"
	err := repository.DB.Exec(truncateStatement).Error

	return err
}
