package models

import "gorm.io/gorm"

type ChartRelease struct {
	gorm.Model
	CiIdentifier    *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:chart-release"`
	Chart           *Chart
	ChartID         uint
	Cluster         *Cluster
	ClusterID       *uint
	DestinationType string
	Environment     *Environment
	EnvironmentID   *uint
	Name            string
	Namespace       string
	ChartReleaseVersion
	Subdomain               *string
	Protocol                *string
	Port                    *uint
	PagerdutyIntegration    *PagerdutyIntegration
	PagerdutyIntegrationID  *uint
	IncludeInBulkChangesets *bool
}

func (c ChartRelease) TableName() string {
	return "v2_chart_releases"
}

func (c ChartRelease) GetCiIdentifier() CiIdentifier {
	if c.CiIdentifier != nil {
		return *c.CiIdentifier
	} else {
		return CiIdentifier{ResourceType: "chart-release", ResourceID: c.ID}
	}
}
