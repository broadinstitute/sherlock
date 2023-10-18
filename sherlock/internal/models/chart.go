package models

import "gorm.io/gorm"

type Chart struct {
	gorm.Model
	CiIdentifier *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:chart"`
	Name         string
	// Mutable
	ChartRepo             *string
	AppImageGitRepo       *string
	AppImageGitMainBranch *string
	ChartExposesEndpoint  *bool
	DefaultSubdomain      *string
	DefaultProtocol       *string
	DefaultPort           *uint
	LegacyConfigsEnabled  *bool
	Description           *string
	PlaybookURL           *string
	PactParticipant       *bool
}

func (c Chart) TableName() string {
	return "charts"
}

func (c Chart) GetCiIdentifier() CiIdentifier {
	if c.CiIdentifier != nil {
		return *c.CiIdentifier
	} else {
		return CiIdentifier{ResourceType: "chart", ResourceID: c.ID}
	}
}
