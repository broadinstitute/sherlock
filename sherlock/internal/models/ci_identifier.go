package models

import (
	"gorm.io/gorm"
)

type CiIdentifier struct {
	gorm.Model
	ResourceType string
	ResourceID   uint
	// Mutable
	CiRuns []CiRun `gorm:"many2many:v2_ci_runs_for_identifiers"`
}

func (c *CiIdentifier) TableName() string {
	return "v2_ci_identifiers"
}
