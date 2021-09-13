package builds

import (
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/internal/services"
	"gorm.io/gorm"
)

type dataStore struct {
	*gorm.DB
}

// Build is the structure used to represent a build entity in sherlock's db persistence layer
type Build struct {
	ID            int
	VersionString string
	CommitSha     string
	BuildURL      string
	BuiltAt       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ServiceID     int
	Service       services.Service
}

type buildStore interface {
	listAll() ([]Build, error)
	createNew(CreateBuildRequest) (Build, error)
}

func newBuildStore(dbConn *gorm.DB) dataStore {
	return dataStore{dbConn}
}

func (db dataStore) listAll() ([]Build, error) {
	builds := make([]Build, 0)

	if err := db.Preload("Service").Find(&builds).Error; err != nil {
		return []Build{}, fmt.Errorf("error listing builds: %v", err)
	}

	return builds, nil
}

func (db dataStore) createNew(newBuild CreateBuildRequest) (Build, error) {
	return Build{}, nil
}
