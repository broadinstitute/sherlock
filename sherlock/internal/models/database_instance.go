package models

import "gorm.io/gorm"

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
	if d.ChartRelease == nil {
		if err := tx.Take(&d.ChartRelease, d.ChartReleaseID).Error; err != nil {
			return err
		}
	}
	return d.ChartRelease.errorIfForbidden(tx)
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
