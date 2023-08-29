package models

import (
	"gorm.io/gorm"
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
	NewAppVersions   []*AppVersion   `gorm:"many2many:v2_changeset_new_app_versions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	NewChartVersions []*ChartVersion `gorm:"many2many:v2_changeset_new_chart_versions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}

func (c Changeset) TableName() string {
	return "v2_changesets"
}

func (c Changeset) GetCiIdentifier() CiIdentifier {
	if c.CiIdentifier != nil {
		return *c.CiIdentifier
	} else {
		return CiIdentifier{ResourceType: "changeset", ResourceID: c.ID}
	}
}