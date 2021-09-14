package builds

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/broadinstitute/sherlock/internal/services"
	"gorm.io/gorm"
)

const duplicateVersionStringErrorCheck string = `duplicate key value violates unique constraint "builds_version_string_key" (SQLSTATE 23505)`

// ErrBuildNotFound is returned when a specific build look up fails
var (
	ErrBuildNotFound          error = gorm.ErrRecordNotFound
	ErrDuplicateVersionString       = errors.New("field version_string for builds must be unique")
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
	createNew(*Build) (*Build, error)
	getByID(int) (*Build, error)
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

func (db dataStore) createNew(newBuild *Build) (*Build, error) {
	err := db.Create(newBuild).Error
	if err != nil {
		// check for error due to duplicate VersionString field
		if strings.Contains(err.Error(), duplicateVersionStringErrorCheck) {
			return nil, ErrDuplicateVersionString
		}
		return nil, fmt.Errorf("error persisting new build: %v", err)
	}
	// retrieve the same build record back from DB so it can be returned with associations updated
	// gorm will not update the associations in the input struct when performing create operations,
	// even though those associations will be modeled properly in the db
	err = db.Preload("Service").First(newBuild, newBuild.ID).Error
	return newBuild, err
}

func (db dataStore) getByID(id int) (*Build, error) {
	build := &Build{}

	if err := db.Preload("Service").First(build, id).Error; err != nil {
		return nil, err
	}
	return build, nil
}
