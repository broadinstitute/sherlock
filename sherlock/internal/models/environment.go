package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type Environment struct {
	gorm.Model
	CiIdentifier              *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:environment"`
	Base                      string
	Lifecycle                 string
	Name                      string
	NamePrefix                string
	TemplateEnvironment       *Environment
	TemplateEnvironmentID     *uint
	ValuesName                string
	AutoPopulateChartReleases *bool
	UniqueResourcePrefix      string
	DefaultNamespace          string
	// Mutable
	DefaultCluster              *Cluster
	DefaultClusterID            *uint
	DefaultFirecloudDevelopRef  *string
	Owner                       *User
	OwnerID                     *uint
	LegacyOwner                 *string
	RequiresSuitability         *bool
	BaseDomain                  *string
	NamePrefixesDomain          *bool
	HelmfileRef                 *string
	PreventDeletion             *bool
	DeleteAfter                 sql.NullTime
	Description                 *string
	PagerdutyIntegration        *PagerdutyIntegration
	PagerdutyIntegrationID      *uint
	Offline                     *bool
	OfflineScheduleBeginEnabled *bool
	OfflineScheduleBeginTime    *string
	OfflineScheduleEndEnabled   *bool
	OfflineScheduleEndTime      *string
	OfflineScheduleEndWeekends  *bool
}

func (e Environment) TableName() string {
	return "v2_environments"
}

func (e Environment) GetCiIdentifier() *CiIdentifier {
	if e.CiIdentifier != nil {
		return e.CiIdentifier
	} else {
		return &CiIdentifier{ResourceType: "environment", ResourceID: e.ID}
	}
}
