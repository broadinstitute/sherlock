package models

import (
	"gorm.io/gorm"
)

type CiIdentifier struct {
	gorm.Model
	ResourceType string
	ResourceID   uint
	// Mutable
	CiRuns []CiRun `gorm:"many2many:ci_runs_for_identifiers"`
}

func (c *CiIdentifier) TableName() string {
	return "ci_identifiers"
}
