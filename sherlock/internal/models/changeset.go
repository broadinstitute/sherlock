package models

import (
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
	if c.From.AppVersionExact != c.To.AppVersionExact && c.To.AppVersionExact != nil {
		if includeFrom && c.From.AppVersionExact != nil {
			summaryParts = append(summaryParts, fmt.Sprintf("app %s⭢%s", *c.From.AppVersionExact, *c.To.AppVersionExact))
		} else {
			summaryParts = append(summaryParts, fmt.Sprintf("app⭢%s", *c.To.AppVersionExact))
		}
	}
	if c.From.ChartVersionExact != c.To.ChartVersionExact && c.To.ChartVersionExact != nil {
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
