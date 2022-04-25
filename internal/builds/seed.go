package builds

import (
	"fmt"
	v1_models2 "github.com/broadinstitute/sherlock/internal/models/v1_models"
	"time"

	"gorm.io/gorm"
)

// Seed is a testing utility used in integration tests
// to populate a postgres DB with fake Build entities
func Seed(db *gorm.DB) ([]v1_models2.Build, error) {
	// get existing services to make sure ids are valid.

	// used to verify we can explicity set BuiltAt rather than just defaulting to current time
	sixHoursAgo := time.Now().Add(-6 * time.Hour)
	var services []v1_models2.Service
	if err := db.Find(&services).Error; err != nil {
		return []v1_models2.Build{}, fmt.Errorf("error retrieving existing services to reference in seeded builds: %v", err)
	}
	builds := []v1_models2.Build{
		{
			VersionString: "gcr.io/cromwell:0.1.0",
			CommitSha:     "k2jh34",
			BuildURL:      "https://build.1.log",
			ServiceID:     services[0].ID,
		},
		{
			VersionString: "gcr.io/cromwell:0.2.0",
			CommitSha:     "lk2j344",
			BuildURL:      "https://build.2.log",
			ServiceID:     services[0].ID,
		},
		{
			VersionString: "grc.io/leonardo:0.1.0",
			CommitSha:     "k2jh34",
			BuildURL:      "https://build.1.log",
			ServiceID:     services[1].ID,
		},
		{
			VersionString: "gcr.io/workspacemanager:1.1.0",
			CommitSha:     "lk23j4",
			BuildURL:      "https://build.3.log",
			ServiceID:     services[2].ID,
		},
		{
			VersionString: "gcr.io/workspacemanager:1.1.1",
			CommitSha:     "asdfbvf",
			BuildURL:      "https://build.3.log",
			ServiceID:     services[2].ID,
		},
		{
			VersionString: "gcr.io/workspacemanager:1.2.0",
			CommitSha:     "6a5s4df",
			BuildURL:      "https://build.3.log",
			BuiltAt:       sixHoursAgo,
			ServiceID:     services[2].ID,
		},
	}

	err := db.Create(&builds).Error
	if err != nil {
		return nil, err
	}
	err = db.Preload("Service").Find(&builds).Error
	return builds, err
}
