package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
)

type DatabaseInstance struct {
	gorm.Model
	ChartRelease   *ChartRelease
	ChartReleaseID uint

	Platform      *string
	GoogleProject *string
	InstanceName  *string

	DefaultDatabase *string
}

func (d *DatabaseInstance) TableName() string {
	return "database_instances"
}

func (d *DatabaseInstance) errorIfForbidden(tx *gorm.DB) error {
	if d.ChartReleaseID == 0 {
		return fmt.Errorf("(%s) database instance wasn't properly loaded, unable to check permissions on chart release", errors.InternalServerError)
	}
	var chartRelease ChartRelease
	if err := tx.Take(&chartRelease, d.ChartReleaseID).Error; err != nil {
		return fmt.Errorf("unable to get chart release to determine permissions on database instance: %w", err)
	}
	return chartRelease.errorIfForbidden(tx)
}

// BeforeCreate checks permissions
func (d *DatabaseInstance) BeforeCreate(tx *gorm.DB) error {
	return d.errorIfForbidden(tx)
}

// BeforeUpdate checks permissions
func (d *DatabaseInstance) BeforeUpdate(tx *gorm.DB) error {
	return d.errorIfForbidden(tx)
}

// BeforeDelete checks permissions
func (d *DatabaseInstance) BeforeDelete(tx *gorm.DB) error {
	return d.errorIfForbidden(tx)
}
