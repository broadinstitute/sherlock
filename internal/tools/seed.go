package tools

import (
	"github.com/broadinstitute/sherlock/internal/builds"
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

	err := db.Create(&services).Error
	if err != nil {
		return nil, err
	}
	err = db.Find(&services).Error
	return services, err
}

func SeedBuilds(db *gorm.DB) ([]builds.Build, error) {
	builds := []builds.Build{
		{
			VersionString: "0.1.0",
			CommitSha:     "k2jh34",
			BuildURL:      "https://build.1.log",
			ServiceID:     1,
		},
		{
			VersionString: "0.2.0",
			CommitSha:     "lk2j344",
			BuildURL:      "https://build.2.log",
			ServiceID:     1,
		},
		{
			VersionString: "0.1.0",
			CommitSha:     "k2jh34",
			BuildURL:      "https://build.1.log",
			ServiceID:     2,
		},
		{
			VersionString: "1.1.0",
			CommitSha:     "lk23j4",
			BuildURL:      "https://build.3.log",
			ServiceID:     3,
		},
		{
			VersionString: "1.1.1",
			CommitSha:     "asdfbvf",
			BuildURL:      "https://build.3.log",
			ServiceID:     3,
		},
		{
			VersionString: "1.2.0",
			CommitSha:     "6a5s4df",
			BuildURL:      "https://build.3.log",
			ServiceID:     3,
		},
	}
	err := db.Create(&builds).Error
	if err != nil {
		return nil, err
	}
	err = db.Find(&builds).Error
	return builds, err
}

// Truncate cleans up tables after integration tests
func Truncate(db *gorm.DB) error {
	// gorm doesn't seem to support truncate operations which are essential to cleaning up after
	// integration tests (and the only use case of this function so doing it with raw sql)
	truncateStatement := "TRUNCATE TABLE services, builds, environments, service_instances, deploys"
	err := db.Exec(truncateStatement).Error

	return err
}
