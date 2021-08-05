package tools

import (
	"fmt"

	"github.com/broadinstitute/sherlock/internal/services"
	"gorm.io/gorm"
)

// SeedServices is a test utility that will populate a database with a predetermined list of "services"
// to be used for running integration tests against a real database
func SeedServices(db *gorm.DB) ([]services.Service, error) {
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

	db.Create(&services)
	db.Find(&services)
	if db.Error != nil {
		return nil, db.Error
	}
	return services, nil
}

// Truncate cleans up tables after integration tests
func Truncate(db *gorm.DB) error {
	// gorm doesn't seem to support truncate operations which are essential to cleaning up after
	// integration tests (and the only use case of this function so doing it with raw sql)
	truncateStatement := "TRUNCATE TABLE services, builds, environments, service_instances, deploys"
	db.Exec(truncateStatement)
	if db.Error != nil {
		return fmt.Errorf("error truncating tables: %v", db.Error)
	}
	return nil
}
