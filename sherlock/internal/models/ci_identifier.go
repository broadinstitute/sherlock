package models

import (
	"fmt"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"gorm.io/gorm"
)

type CiIdentifier struct {
	gorm.Model
	ResourceType string
	ResourceID   uint
	// Mutable
	CiRuns []CiRun `gorm:"many2many:ci_runs_for_identifiers"`

	// ResourceStatus is ignored by Gorm and isn't stored in the database -- at least, not
	// on the CiIdentifier type itself. The data is actually stored on CiRunIdentifierJoin,
	// and this field exists so that when a CiIdentifier is loaded via a CiRun it can hold
	// the resource-specific status from the join table along the way.
	ResourceStatus *string `gorm:"-:all"`
}

func (c *CiIdentifier) FillCiRunResourceStatuses(db *gorm.DB) error {
	var joinEntries []CiRunIdentifierJoin
	if err := db.
		Model(&CiRunIdentifierJoin{}).
		Where("ci_identifier_id = ? AND ci_run_id IN ? AND resource_status IS NOT NULL",
			c.ID, utils.Map(c.CiRuns, func(cr CiRun) uint { return cr.ID })).
		Limit(len(c.CiRuns)).
		Find(&joinEntries).
		Error; err != nil {
		return fmt.Errorf("failed to query join table for ci run resource statuses: %w", err)
	}
	for _, joinEntry := range joinEntries {
		if joinEntry.ResourceStatus != nil {
			for index, ciRun := range c.CiRuns {
				if ciRun.ID == joinEntry.CiRunID && c.ID == joinEntry.CiIdentifierID {
					// dereference and reference so we are extra sure we don't cross wires while iterating
					ciRun.ResourceStatus = utils.PointerTo(*joinEntry.ResourceStatus)
					c.CiRuns[index] = ciRun
				}
			}
		}
	}
	return nil
}
