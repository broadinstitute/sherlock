package db

import (
	"fmt"

	"github.com/broadinstitute/sherlock/internal/builds"
	"gorm.io/gorm"
)

type BuildModel struct {
	DB *gorm.DB
}

func (b *BuildModel) ListAll() ([]builds.Build, error) {
	builds := []builds.Build{}

	if err := b.DB.Preload("Service").Find(&builds).Error; err != nil {
		return nil, fmt.Errorf("Error retrieving builds: %v", err)
	}

	return builds, nil
}

// NewBuildModel constructs an object that can perform
// crud operations on Build entities stored in a postgres database
func NewBuildModel(dbConn *gorm.DB) *BuildModel {
	return &BuildModel{DB: dbConn}
}
