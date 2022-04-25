package v1_models

// v1_models.go contains the type for modeling build entities in sherlocks database
// and methods for interacting with the persistence layer. It should only contain
// logic related to interacting with build entities in sherlock's db

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

const duplicateVersionStringErrorCheck string = `duplicate key value violates unique constraint "builds_version_string_key" (SQLSTATE 23505)`

// ErrBuildNotFound is returned when a specific build look up fails
var (
	ErrBuildNotFound          error = gorm.ErrRecordNotFound
	ErrDuplicateVersionString error = errors.New("field version_string for builds must be unique")
	// ErrBadCreateRequest is an error type used when a create servie request fails validation checks
	ErrBadCreateRequest error = errors.New("error invalid create build request")
)

type buildStore struct {
	*gorm.DB
}

// Build is the structure used to represent a build entity in sherlock's db persistence layer
type Build struct {
	ID            int
	VersionString string `gorm:"not null;default:null"`
	CommitSha     string
	BuildURL      string
	BuiltAt       time.Time `gorm:"autoCreateTime"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ServiceID     int
	Service       Service `gorm:"foreignKey:ServiceID;references:ID"`
}

type BuildStore interface {
	ListAll() ([]Build, error)
	CreateNew(Build) (Build, error)
	GetByID(int) (Build, error)
	GetByVersionString(string) (Build, error)
}

func NewBuildStore(dbConn *gorm.DB) buildStore {
	return buildStore{dbConn}
}

func (db buildStore) ListAll() ([]Build, error) {
	builds := make([]Build, 0)

	if err := db.Preload("Service").Find(&builds).Error; err != nil {
		return []Build{}, fmt.Errorf("error listing builds: %v", err)
	}

	return builds, nil
}

func (db buildStore) CreateNew(newBuild Build) (Build, error) {
	err := db.Create(&newBuild).Error
	if err != nil {
		// check for error due to duplicate VersionString field
		if strings.Contains(err.Error(), duplicateVersionStringErrorCheck) {
			return Build{}, ErrDuplicateVersionString
		}
		return Build{}, ErrBadCreateRequest
	}
	// retrieve the same build record back from DB so it can be returned with associations updated
	// gorm will not update the associations in the input struct when performing create operations,
	// even though those associations will be modeled properly in the db
	err = db.Preload("Service").First(&newBuild, newBuild.ID).Error
	return newBuild, err
}

func (db buildStore) GetByID(id int) (Build, error) {
	build := Build{}

	if err := db.Preload("Service").First(&build, id).Error; err != nil {
		return Build{}, err
	}
	return build, nil
}

func (db buildStore) GetByVersionString(versionString string) (Build, error) {
	build := Build{}

	if err := db.Preload("Service").First(&build, &Build{VersionString: versionString}).Error; err != nil {
		return Build{}, err
	}

	return build, nil
}
