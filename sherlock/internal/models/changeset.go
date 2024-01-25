package models

import (
	"cmp"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Changeset struct {
	gorm.Model
	CiIdentifier   *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:changeset"`
	ChartRelease   *ChartRelease
	ChartReleaseID uint

	From             ChartReleaseVersion `gorm:"embedded;embeddedPrefix:from_"`
	To               ChartReleaseVersion `gorm:"embedded;embeddedPrefix:to_"`
	AppliedAt        *time.Time
	SupersededAt     *time.Time
	NewAppVersions   []*AppVersion   `gorm:"many2many:changeset_new_app_versions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	NewChartVersions []*ChartVersion `gorm:"many2many:changeset_new_chart_versions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`

	PlannedBy   *User
	PlannedByID *uint
	AppliedBy   *User
	AppliedByID *uint
}

func (c *Changeset) GetCiIdentifier() CiIdentifier {
	if c.CiIdentifier != nil {
		return *c.CiIdentifier
	} else {
		return CiIdentifier{ResourceType: "changeset", ResourceID: c.ID}
	}
}

func (c *Changeset) Summarize(includeFrom bool) string {
	summaryParts := make([]string, 0, 3)
	if c.To.AppVersionExact != nil && (c.From.AppVersionExact == nil || *c.From.AppVersionExact != *c.To.AppVersionExact) {
		if includeFrom && c.From.AppVersionExact != nil {
			summaryParts = append(summaryParts, fmt.Sprintf("app %s⭢%s", *c.From.AppVersionExact, *c.To.AppVersionExact))
		} else {
			summaryParts = append(summaryParts, fmt.Sprintf("app⭢%s", *c.To.AppVersionExact))
		}
	}
	if c.To.ChartVersionExact != nil && (c.From.ChartVersionExact == nil || *c.From.ChartVersionExact != *c.To.ChartVersionExact) {
		if includeFrom && c.From.ChartVersionExact != nil {
			summaryParts = append(summaryParts, fmt.Sprintf("chart %s⭢%s", *c.From.ChartVersionExact, *c.To.ChartVersionExact))
		} else {
			summaryParts = append(summaryParts, fmt.Sprintf("chart⭢%s", *c.To.ChartVersionExact))
		}
	}
	if len(summaryParts) == 0 {
		summaryParts = append(summaryParts, "configuration change")
	}
	return strings.Join(summaryParts, ", ")
}

func CompareChangesetsByName(a, b Changeset) int {
	if a.ChartRelease == nil && b.ChartRelease == nil {
		return 0
	} else if a.ChartRelease == nil {
		return -1
	} else if b.ChartRelease == nil {
		return 1
	} else {
		return cmp.Compare(a.ChartRelease.Name, b.ChartRelease.Name)
	}
}

func UsersFromChangesets(changesets []Changeset) []User {
	var users []User
	for _, changeset := range changesets {
		if changeset.AppliedBy != nil {
			// Check if this user is already in our list before adding it
			var exists bool
			for _, existing := range users {
				if existing.ID == changeset.AppliedBy.ID {
					exists = true
					break
				}
			}
			if !exists {
				users = append(users, *changeset.AppliedBy)
			}
		}
		// Only go through the planned by if it's different from the applied by
		if changeset.PlannedBy != nil && (changeset.AppliedBy == nil || changeset.PlannedBy.ID != changeset.AppliedBy.ID) {
			var exists bool
			for _, existing := range users {
				if existing.ID == changeset.PlannedBy.ID {
					exists = true
					break
				}
			}
			if !exists {
				users = append(users, *changeset.PlannedBy)
			}
		}
	}
	return users
}
