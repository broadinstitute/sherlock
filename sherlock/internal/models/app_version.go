package models

import "gorm.io/gorm"

type AppVersion struct {
	gorm.Model
	CiIdentifier       *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:app-version"`
	Chart              *Chart
	ChartID            uint
	AppVersion         string
	GitCommit          string
	GitBranch          string
	Description        string
	ParentAppVersion   *AppVersion
	ParentAppVersionID *uint
}

func (a AppVersion) TableName() string {
	return "v2_app_versions"
}

func (a AppVersion) GetCiIdentifier() *CiIdentifier {
	if a.CiIdentifier != nil {
		return a.CiIdentifier
	} else {
		return &CiIdentifier{ResourceType: "app-version", ResourceID: a.ID}
	}
}
