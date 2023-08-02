package models

import "gorm.io/gorm"

type ChartVersion struct {
	gorm.Model
	CiIdentifier         *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:chart-version"`
	Chart                *Chart
	ChartID              uint
	ChartVersion         string
	Description          string
	ParentChartVersion   *ChartVersion
	ParentChartVersionID *uint
}

func (c ChartVersion) TableName() string {
	return "v2_chart_versions"
}

func (c ChartVersion) GetCiIdentifier() *CiIdentifier {
	if c.CiIdentifier != nil {
		return c.CiIdentifier
	} else {
		return &CiIdentifier{ResourceType: "chart-version", ResourceID: c.ID}
	}
}
